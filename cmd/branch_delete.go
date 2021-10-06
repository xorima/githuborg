package cmd

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xorima/githuborg/pkg/branch"
)

// branchDeleteCmd represents the branch delete command
var branchDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete all matching branches",
	Long:  `Deletes all branches that match the given branch-name in all repos within the org with the given topic name`,
	Run: func(cmd *cobra.Command, args []string) {
		branchName := getStringFlag("branch-name", cmd)
		owner := getStringFlag("org", cmd)
		topic := getStringFlag("topic", cmd)

		token := os.Getenv("GITHUB_TOKEN")
		ctx := context.Background()
		client := branch.CreateGithubClient(ctx, token)
		repos, _ := branch.SearchForRepoByTopicInOrg(ctx, owner, topic, client)

		for _, repo := range repos {
			// branches, err := listBranchesForRepo(owner, v.GetName(), client, ctx)
			branchExists := branch.BranchExists(client, ctx, owner, repo.GetName(), branchName)
			if !branchExists {
				log.Printf("Branch %v does not exist in repo: %v\n", branchName, repo.GetName())
				// skip to next repo, nothing to do here.
				continue
			}
			err := branch.CloseOpenPullRequestByBranch(client, ctx, owner, repo.GetName(), branchName)
			if err != nil {
				log.Fatalf("Error in closing pull request for repo: %v, %v", repo.GetName(), err)
			}
			err = branch.DeleteBranchByName(client, ctx, owner, repo.GetName(), branchName)
			if err != nil {
				log.Fatalf("Error in deleting branch for repo: %v, %v", repo.GetName(), err)
			}
			log.Printf("Branch: %v and any associated PR in Repo: %v/%v have been closed\n", branchName, owner, repo.GetName())
		}

	},
}

func init() {
	branchCmd.AddCommand(branchDeleteCmd)
}

package cmd

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xorima/githuborg/pkg/branch"
)

// branchApproveCmd represents the branch approve command
var branchApproveCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approves all matching branches",
	Long:  `Approves all pull requests that have the given branch-name and are open in all repos within the org with the given topic name`,
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
			branch.ApprovePullRequestByBranch(client, ctx, owner, repo.GetName(), branchName)
			log.Printf("Pull Requests for branch %v in Repo: %v/%v have been approved\n", branchName, owner, repo.GetName())
		}

	},
}

func init() {
	branchCmd.AddCommand(branchApproveCmd)
}

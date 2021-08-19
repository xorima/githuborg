package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	owner := "sous-chefs"
	token := os.Getenv("GITHUB_TOKEN")
	client := createGithubClient(token)
	repos, _ := searchForRepoByTopicInOrg(owner, "bot-trainer", client)
	branchName := "automated/standardfiles"

	for _, v := range repos {
		fmt.Print("**************")
		fmt.Println(v.GetName())

		branch, _, err := client.Repositories.ListBranches(context.Background(), owner, v.GetName(), nil)
		if err != nil {
			log.Fatal(err)
		}
		for _, b := range branch {
			fmt.Println(b.GetName())
			if b.GetName() == branchName {
				// deleteOpenPullRequestByBranch()
				opt := &github.PullRequestListOptions{Head: b.GetName(), State: "open"}
				pr, _, err := client.PullRequests.List(context.Background(), owner, v.GetName(), opt)
				if err != nil {
					log.Fatal(err)
				}
				println(len(pr))
				if len(pr) == 1 {
					fmt.Println(pr[0].GetNumber())
					pr_update := pr[0]
					state := "closed"
					pr_update.State = &state
					client.PullRequests.Edit(context.Background(), owner, v.GetName(), pr[0].GetNumber(), pr_update)
				}
				// deleteBranchByName()
				// client.Repositories.GetBranch(context.Background(), owner, v.GetName(), branchName)

				// Branch needs to be prefixed with head, because of course it does.
				branchNameWithHead := fmt.Sprintf("heads/%v", branchName)

				resp, err := client.Git.DeleteRef(context.Background(), owner, v.GetName(), branchNameWithHead)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(resp.StatusCode)

			}
		}
	}
}

func createGithubClient(token string) *github.Client {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func searchForRepoByTopicInOrg(orgName, topic string, client *github.Client) (repositories []github.Repository, err error) {
	opts := &github.SearchOptions{Sort: "created", Order: "asc"}
	query := fmt.Sprintf("topic:%v org:%v", topic, orgName)
	res, _, err := client.Search.Repositories(context.Background(), query, opts)

	if err != nil {
		return nil, err
	}
	return res.Repositories, nil
}

// func listBranchesForRepo(orgName, repoName string, client *github.Client) (branches []*github.Branch, err error) {
// 	branches, _, err := client.Repositories.ListBranches(context.Background(), orgName, repoName, nil)
// 	if err != nil {

// 	}
// 	return branches, err
// }

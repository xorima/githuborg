package branch

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func CreateGithubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func SearchForRepoByTopicInOrg(ctx context.Context, orgName, topic string, client *github.Client) (repositories []github.Repository, err error) {
	opts := &github.SearchOptions{Sort: "created", Order: "asc", ListOptions: github.ListOptions{PerPage: 100}}
	query := fmt.Sprintf("topic:%v org:%v", topic, orgName)

	var allRepos []github.Repository

	for {
		results, resp, err := client.Search.Repositories(ctx, query, opts)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, results.Repositories...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	if err != nil {
		return nil, err
	}
	return allRepos, nil
}

func BranchExists(client *github.Client, ctx context.Context, owner, repoName, branchName string) bool {
	_, response, err := client.Repositories.GetBranch(ctx, owner, repoName, branchName)
	if response.Response.StatusCode == 404 {
		return false
	}
	if err != nil {
		log.Fatalf("Error occoured in getting branch for %v stacktrace: %v", repoName, err)
	}
	return true
}

func listOpenPullRequestByBranch(client *github.Client, ctx context.Context, owner, repoName, branchName string) (*github.PullRequest, error) {

	// deleteOpenPullRequestByBranch()
	opt := &github.PullRequestListOptions{Head: branchName, State: "open"}
	pr, _, err := client.PullRequests.List(ctx, owner, repoName, opt)
	if err != nil {
		return nil, err
	}
	// No pr exists, so return
	if len(pr) != 1 {
		return nil, err
	}
	return pr[0], err
}

func CloseOpenPullRequestByBranch(client *github.Client, ctx context.Context, owner, repoName, branchName string) error {
	pullRequest, err := listOpenPullRequestByBranch(client, ctx, owner, repoName, branchName)
	if err != nil {
		return err
	}
	if &pullRequest != nil {
		return nil
	}
	state := "closed"
	pullRequest.State = &state
	_, _, err = client.PullRequests.Edit(ctx, owner, repoName, pullRequest.GetNumber(), pullRequest)
	return err
}

func DeleteBranchByName(client *github.Client, ctx context.Context, owner, repoName, branchName string) error {
	branchNameWithHead := fmt.Sprintf("heads/%v", branchName)
	_, err := client.Git.DeleteRef(ctx, owner, repoName, branchNameWithHead)
	return err
}

func ApprovePullRequestByBranch(client *github.Client, ctx context.Context, owner, repoName, branchName string) {
	pullRequest, err := listOpenPullRequestByBranch(client, ctx, owner, repoName, branchName)
	if err != nil {
		log.Fatal(err)
	}

	approve := "APPROVE"
	pullRequestReview := &github.PullRequestReviewRequest{
		Event: &approve,
	}

	_, _, err = client.PullRequests.CreateReview(ctx, owner, repoName, pullRequest.GetNumber(), pullRequestReview)
	if err != nil {
		log.Fatal(err)
	}
}

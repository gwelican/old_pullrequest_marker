package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v24/github"
	"github.com/namsral/flag"
	"golang.org/x/oauth2"
	"net/url"
	"os"
	"time"
)

func main() {

	var days int
	var token string
	var owner string
	var repo string

	flag.IntVar(&days, "days", 5, "")
	flag.StringVar(&token, "token", "", "")
	flag.StringVar(&owner, "owner", "", "")
	flag.StringVar(&repo, "repo", "", "")
	flag.Parse()
	ctx := context.Background()

	client := createGithubClient(ctx, token)
	pullRequests := getPullRequests(client, ctx, owner, repo)

	for _, r := range pullRequests {

		duration := time.Duration(days*24) * time.Hour
		if time.Now().After(r.CreatedAt.Add(duration)) {
			labels := []string{"old"}
			fmt.Printf("Marking %d PR with old flag\n", r.GetNumber())
			_, _, err := client.Issues.AddLabelsToIssue(ctx, "ECP", "pacman", *r.Number, labels)
			checkErrorAndExit(err)
		}
	}

}

func getPullRequests(client *github.Client, ctx context.Context, owner string, repo string) []*github.PullRequest {
	options := &github.PullRequestListOptions{Sort: "updated", Direction: "desc"}
	pullRequests, _, err := client.PullRequests.List(ctx, owner, repo, options)
	checkErrorAndExit(err)
	return pullRequests
}

func createGithubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	githubBaseUrl, err := url.Parse("https://github.expedia.biz/api/v3/")
	client.BaseURL = githubBaseUrl
	checkErrorAndExit(err)
	return client
}

func checkErrorAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

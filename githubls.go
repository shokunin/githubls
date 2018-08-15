package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func main() {

	githubOrg := flag.String("org", "", "Github Organization Name")
	flag.Parse()

	ctx := context.Background()

	// oauth
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 20},
	}

	// check env var
	if os.Getenv("GITHUB_TOKEN") == "" {
		fmt.Println("Please set the GITHUB_TOKEN environment variable")
		os.Exit(1)
	}

	// check organization set
	if *githubOrg == "" {
		fmt.Println("Please see usage: githubls -h")
		os.Exit(1)
	}

	// get all pages of results
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, *githubOrg, opt)
		if err != nil {
			log.Fatal(err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for _, repo := range allRepos {
		fmt.Println(*repo.Name)
	}
}

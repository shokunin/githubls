package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func main() {

	ctx := context.Background()

	// oauth
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, "", nil)

	if err != nil {
		log.Fatal(err)
	}

	for i, repo := range repos {
		fmt.Println(i)
		fmt.Println(repo)
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Client struct {
	ctx    context.Context
	github *github.Client
}

func main() {
	org := flag.String("org", "", "the github organization name")
	repo := flag.String("repo", "", "the repository name")
	prid := flag.Int("pr", 0, "the pull request id")
	flag.Parse()

	if *org == "" || *repo == "" || *prid == 0 {
		panic("Ensure org, repo, and pr are set")
	}
	// should run the program and output the list of files
	client, _ := newClient()

	pageNumber := 0
	for {
		opts := github.ListOptions{
			PerPage: 300,
		}
		if pageNumber != 0 {
			opts.Page = pageNumber
		}

		pageFiles, resp, err := client.github.PullRequests.ListFiles(client.ctx, *org, *repo, *prid, &opts)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode == 200 {
			for _, f := range pageFiles {
				fmt.Printf("%+v\n", f.GetFilename())
			}
		} else {
			// fmt.Printf("%+v\n", pageFiles)
			// fmt.Printf("%+v\n", resp)
			// fmt.Printf("%+v\n", err)
		}
		if resp.NextPage == 0 {
			break
		}
		pageNumber = resp.NextPage
	}
}

func newClient() (*Client, error) {
	token := os.Getenv("GITHUB_OAUTH_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	// If we're using github.com then we don't need to do any additional configuration
	// for the client. It we're using Github Enterprise, then we need to manually
	// set the base url for the API.
	// if hostname != "github.com" {
	// 	// baseURL := fmt.Sprintf("https://%s/api/v3/", hostname)
	// 	base, err := url.Parse(baseURL)
	// 	if err != nil {
	// 		return nil, errors.Wrapf(err, "Invalid github hostname trying to parse %s", baseURL)
	// 	}
	// 	client.BaseURL = base
	// }

	return &Client{
		ctx:    ctx,
		github: client,
	}, nil
}

func (c *Client) PullRequestFiles() {
	// with the org, repo, pr id list the files
	// pagination is necessary
}

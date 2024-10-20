package gh

import (
	"context"
	"fmt"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type ghApi struct {
	client *githubv4.Client
}

func NewApi(access_token string) *ghApi {
	client := &ghApi{}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client.client = githubv4.NewClient(httpClient)

	return client
}

type PullRequest struct {
	Title     string
	Permalink string
	Additions int
	Deletions int
	CreatedAt time.Time
	IsDraft   bool
	Mergeable githubv4.MergeableState
	Merged    bool
	Closed    bool
}

func (api ghApi) GetOpenPRs(repo_owner string, repo_name string) []PullRequest {
	// Define query variables
	variables := map[string]interface{}{
		"owner":  githubv4.String(repo_owner),
		"name":   githubv4.String(repo_name),
		"states": []githubv4.PullRequestState{githubv4.PullRequestStateOpen},
	}

	// Construct query
	var query struct {
		Repository struct {
			PullRequests struct {
				Nodes []PullRequest
			} `graphql:"pullRequests(first: 20, states: $states)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	// Execute query
	err := api.client.Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println("ERROR", err.Error())
	}
	return query.Repository.PullRequests.Nodes
}

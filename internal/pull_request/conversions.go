package pull_request

import (
	"github.com/devenjarvis/pr-staton/internal/model"
	"github.com/devenjarvis/pr-staton/pkg/gh"
	"github.com/shurcooL/githubv4"
)

func pullRequestFromGhPullRequest(ghPullRequest gh.PullRequest) model.PullRequest {
	modelPullRequest := model.PullRequest{}

	modelPullRequest.Additions = ghPullRequest.Additions
	modelPullRequest.Closed = ghPullRequest.Closed
	modelPullRequest.CreatedAt = ghPullRequest.CreatedAt
	modelPullRequest.Deletions = ghPullRequest.Deletions
	modelPullRequest.IsDraft = ghPullRequest.IsDraft
	modelPullRequest.Merged = ghPullRequest.Merged
	modelPullRequest.Permalink = ghPullRequest.Permalink
	modelPullRequest.Title = ghPullRequest.Title

	switch ghPullRequest.Mergeable {
	case githubv4.MergeableStateConflicting:
		modelPullRequest.Mergeable = model.MergeableStateConflicting
	case githubv4.MergeableStateMergeable:
		modelPullRequest.Mergeable = model.MergeableStateMergeable
	case githubv4.MergeableStateUnknown:
		modelPullRequest.Mergeable = model.MergeableStateUnknown
	}

	return modelPullRequest
}

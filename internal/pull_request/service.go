package pull_request

import (
	"time"

	"github.com/devenjarvis/pr-staton/internal/model"
	"github.com/devenjarvis/pr-staton/pkg/gh"
)

type Api interface {
	GetOpenPRs(string, string) []gh.PullRequest
}

func NewService(ghApi Api) *service {
	return &service{ghApi: ghApi}
}

type service struct {
	ghApi Api
}

func (s service) GetReviewablePRs(repo_owner string, repo_name string) []model.PullRequest {
	openPRs := s.ghApi.GetOpenPRs(repo_owner, repo_name)
	reviewablePRs := []model.PullRequest{}
	minOpen := time.Now().Add(-24 * time.Hour)

	for _, pr := range openPRs {
		if pr.CreatedAt.Before(minOpen) {
			internalPR := pullRequestFromGhPullRequest(pr)
			reviewablePRs = append(reviewablePRs, internalPR)
		}
	}

	return reviewablePRs

}

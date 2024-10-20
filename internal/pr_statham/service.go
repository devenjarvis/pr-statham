package pr_statham

import (
	"fmt"

	"github.com/devenjarvis/pr-staton/internal/model"
)

type GhService interface {
	GetReviewablePRs(string, string) []model.PullRequest
}

type ChatService interface {
	SendMessage(*model.Message) error
}

func NewService(ghService GhService, chatService ChatService) *service {
	return &service{ghService: ghService, chatService: chatService}
}

type service struct {
	ghService   GhService
	chatService ChatService
}

func (s service) SendReviewablePRs(repoOwner string, repoName string, channel string) error {
	// Get reviewable PRs
	reviewablePRs := s.ghService.GetReviewablePRs(repoOwner, repoName)

	// Initialize message
	message := &model.Message{
		// Channel: channel,
		Blocks: []model.Block{
			model.Section{
				Type: "section",
				Text: model.Text{
					Type: "mrkdwn",
					Text: "Good morning! The following PRs are in need of review",
				},
			},
		},
	}

	// Add reviewable PRs to message
	for _, pr := range reviewablePRs {
		reviewMessage := fmt.Sprintf("<%s|%s> [+%d, -%d]", pr.Permalink, pr.Title, pr.Additions, pr.Deletions)
		reviewText := model.Text{
			Type: "mrkdwn",
			Text: reviewMessage,
		}
		reviewBlock := model.Context{
			Type: "context",
			Elements: []model.Element{
				reviewText,
			},
		}
		message.Blocks = append(message.Blocks, reviewBlock)
	}

	// Send message
	return s.chatService.SendMessage(message)
}

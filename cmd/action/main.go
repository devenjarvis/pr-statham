package main

import (
	"os"

	"github.com/devenjarvis/pr-staton/internal/chat"
	"github.com/devenjarvis/pr-staton/internal/prpulse"
	"github.com/devenjarvis/pr-staton/internal/pull_request"
	"github.com/devenjarvis/pr-staton/pkg/gh"
	"github.com/devenjarvis/pr-staton/pkg/slack"
)

func main() {
	// Init GH
	ghAccessToken := os.Getenv("GITHUB_TOKEN")
	ghApi := gh.NewApi(ghAccessToken)
	ghService := pull_request.NewService(ghApi)

	// Init Slack
	webhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	slackApi := slack.NewApi(webhookUrl)
	chatService := chat.NewService(slackApi)

	// Init PRPulse
	prPulseService := prpulse.NewService(ghService, chatService)

	// Send selected notification
	notificationType := "reviewable_prs"

	switch notificationType {
	case "reviewable_prs":
		repoOwner := "slack-go"
		repoName := "slack"
		channel := "channel"
		prPulseService.SendReviewablePRs(repoOwner, repoName, channel)
	}

}

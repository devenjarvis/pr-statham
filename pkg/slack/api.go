package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devenjarvis/pr-staton/internal/model"
)

type slackApi struct {
	webhookUrl string
}

func NewApi(webhookUrl string) *slackApi {
	client := &slackApi{webhookUrl: webhookUrl}

	return client
}

func (api slackApi) SendMessage(message *model.Message) error {
	// Marshal the message
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// fmt.Println(bytes.NewBuffer(jsonData))

	// Create the request
	req, err := http.NewRequest("POST", api.webhookUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Do something with the response?
	fmt.Println(resp)

	// Close the response
	defer resp.Body.Close()

	// No error
	return nil
}

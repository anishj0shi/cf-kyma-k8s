package slack_client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type slackClient struct {
	webhookURL string
}

func (s slackClient) SendMessage(message string) {
	payload := &slackPayload{Text: message}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		panic("Error Marshalling Slack payload")
	}
	res, err := http.Post(s.webhookURL, "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		panic(err)
	}
	log.Print(res.StatusCode)

}

type SlackClient interface {
	SendMessage(message string)
}

func NewSlackClient() SlackClient {
	return &slackClient{
		webhookURL: os.Getenv("SLACK_TOKEN"),
	}
}

type slackPayload struct {
	Text string `json:"text"`
}

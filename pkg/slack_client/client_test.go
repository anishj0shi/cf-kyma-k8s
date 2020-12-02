package slack_client

import "testing"

func TestSlackIntegration(t *testing.T) {
	client := NewSlackClient()
	client.SendMessage("Hello From Test")
}

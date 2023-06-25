package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"taskmine/domain/service"
)

type SlackNotifier struct {
	WebhookURL string
}

type SlackMessage struct {
	Text string `json:"text"`
}

func NewSlackNotifier(url string) service.WebhookNotifier {
	return &SlackNotifier{WebhookURL: url}
}

func (notifier *SlackNotifier) Notify(message string) error {
	slackMessage := SlackMessage{
		Text: message,
	}
	jsonData, err := json.Marshal(slackMessage)
	if err != nil {
		return err
	}

	response, err := http.Post(notifier.WebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("faild to send Slack with json. respnse states code is%d", response.StatusCode)
	}

	return nil
}

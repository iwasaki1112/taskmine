package notifier

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"taskmine/helper"
	"testing"

	"github.com/joho/godotenv"
)

type mockClient struct {
	http.Client
}

var slackWebhookURL string

func TestMain(m *testing.M) {
	path, err := helper.FindRoot()
	if err != nil {
		log.Printf("No get root path")
	}

	err = godotenv.Load(path + "/.env")
	if err != nil {
		log.Printf("No .env file found: %s", path)
	}

	slackWebhookURL = os.Getenv("SLACK_WEBHOOK_URL")
	os.Exit(m.Run())

}

func TestNotifyWithMock(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	notifier := NewSlackNotifier(server.URL)
	err := notifier.Notify("this is test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestNotifyWithSlackServer(t *testing.T) {
	notifier := NewSlackNotifier(slackWebhookURL)
	err := notifier.Notify("this is test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

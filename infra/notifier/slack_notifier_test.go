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
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Printf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	notifier := NewSlackNotifier(server.URL)
	err := notifier.Notify("this is test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

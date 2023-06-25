package service

type WebhookNotifier interface {
	Notify(message string) error
}

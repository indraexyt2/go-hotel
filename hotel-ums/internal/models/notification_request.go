package models

type NotificationRequest struct {
	TemplateName string
	Recipient    string
	Placeholder  map[string]string
}

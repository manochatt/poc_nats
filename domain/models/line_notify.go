package models

const (
	LineNotifyURL = "https://api.line.me/v2/bot/message/push"
)

type LineMessage struct {
	To       string                   `json:"to"`
	Messages []map[string]interface{} `json:"messages"`
}

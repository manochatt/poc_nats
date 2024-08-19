package domain

import (
	"bytes"
	"context"
)

const (
	LineNotifyURL = "https://api.line.me/v2/bot/message/push"
)

type LineNotifyDTO struct {
	To      string                   `json:"to"`
	Message []map[string]interface{} `json:"messages"`
}

type LineNotifyRepository interface {
	SendNotify(c context.Context, payload *bytes.Buffer) error
}

type LineNotifyUsecase interface {
	SendNotify(c context.Context, payload *bytes.Buffer) error
}

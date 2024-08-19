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

type MessageValue struct {
	Title         string `json:"title"`
	Place         string `json:"place"`
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
}

type LineMessageDTO struct {
	ToID         string       `json:"toID"`
	ProjectID    string       `json:"projectID"`
	MessageValue MessageValue `json:"messageValue"`
}

type LineMessage struct {
	To       string                   `json:"to"`
	Messages []map[string]interface{} `json:"messages"`
}

type LineNotifyRepository interface {
	SendNotify(c context.Context, payload *bytes.Buffer) error
	UpdateMessage(c context.Context, payload *[]byte, messageValue MessageValue)
}

type LineNotifyUsecase interface {
	SendNotify(c context.Context, payload *bytes.Buffer) error
	UpdateMessage(c context.Context, payload *[]byte, messageValue MessageValue)
}

// func validateDatetime(sl validator.StructLevel) {
// 	msg := sl.Current().Interface().(MessageValue)

// 	// Parse StartDateTime and EndDateTime
// 	startTime, err1 := time.Parse(time.RFC3339, msg.StartDateTime)
// 	endTime, err2 := time.Parse(time.RFC3339, msg.EndDateTime)

// 	if err1 != nil || err2 != nil {
// 		return
// 	}

// 	// Check if EndDateTime is after StartDateTime
// 	if !endTime.After(startTime) {
// 		sl.ReportError(msg.EndDateTime, "EndDateTime", "endDateTime", "gtfield", "StartDateTime")
// 	}
// }

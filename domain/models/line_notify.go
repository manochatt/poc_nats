package models

import (
	"bytes"
	"context"

	"github.com/manochatt/line-noti/domain/requests"
)

const (
	LineNotifyURL = "https://api.line.me/v2/bot/message/push"
)

type LineMessage struct {
	To       string                   `json:"to"`
	Messages []map[string]interface{} `json:"messages"`
}

type LineNotifyRepository interface {
	SendNotify(c context.Context, payload *bytes.Buffer) error
	UpdateMessage(c context.Context, payload *[]byte, messageValue requests.MessageValue)
}

type LineNotifyUsecase interface {
	SendNotify(c context.Context, payload *bytes.Buffer) error
	UpdateMessage(c context.Context, payload *[]byte, messageValue requests.MessageValue)
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

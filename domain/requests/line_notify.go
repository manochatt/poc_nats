package requests

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LineNotifyRequest struct {
	To      string                   `json:"to"`
	Message []map[string]interface{} `json:"messages"`
}

type MessageValue struct {
	Title         string `json:"title"`
	Place         string `json:"place"`
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
}

func (mv *MessageValue) Validate() error {
	return validation.ValidateStruct(mv,
		validation.Field(&mv.StartDateTime, validation.Required, validation.Date(time.RFC3339).Error("Invalid date time format")),
		validation.Field(&mv.EndDateTime, validation.Required, validation.Date(time.RFC3339).Error("Invalid date time format")),
	)
}

type LineMessageRequest struct {
	ToID         string       `json:"toID"`
	ProjectID    string       `json:"projectID"`
	MessageValue MessageValue `json:"messageValue"`
}

func (r *LineMessageRequest) Validate() error {
	return validation.Errors{
		"ToID":         validation.Validate(&r.ToID, validation.Required),
		"ProjectID":    validation.Validate(&r.ProjectID, validation.Required, is.Hexadecimal),
		"MessageValue": r.MessageValue.Validate(), // Validate the embedded struct
	}.Filter()
}

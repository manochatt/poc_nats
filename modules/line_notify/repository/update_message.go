package repository

import (
	"context"
	"strings"

	"github.com/manochatt/line-noti/domain"
)

func (lnr *lineNotifyRepository) UpdateMessage(c context.Context, payload *[]byte, messageValue domain.MessageValue) {
	replacer := strings.NewReplacer(
		"${Title}", messageValue.Title,
		"${Place}", messageValue.Place,
		"${Time}", messageValue.StartDateTime,
	)

	message := replacer.Replace(string(*payload))
	*payload = []byte(message)
}

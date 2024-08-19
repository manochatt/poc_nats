package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain"
)

func (lnu *lineNotifyUsecase) UpdateMessage(c context.Context, payload *[]byte, messageValue domain.MessageValue) {
	ctx, cancel := context.WithTimeout(c, lnu.contextTimeout)
	defer cancel()
	lnu.lineNotifyRepository.UpdateMessage(ctx, payload, messageValue)
}

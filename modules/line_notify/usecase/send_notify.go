package usecase

import (
	"bytes"
	"context"
)

func (lnu *lineNotifyUsecase) SendNotify(c context.Context, payload *bytes.Buffer) error {
	ctx, cancel := context.WithTimeout(c, lnu.contextTimeout)
	defer cancel()
	return lnu.lineNotifyRepository.SendNotify(ctx, payload)
}

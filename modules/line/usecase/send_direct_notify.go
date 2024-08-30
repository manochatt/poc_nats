package usecase

import (
	"bytes"
	"context"
)

func (lu *lineUsecase) SendDirectNotify(c context.Context, payload *bytes.Buffer) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineRepository.SendNotify(ctx, payload)
}

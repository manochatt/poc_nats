package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain"
)

func (lu *lineTemplateUsecase) FetchByID(c context.Context, id string) ([]domain.LineTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineTemplateRepository.FetchByID(ctx, id)
}

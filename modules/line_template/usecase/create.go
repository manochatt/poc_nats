package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain/models"
)

func (lu *lineTemplateUsecase) Create(c context.Context, lineTemplate *models.LineTemplate) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineTemplateRepository.Create(ctx, lineTemplate)
}

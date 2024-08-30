package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain/models"
)

func (lu *lineUsecase) FindLineTemplateById(c context.Context, id string) ([]models.LineTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineRepository.FindLineTemplateById(ctx, id)
}

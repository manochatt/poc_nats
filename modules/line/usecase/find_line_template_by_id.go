package usecase

import (
	"context"

	line_models "github.com/manochatt/line-noti/domain/line/models"
)

func (lu *lineUsecase) FindLineTemplateById(c context.Context, id string) (line_models.LineTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineRepository.FindLineTemplateById(ctx, id)
}

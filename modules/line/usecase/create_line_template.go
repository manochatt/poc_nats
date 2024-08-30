package usecase

import (
	"context"

	line_models "github.com/manochatt/line-noti/domain/line/models"
)

func (lu *lineUsecase) CreateLineTemplate(c context.Context, lineTemplate *line_models.LineTemplate) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineRepository.CreateLineTemplate(ctx, lineTemplate)
}

package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain/models"
)

func (lu *lineUsecase) CreateLineTemplate(c context.Context, lineTemplate *models.LineTemplate) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineRepository.CreateLineTemplate(ctx, lineTemplate)
}

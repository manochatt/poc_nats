package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain/models"
)

func (lu *lineUsecase) FindLineTemplateByProjectId(c context.Context, projectID string) ([]models.LineTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineRepository.FindLineTemplateByProjectId(ctx, projectID)
}

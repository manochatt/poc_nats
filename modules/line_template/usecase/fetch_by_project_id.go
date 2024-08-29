package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain/models"
)

func (lu *lineTemplateUsecase) FetchByProjectID(c context.Context, projectID string) ([]models.LineTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineTemplateRepository.FetchByProjectID(ctx, projectID)
}

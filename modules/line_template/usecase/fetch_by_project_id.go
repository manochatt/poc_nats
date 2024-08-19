package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain"
)

func (lu *lineTemplateUsecase) FetchByProjectID(c context.Context, projectID string) ([]domain.LineTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineTemplateRepository.FetchByProjectID(ctx, projectID)
}

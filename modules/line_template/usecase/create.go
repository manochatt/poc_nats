package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain"
)

func (lu *lineTemplateUsecase) Create(c context.Context, lineTemplate *domain.LineTemplate) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineTemplateRepository.Create(ctx, lineTemplate)
}

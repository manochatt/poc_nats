package usecase

import (
	"context"

	"github.com/manochatt/line-noti/domain/requests"
)

func (lu *lineUsecase) UpdateLineTemplate(c context.Context, id string, updateRequest *requests.UpdateLineTemplateRequest) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineRepository.UpdateLineTemplate(ctx, id, updateRequest)
}

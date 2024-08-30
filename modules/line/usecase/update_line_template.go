package usecase

import (
	"context"

	line_requests "github.com/manochatt/line-noti/domain/line/requests"
)

func (lu *lineUsecase) UpdateLineTemplate(c context.Context, id string, updateRequest *line_requests.UpdateLineTemplateRequest) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.lineRepository.UpdateLineTemplate(ctx, id, updateRequest)
}

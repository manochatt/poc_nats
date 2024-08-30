package usecase

import (
	"bytes"
	"context"
	"time"

	line_models "github.com/manochatt/line-noti/domain/line/models"
	line_requests "github.com/manochatt/line-noti/domain/line/requests"
	"github.com/manochatt/line-noti/modules/line/repository"
)

type LineUsecase interface {
	CreateLineTemplate(c context.Context, lineTemplate *line_models.LineTemplate) error
	FindLineTemplateById(c context.Context, ID string) (line_models.LineTemplate, error)
	FindLineTemplateByProjectId(c context.Context, projectID string) (line_models.LineTemplate, error)
	UpdateLineTemplate(c context.Context, ID string, updateRequest *line_requests.UpdateLineTemplateRequest) error
	SendDirectNotify(c context.Context, payload *bytes.Buffer) error
	SendMessage(c context.Context, lineMessageRequest line_requests.LineMessageRequest) error
}

type lineUsecase struct {
	lineRepository repository.LineRepository
	contextTimeout time.Duration
}

func NewLineUsecase(lineRepository repository.LineRepository, timeout time.Duration) LineUsecase {
	return &lineUsecase{
		lineRepository: lineRepository,
		contextTimeout: timeout,
	}
}

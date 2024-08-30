package usecase

import (
	"bytes"
	"context"
	"time"

	"github.com/manochatt/line-noti/domain/models"
	"github.com/manochatt/line-noti/domain/requests"
	"github.com/manochatt/line-noti/modules/line/repository"
)

type LineUsecase interface {
	CreateLineTemplate(c context.Context, lineTemplate *models.LineTemplate) error
	FindLineTemplateById(c context.Context, ID string) ([]models.LineTemplate, error)
	FindLineTemplateByProjectId(c context.Context, projectID string) ([]models.LineTemplate, error)
	UpdateLineTemplate(c context.Context, ID string, updateRequest *requests.UpdateLineTemplateRequest) error
	SendDirectNotify(c context.Context, payload *bytes.Buffer) error
	SendMessage(c context.Context, lineMessageRequest requests.LineMessageRequest) error
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

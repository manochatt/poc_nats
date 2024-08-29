package usecase

import (
	"time"

	"github.com/manochatt/line-noti/domain/models"
)

type lineTemplateUsecase struct {
	lineTemplateRepository models.LineTemplateRepository
	contextTimeout         time.Duration
}

func NewLineTemplateUsecase(lineTemplateRepository models.LineTemplateRepository, timeout time.Duration) models.LineTemplateUsecase {
	return &lineTemplateUsecase{
		lineTemplateRepository: lineTemplateRepository,
		contextTimeout:         timeout,
	}
}

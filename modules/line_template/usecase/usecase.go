package usecase

import (
	"time"

	"github.com/manochatt/line-noti/domain"
)

type lineTemplateUsecase struct {
	lineTemplateRepository domain.LineTemplateRepository
	contextTimeout         time.Duration
}

func NewLineTemplateUsecase(lineTemplateRepository domain.LineTemplateRepository, timeout time.Duration) domain.LineTemplateUsecase {
	return &lineTemplateUsecase{
		lineTemplateRepository: lineTemplateRepository,
		contextTimeout:         timeout,
	}
}

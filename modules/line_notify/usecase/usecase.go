package usecase

import (
	"time"

	"github.com/manochatt/line-noti/domain"
)

type lineNotifyUsecase struct {
	lineNotifyRepository domain.LineNotifyRepository
	contextTimeout       time.Duration
}

func NewLineNotifyUsecase(lineNotifyRepository domain.LineNotifyRepository, timeout time.Duration) domain.LineNotifyUsecase {
	return &lineNotifyUsecase{
		lineNotifyRepository: lineNotifyRepository,
		contextTimeout:       timeout,
	}
}

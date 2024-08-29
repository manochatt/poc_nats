package usecase

import (
	"time"

	"github.com/manochatt/line-noti/domain/models"
)

type lineNotifyUsecase struct {
	lineNotifyRepository models.LineNotifyRepository
	contextTimeout       time.Duration
}

func NewLineNotifyUsecase(lineNotifyRepository models.LineNotifyRepository, timeout time.Duration) models.LineNotifyUsecase {
	return &lineNotifyUsecase{
		lineNotifyRepository: lineNotifyRepository,
		contextTimeout:       timeout,
	}
}

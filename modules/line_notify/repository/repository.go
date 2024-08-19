package repository

import (
	"bytes"

	"github.com/manochatt/line-noti/domain"
)

type lineNotifyRepository struct {
	payload *bytes.Buffer
}

func NewLineNotifyRepository() domain.LineNotifyRepository {
	return &lineNotifyRepository{}
}

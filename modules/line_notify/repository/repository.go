package repository

import (
	"bytes"

	"github.com/manochatt/line-noti/domain/models"
)

type lineNotifyRepository struct {
	payload *bytes.Buffer
}

func NewLineNotifyRepository() models.LineNotifyRepository {
	return &lineNotifyRepository{}
}

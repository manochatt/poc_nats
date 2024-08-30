package repository

import (
	"bytes"
	"context"

	"github.com/manochatt/line-noti/domain/models"
	"github.com/manochatt/line-noti/domain/requests"
	"github.com/manochatt/line-noti/mongo"
)

type LineRepository interface {
	CreateLineTemplate(c context.Context, lineTemplate *models.LineTemplate) error
	FindLineTemplateById(c context.Context, ID string) ([]models.LineTemplate, error)
	FindLineTemplateByProjectId(c context.Context, projectID string) ([]models.LineTemplate, error)
	UpdateLineTemplate(c context.Context, ID string, updateRequest *requests.UpdateLineTemplateRequest) error
	SendNotify(c context.Context, payload *bytes.Buffer) error
}

type lineRepository struct {
	database   mongo.Database
	collection string
}

// type lineNotifyRepository struct {
// 	payload *bytes.Buffer
// }

func NewLineRepository(db mongo.Database, collection string) LineRepository {
	return &lineRepository{
		database:   db,
		collection: collection,
	}
}

package repository

import (
	"bytes"
	"context"

	line_models "github.com/manochatt/line-noti/domain/line/models"
	line_requests "github.com/manochatt/line-noti/domain/line/requests"
	"github.com/manochatt/line-noti/mongo"
)

type LineRepository interface {
	CreateLineTemplate(c context.Context, lineTemplate *line_models.LineTemplate) error
	FindLineTemplateById(c context.Context, ID string) ([]line_models.LineTemplate, error)
	FindLineTemplateByProjectId(c context.Context, projectID string) ([]line_models.LineTemplate, error)
	UpdateLineTemplate(c context.Context, ID string, updateRequest *line_requests.UpdateLineTemplateRequest) error
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

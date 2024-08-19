package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLineTemplate = "line_template"
)

type CreateLineTemplateDTO struct {
	ID        primitive.ObjectID       `bson:"_id" json:"-"`
	ProjectID primitive.ObjectID       `bson:"projectID" json:"projectID"`
	Messages  []map[string]interface{} `bson:"messages" json:"messages"`
}

type LineTemplate struct {
	ID        primitive.ObjectID       `bson:"_id"`
	ProjectID primitive.ObjectID       `bson:"projectID"`
	Messages  []map[string]interface{} `bson:"messages"`
}

type LineTemplateRepository interface {
	Create(c context.Context, lineTemplate *LineTemplate) error
	FetchByID(c context.Context, ID string) ([]LineTemplate, error)
	FetchByProjectID(c context.Context, projectID string) ([]LineTemplate, error)
}

type LineTemplateUsecase interface {
	Create(c context.Context, lineTemplate *LineTemplate) error
	FetchByID(c context.Context, ID string) ([]LineTemplate, error)
	FetchByProjectID(c context.Context, projectID string) ([]LineTemplate, error)
}

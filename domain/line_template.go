package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLineTemplate = "line_template"
)

type Message struct {
	Type    string `bson:"type" json:"type"`
	AltText string `bson:"altText" json:"altText"`
}

type LineTemplate struct {
	ID       primitive.ObjectID `bson:"_id" json:"-"`
	ToID     primitive.ObjectID `bson:"toID" json:"toID"`
	Messages []Message          `bson:"messages" json:"messages"`
}

type LineTemplateRepository interface {
	Create(c context.Context, lineTemplate *LineTemplate) error
	FetchByID(c context.Context, ID string) ([]LineTemplate, error)
}

type LineTemplateUsecase interface {
	Create(c context.Context, lineTemplate *LineTemplate) error
	FetchByID(c context.Context, ID string) ([]LineTemplate, error)
}

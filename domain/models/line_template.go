package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLineTemplate = "line_template"
)

type LineTemplate struct {
	ID        primitive.ObjectID       `bson:"_id"`
	ProjectID primitive.ObjectID       `bson:"projectID"`
	Messages  []map[string]interface{} `bson:"messages"`
}

package repository

import (
	"context"

	line_models "github.com/manochatt/line-noti/domain/line/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *lineRepository) FindLineTemplateById(c context.Context, id string) (line_models.LineTemplate, error) {
	collection := lr.database.Collection(lr.collection)

	var lineTemplate line_models.LineTemplate

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return lineTemplate, err
	}

	result := collection.FindOne(c, bson.M{"_id": objId})
	err = result.Decode(&lineTemplate)
	if err != nil {
		return lineTemplate, err
	}

	return lineTemplate, err
}

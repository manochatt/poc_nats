package repository

import (
	"context"

	line_models "github.com/manochatt/line-noti/domain/line/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *lineRepository) FindLineTemplateById(c context.Context, id string) ([]line_models.LineTemplate, error) {
	collection := lr.database.Collection(lr.collection)

	var lineTemplates []line_models.LineTemplate

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return lineTemplates, err
	}

	cursor, err := collection.Find(c, bson.M{"_id": objId})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &lineTemplates)
	if lineTemplates == nil {
		return []line_models.LineTemplate{}, err
	}

	return lineTemplates, err
}

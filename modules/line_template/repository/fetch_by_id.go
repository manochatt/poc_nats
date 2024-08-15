package repository

import (
	"context"

	"github.com/manochatt/line-noti/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *lineTemplateRepository) FetchByID(c context.Context, id string) ([]domain.LineTemplate, error) {
	collection := lr.database.Collection(lr.collection)

	var lineTemplates []domain.LineTemplate

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
		return []domain.LineTemplate{}, err
	}

	return lineTemplates, err
}

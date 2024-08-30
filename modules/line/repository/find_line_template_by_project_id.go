package repository

import (
	"context"

	"github.com/manochatt/line-noti/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *lineRepository) FindLineTemplateByProjectId(c context.Context, projectID string) ([]models.LineTemplate, error) {
	collection := lr.database.Collection(lr.collection)

	var lineTemplates []models.LineTemplate

	objId, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		return lineTemplates, err
	}

	cursor, err := collection.Find(c, bson.M{"projectID": objId})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &lineTemplates)
	if lineTemplates == nil {
		return []models.LineTemplate{}, err
	}

	return lineTemplates, err

}

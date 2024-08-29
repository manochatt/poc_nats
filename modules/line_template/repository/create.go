package repository

import (
	"context"

	"github.com/manochatt/line-noti/domain/models"
)

func (lr *lineTemplateRepository) Create(c context.Context, lineTemplate *models.LineTemplate) error {
	collection := lr.database.Collection(lr.collection)

	_, err := collection.InsertOne(c, lineTemplate)

	return err
}

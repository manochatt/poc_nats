package repository

import (
	"context"

	"github.com/manochatt/line-noti/domain"
)

func (lr *lineTemplateRepository) Create(c context.Context, lineTemplate *domain.LineTemplate) error {
	collection := lr.database.Collection(lr.collection)

	_, err := collection.InsertOne(c, lineTemplate)

	return err
}

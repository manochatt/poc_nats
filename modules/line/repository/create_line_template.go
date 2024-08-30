package repository

import (
	"context"

	line_models "github.com/manochatt/line-noti/domain/line/models"
)

func (lr *lineRepository) CreateLineTemplate(c context.Context, lineTemplate *line_models.LineTemplate) error {
	collection := lr.database.Collection(lr.collection)

	_, err := collection.InsertOne(c, lineTemplate)

	return err
}

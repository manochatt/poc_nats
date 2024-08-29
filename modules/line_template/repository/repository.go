package repository

import (
	"github.com/manochatt/line-noti/domain/models"
	"github.com/manochatt/line-noti/mongo"
)

type lineTemplateRepository struct {
	database   mongo.Database
	collection string
}

func NewLineTemplateRepository(db mongo.Database, collection string) models.LineTemplateRepository {
	return &lineTemplateRepository{
		database:   db,
		collection: collection,
	}
}

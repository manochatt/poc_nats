package repository

import (
	"github.com/manochatt/line-noti/domain"
	"github.com/manochatt/line-noti/mongo"
)

type lineTemplateRepository struct {
	database   mongo.Database
	collection string
}

func NewLineTemplateRepository(db mongo.Database, collection string) domain.LineTemplateRepository {
	return &lineTemplateRepository{
		database:   db,
		collection: collection,
	}
}

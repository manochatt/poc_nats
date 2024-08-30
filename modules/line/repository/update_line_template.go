package repository

import (
	"context"
	"errors"

	line_requests "github.com/manochatt/line-noti/domain/line/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lr *lineRepository) UpdateLineTemplate(c context.Context, id string, updateRequest *line_requests.UpdateLineTemplateRequest) error {
	collection := lr.database.Collection(lr.collection)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updateFields := buildUpdateFields(updateRequest)
	if len(updateFields) == 0 {
		return errors.New("no fields to update")
	}

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": updateFields}

	_, err = collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func buildUpdateFields(updateRequest *line_requests.UpdateLineTemplateRequest) bson.M {
	updateFields := bson.M{}

	if !updateRequest.ProjectID.IsZero() {
		updateFields["projectID"] = updateRequest.ProjectID
	}
	if updateRequest.Messages != nil && len(updateRequest.Messages) > 0 {
		updateFields["messages"] = updateRequest.Messages
	}

	return updateFields
}

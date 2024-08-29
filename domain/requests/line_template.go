package requests

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateLineTemplateRequest struct {
	ID        primitive.ObjectID       `bson:"_id" json:"-"`
	ProjectID primitive.ObjectID       `bson:"projectID" json:"projectID"`
	Messages  []map[string]interface{} `bson:"messages" json:"messages"`
}

func (r *CreateLineTemplateRequest) Validate() {}

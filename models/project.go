package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name,omitempty" json:"name,omitempty"`
	UserID primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
}

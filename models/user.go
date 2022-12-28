package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UID   string             `bson:"uid,omitempty" json:"uid,omitempty"`
	Email string             `bson:"email,omitempty" json:"email,omitempty"`
}

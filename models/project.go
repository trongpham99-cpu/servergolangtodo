package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string               `bson:"name,omitempty" json:"name,omitempty"`
	PhotoURL       string               `bson:"photoURL,omitempty" json:"photoURL,omitempty"`
	UserID         primitive.ObjectID   `bson:"userID,omitempty" json:"userID,omitempty"`
	UserAssignedID []primitive.ObjectID `bson:"userAssignedID,omitempty" json:"userAssignedID,omitempty"`
	CreateAt       time.Time            `json:"created_at" bson:"created_at"`
	UpdateAt       time.Time            `json:"updated_at" bson:"updated_at"`
}

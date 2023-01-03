package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Status      string             `bson:"status,omitempty" json:"status,omitempty"`
	ProjectID   primitive.ObjectID `bson:"projectID,omitempty" json:"projectID,omitempty"`
	UserID      primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	Page        int                `bson:"page ,omitempty" json:"page"`
	Count       int                `bson:"count,omitempty" json:"count"`
	CreateAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdateAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

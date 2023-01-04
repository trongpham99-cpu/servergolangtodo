package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Status      string             `bson:"status,omitempty" json:"status,omitempty"`
	PhotoURL    string             `bson:"photoURL,omitempty" json:"photoURL,omitempty"`
	ProjectID   primitive.ObjectID `bson:"projectID,omitempty" json:"projectID,omitempty"`
	UserID      primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	Page        int                `bson:"page" json:"page"`
	Count       int                `bson:"count" json:"count"`
	Deadline    int64              `json:"deadline" bson:"deadline"`
	FinishedAt  int64              `json:"finished_at" bson:"finished_at"`
	CreateAt    int64              `json:"created_at" bson:"created_at"`
	UpdateAt    int64              `json:"updated_at" bson:"updated_at"`
}

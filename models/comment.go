package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TaskID   primitive.ObjectID `bson:"taskId,omitempty" json:"taskId,omitempty"`
	UserID   primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	Text     string             `bson:"text,omitempty" json:"text,omitempty"`
	CreateAt time.Time          `json:"created_at" bson:"created_at"`
	UpdateAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CommentBucket struct {
	TaskID   primitive.ObjectID `bson:"taskId,omitempty" json:"taskId,omitempty"`
	Comments []Comment          `bson:"comments,omitempty" json:"comments"`
	Page     int                `bson:"page,omitempty" json:"page"`
	Count    int                `bson:"count,omitempty" json:"count"`
	CreateAt time.Time          `json:"created_at" bson:"created_at"`
	UpdateAt time.Time          `json:"updated_at" bson:"updated_at"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoModel struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Title   string             `json:"title,omitempty"`
	Content string             `json:"content,omitempty"`
}

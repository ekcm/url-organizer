package UrlFolder

import "go.mongodb.org/mongo-driver/bson/primitive"

type UrlFolder struct {
	ID         primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	FolderName string               `json:"folderName" bson:"folderName"`
	UrlLinks   []primitive.ObjectID `json:"urlLinks" bson:"urlLinks"`
}

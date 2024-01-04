package models

import (
	"github.com/ekcm/url-organizer/UrlLink/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlFolder struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FolderName string             `json:"folderName" bson:"folderName"`
	UrlLinks   []*models.UrlLink          `json:"urlLinks" bson:"urlLinks"`
}

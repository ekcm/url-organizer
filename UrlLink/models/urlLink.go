package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UrlLink struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Url     string             `json:"url" bson:"url"`
	UrlName string             `json:"urlName" bson:"urlName"`
	UrlType string             `json:"urlType" bson:"urlType"`
}

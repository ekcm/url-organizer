package models

type UrlLink struct {
	UrlId string `json:"urlId" bson:"urlId"`
	Url string `json:"url" bson:"url"`
	UrlName string `json:"urlName" bson:"urlName"`
	UrlType string `json:"urlType" bson:"urlType"`
}
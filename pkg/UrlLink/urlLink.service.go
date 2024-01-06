package UrlLink

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlLinkService interface {
	CreateUrlLink(*UrlLink) error
	GetUrlLink(primitive.ObjectID) (*UrlLink, error)
	GetAll() ([]*UrlLink, error)
	UpdateUrlLink(urlLink *UrlLink) error
	// DeleteUrlLink(urlId string) error
}

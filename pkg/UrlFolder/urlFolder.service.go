package UrlFolder

import (
	"github.com/ekcm/url-organizer/pkg/UrlLink"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlFolderService interface {
	CreateUrlFolder(*UrlFolder) error
	GetUrlFolder(primitive.ObjectID) (*UrlFolder, error)
	AddUrlLink(primitive.ObjectID, primitive.ObjectID) error
	CreateAddUrlLink(objectID primitive.ObjectID, urlLink UrlLink.UrlLink) error
	// GetAllUrlFolder() ([]*UrlFolder, error)
}

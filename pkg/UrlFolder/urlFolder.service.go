package UrlFolder

import "go.mongodb.org/mongo-driver/bson/primitive"

type UrlFolderService interface {
	CreateUrlFolder(*UrlFolder) error
	GetUrlFolder(primitive.ObjectID) (*UrlFolder, error)
	// GetAllUrlFolder() ([]*UrlFolder, error)
}

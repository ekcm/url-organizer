package services

import (
	"github.com/ekcm/url-organizer/UrlLink/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlLinkService interface {
	CreateUrlLink(*models.UrlLink) error
	GetUrlLink(primitive.ObjectID) (*models.UrlLink, error)
	GetAll() ([]*models.UrlLink, error)
	// UpdateUrlLink(urlLink *models.UrlLink) error
	// DeleteUrlLink(urlId string) error
}

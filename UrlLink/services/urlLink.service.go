package services

import "github.com/ekcm/url-organizer/UrlLink/models"

type UrlLinkService interface {
	CreateUrlLink(*models.UrlLink) error
	// GetUrlLink(urlId string) (*models.UrlLink, error)
	// GetUrlLinks() ([]*models.UrlLink, error)
	// UpdateUrlLink(urlLink *models.UrlLink) error
	// DeleteUrlLink(urlId string) error
}

package services

import (
	"context"

	"github.com/ekcm/url-organizer/UrlLink/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlLinkServiceImpl struct {
	urlLinkCollection *mongo.Collection
	ctx               context.Context
}

func NewUrlLinkService(urlLinkCollection *mongo.Collection, ctx context.Context) UrlLinkService {
	return &UrlLinkServiceImpl{
		urlLinkCollection: urlLinkCollection,
		ctx:               ctx,
	}
}

func (u *UrlLinkServiceImpl) CreateUrlLink(url *models.UrlLink) error {
	_, err := u.urlLinkCollection.InsertOne(u.ctx, url)
	return err
}


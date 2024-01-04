package services

import (
	"context"

	"github.com/ekcm/url-organizer/UrlLink/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (u *UrlLinkServiceImpl) GetUrlLink(id primitive.ObjectID) (*models.UrlLink, error) {
	var urlLink *models.UrlLink
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := u.urlLinkCollection.FindOne(u.ctx, query).Decode(&urlLink)
	return urlLink, err
}

func (u *UrlLinkServiceImpl) GetAll() ([]*models.UrlLink, error) {
	var urlLinks []*models.UrlLink
	cursor, err := u.urlLinkCollection.Find(u.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(u.ctx, &urlLinks); err != nil {
		return nil, err
	}
	return urlLinks, nil
}



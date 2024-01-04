package UrlLink

import (
	"context"

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

func (u *UrlLinkServiceImpl) CreateUrlLink(url *UrlLink) error {
	_, err := u.urlLinkCollection.InsertOne(u.ctx, url)
	return err
}

func (u *UrlLinkServiceImpl) GetUrlLink(id primitive.ObjectID) (*UrlLink, error) {
	var urlLink *UrlLink
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := u.urlLinkCollection.FindOne(u.ctx, query).Decode(&urlLink)
	return urlLink, err
}

func (u *UrlLinkServiceImpl) GetAll() ([]*UrlLink, error) {
	var urlLinks []*UrlLink
	cursor, err := u.urlLinkCollection.Find(u.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(u.ctx, &urlLinks); err != nil {
		return nil, err
	}
	return urlLinks, nil
}



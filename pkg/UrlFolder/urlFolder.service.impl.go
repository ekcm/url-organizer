package UrlFolder

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlFolderServiceImpl struct {
	urlFolderCollection *mongo.Collection
	ctx                 context.Context
}

func NewUrlFolderService(urlFolderCollection *mongo.Collection, ctx context.Context) UrlFolderService {
	return &UrlFolderServiceImpl{
		urlFolderCollection: urlFolderCollection,
		ctx:                 ctx,
	}
}

func (u *UrlFolderServiceImpl) CreateUrlFolder(url *UrlFolder) error {
	_, err := u.urlFolderCollection.InsertOne(u.ctx, url)
	return err
}

func (u *UrlFolderServiceImpl) GetUrlFolder(id primitive.ObjectID) (*UrlFolder, error) {
	var urlFolder *UrlFolder
	query := bson.D{bson.E{Key: "_id", Value: id}}
	err := u.urlFolderCollection.FindOne(u.ctx, query).Decode(&urlFolder)
	return urlFolder, err
}


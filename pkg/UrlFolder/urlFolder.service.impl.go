package UrlFolder

import (
	"context"

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


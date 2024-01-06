package UrlFolder

import (
	"context"

	"github.com/ekcm/url-organizer/pkg/UrlLink"
	// "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlFolderServiceImpl struct {
	urlFolderCollection *mongo.Collection
	urlLinkCollection   *mongo.Collection // Add urlLinkCollection field
	ctx                 context.Context
}

func NewUrlFolderService(urlFolderCollection *mongo.Collection, urlLinkCollection *mongo.Collection, ctx context.Context) UrlFolderService {
	return &UrlFolderServiceImpl{
		urlFolderCollection: urlFolderCollection,
		urlLinkCollection:   urlLinkCollection, // Initialize urlLinkCollection field
		ctx:                 ctx,
	}
}

// func NewUrlFolderService(urlFolderCollection *mongo.Collection, urlLinkCollection *mongo.Collection, ctx context.Context) UrlFolderService {
// 	return &UrlFolderServiceImpl{
// 		urlFolderCollection: urlFolderCollection,
// 		urlLinkCollection:   urlLinkCollection, // Initialize urlLinkCollection field
// 		ctx:                 ctx,
// 	}
// }

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

func (u *UrlFolderServiceImpl) AddUrlLink(id primitive.ObjectID, urlId primitive.ObjectID) error {
	query := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "urlLinks", Value: urlId}}}}
	_, err := u.urlFolderCollection.UpdateOne(u.ctx, query, update)
	return err
}

func (u *UrlFolderServiceImpl) CreateAddUrlLink(id primitive.ObjectID, urlLink UrlLink.UrlLink) error {

	// add urlLink to urlLinkCollection
	urlId, err := u.urlLinkCollection.InsertOne(u.ctx, urlLink)
	if err != nil {
		return err
	}

	// call AddUrlLink function
	err = u.AddUrlLink(id, urlId.InsertedID.(primitive.ObjectID))
	return err
}

// func (u *UrlFolderServiceImpl) CreateAddUrlLink(id primitive.ObjectID, urlLink UrlLink.UrlLink) error {
// 	// create urlLink
// 	// var urlLink *UrlLink.UrlLink
// 	// UrlLink := UrlLink.UrlLink{
// 	// 	Url:     "https://www.google.com",
// 	// 	UrlName: "Google",
// 	// 	UrlType: "Search Engine",
// 	// }

// 	// add urlLink to urlLinkCollection
// 	urlId, err := u.urlLinkCollection.InsertOne(u.ctx, urlLink)
// 	if err != nil {
// 		return err
// 	}
	
// 	// adding urlLink ID to urlFolder
// 	// query := bson.D{bson.E{Key: "_id", Value: id}}
// 	// update := bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "urlLinks", Value: urlId}}}}
// 	// _, err := u.urlFolderCollection.UpdateOne(u.ctx, query, update)
// 	// return err

// 	// call AddUrlLink function
// 	err = u.AddUrlLink(id, urlId.InsertedID.(primitive.ObjectID))
// 	return err
// }


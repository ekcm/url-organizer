package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ekcm/url-organizer/pkg/UrlLink"
	"github.com/ekcm/url-organizer/pkg/UrlFolder"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	ctx = context.TODO()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoDBURI := os.Getenv("MONGODB_URI")
	mongoconn := options.Client().ApplyURI(mongoDBURI)
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	// before starting the app, will ping the database to check if it is connected
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	urlLinkCollection = mongoclient.Database("urlLinkdb").Collection("urlLinks")
	urlLinkService = UrlLink.NewUrlLinkService(urlLinkCollection, ctx)
	urlLinkController = UrlLink.New(urlLinkService)

	urlFolderCollection = mongoclient.Database("urlFolderdb").Collection("urlFolders")
	// urlFolderService = UrlFolder.NewUrlFolderService(urlFolderCollection, ctx)
	urlFolderService = UrlFolder.NewUrlFolderService(urlFolderCollection, urlLinkCollection, ctx)
	urlFolderController = UrlFolder.New(urlFolderService)

	// initialize gin server
	server = gin.Default()
}

var (
	server            *gin.Engine
	ctx               context.Context
	mongoclient       *mongo.Client
	err               error
	urlLinkService    UrlLink.UrlLinkService
	urlLinkController UrlLink.UrlLinkController
	urlLinkCollection *mongo.Collection
	urlFolderService		UrlFolder.UrlFolderService
	urlFolderController UrlFolder.UrlFolderController
	urlFolderCollection *mongo.Collection
)

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("v1")
	urlLinkController.RegisterUserRoutes(basepath)
	urlFolderController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}

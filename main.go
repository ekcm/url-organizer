package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ekcm/url-organizer/UrlLink/controllers"
	"github.com/ekcm/url-organizer/UrlLink/services"
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
	urlLinkService = services.NewUrlLinkService(urlLinkCollection, ctx)
	urlLinkController = controllers.New(urlLinkService)

	// initialize gin server
	server = gin.Default()
}

var (
	server            *gin.Engine
	ctx               context.Context
	mongoclient       *mongo.Client
	err               error
	urlLinkService    services.UrlLinkService
	urlLinkController controllers.UrlLinkController
	urlLinkCollection *mongo.Collection
)

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("v1")
	urlLinkController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}
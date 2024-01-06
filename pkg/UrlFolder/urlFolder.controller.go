package UrlFolder

import (
	"net/http"

	"github.com/ekcm/url-organizer/pkg/UrlLink"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlFolderController struct {
	UrlFolderService UrlFolderService
}

func New(urlFolderService UrlFolderService) UrlFolderController {
	return UrlFolderController{
		UrlFolderService: urlFolderService,
	}
}

func (ufc *UrlFolderController) RegisterUserRoutes(router *gin.RouterGroup) {
	urlfolderroute := router.Group("/urlfolder")
	urlfolderroute.POST("/create", ufc.CreateUrlFolder)
	urlfolderroute.GET("/get/:id", ufc.GetUrlFolder)
	urlfolderroute.PUT("/add/:id/:urlid", ufc.AddUrlLink)
	urlfolderroute.PUT("/createAddUrlLink/:id", ufc.CreateAddUrlLink)
	// urlroute.GET("/getall", ufc.GetAllFolder)
}

func (ufc *UrlFolderController) CreateUrlFolder(ctx *gin.Context) {
	var urlFolder UrlFolder

	if err := ctx.ShouldBindJSON(&urlFolder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := ufc.UrlFolderService.CreateUrlFolder(&urlFolder)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "UrlFolder created successfully"})
}

func (ufc *UrlFolderController) GetUrlFolder(ctx *gin.Context) {
	idParam := ctx.Param("id")

	objId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	
	urlFolder, err := ufc.UrlFolderService.GetUrlFolder(objId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": urlFolder})
}

func (ufc *UrlFolderController) AddUrlLink(ctx *gin.Context) {
	idParam := ctx.Param("id")
	urlIdParam := ctx.Param("urlid")

	objId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	urlObjId, err := primitive.ObjectIDFromHex(urlIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err = ufc.UrlFolderService.AddUrlLink(objId, urlObjId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	// uncomment this and comment the below line if you want to receive success message as response instead of updated urlFolder
	urlFolder, err := ufc.UrlFolderService.GetUrlFolder(objId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": urlFolder})
	// ctx.JSON(http.StatusOK, gin.H{"message": "UrlLink added successfully"})
}

func (ufc *UrlFolderController) CreateAddUrlLink(ctx *gin.Context) {
	idParam := ctx.Param("id")

	objId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	// create urlLink object and bind JSON data from the request
	var urlLink UrlLink.UrlLink
	if err := ctx.ShouldBindJSON(&urlLink); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = ufc.UrlFolderService.CreateAddUrlLink(objId, urlLink) 
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	urlFolder, err := ufc.UrlFolderService.GetUrlFolder(objId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": urlFolder})
}



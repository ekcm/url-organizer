package UrlFolder

import (
	"net/http"

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

// func (ufc *UrlFolderController) GetAllFolder(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{"message": "All UrlFolders retrieved successfully"})
// }

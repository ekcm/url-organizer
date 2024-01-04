package controllers

import (
	"net/http"

	"github.com/ekcm/url-organizer/UrlLink/models"
	"github.com/ekcm/url-organizer/UrlLink/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlLinkController struct {
	UrlService services.UrlLinkService
}

func New(urlService services.UrlLinkService) UrlLinkController {
	return UrlLinkController{
		UrlService: urlService,
	}
}

func (uc *UrlLinkController) RegisterUserRoutes(router *gin.RouterGroup) {
	userroute := router.Group("/url")
	userroute.POST("/create", uc.CreateUrlLink)
	userroute.GET("/get/:id", uc.GetUrlLink)
	userroute.GET("/getall", uc.GetAll)
}

func (uc *UrlLinkController) CreateUrlLink(ctx *gin.Context) {
	var urlLink models.UrlLink
	// attempt to bind JSON request body to 'urlLink' variable
	if err := ctx.ShouldBindJSON(&urlLink); err != nil {
		// if there is an error in binding, respond with a bad request status
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// call the CreateUrlLink method of the urlLinkService with the urlLink info
	err := uc.UrlService.CreateUrlLink(&urlLink)

	// check if there was an error during url creation process
	if err != nil {
		// if there is an error, respond with bad gateway status and error message
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	// if the user creation was successful, respond with a success status and a message
	ctx.JSON(http.StatusOK, gin.H{"message": "UrlLink created successfully"})
}

func (uc *UrlLinkController) GetUrlLink(ctx *gin.Context) {
	idParam := ctx.Param("id")

	// convert the string to objectID
	objId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	urlLink, err := uc.UrlService.GetUrlLink(objId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, urlLink)
}

func (uc *UrlLinkController) GetAll(ctx *gin.Context) {
	urlLinks, err := uc.UrlService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, urlLinks)
}

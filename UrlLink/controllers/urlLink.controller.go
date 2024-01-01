package controllers

import (
	"net/http"

	"github.com/ekcm/url-organizer/UrlLink/models"
	"github.com/ekcm/url-organizer/UrlLink/services"
	"github.com/gin-gonic/gin"
)

type UrlLinkController struct {
	UrlService services.UrlLinkService
}

func New(urlService services.UrlLinkService) UrlLinkController {
	return UrlLinkController{
		UrlService: urlService,
	}
}

func (uc *UrlLinkController) RegisterUserRoutes(router *gin.RouterGroup){
	userroute := router.Group("/url")
	userroute.POST("/create", uc.CreateUrlLink)
	// userroute.GET("/:urlId", uc.GetUrlLink)
	// userroute.GET("/", uc.GetUrlLinks)
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


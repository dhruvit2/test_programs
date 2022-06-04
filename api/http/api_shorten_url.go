package http

import (
	"github.com/dhruvit2/url_shortner/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	uriV1       = "/api/v1/url-shortner"
	uriV1CreateUrl = "/create/short-url"
	uriV1GetUrl   = "/get/url"
)

const (
	HOST_URL = "http://localhost:9808/"
)

type V1API APIHandler

func (api *V1API) GetV1Routes(e *gin.Engine) {

	mgmtGrp := e.Group(uriV1)

	mgmtGrp.POST(uriMgtCreateUrl, api.CreateShortenUrl)
	mgmtGrp.GET(uriMgtNotifications, api.GetShortenUrl)
}

func (api *V1API) CreateShortUrl(c *gin.Context) {
	var creationRequest model.UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl, err := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":   err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": HOST_URL + shortUrl,
	})
}

func (api *V1API) HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	if shortUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":   "Invalid shot url",
		})
		return
	}

	initialUrl, err := store.RetrieveInitialUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":   err.Error()})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, initialUrl)
}
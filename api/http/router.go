package http

import (
	urlshrtner "github.com/dhruvit2/url-shortner/shortner"
	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	App urlshrtner.App
}

func NewAPIHandler(app urlshrtner.App) *APIHandler {
	return &APIHandler{
		App: app,
	}
}

func NewRouter(app urlshrtner.App) http.Handler {

	router := gin.Default()

	apiHandler := NewAPIHandler(app)

	v1API := (*V1API)(apiHandler)
	v1API.GetV1Routes(router)

	return router
}
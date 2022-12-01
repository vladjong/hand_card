package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	// userBalance usecase.UserBalanse
}

func New() *handler {
	return &handler{}
}

func (h *handler) NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	return router
}

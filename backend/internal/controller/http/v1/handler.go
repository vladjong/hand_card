package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/vladjong/hand_card/internal/domain/use_cases"
)

type handler struct {
	authUseCase usecases.AuthUseCaser
}

func New(authUseCase usecases.AuthUseCaser) *handler {
	return &handler{
		authUseCase: authUseCase,
	}
}

func (h *handler) NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.UserIdentity)
	{
		api.GET("/cards")
	}

	return router
}

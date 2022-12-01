package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrorUserIdNotFound = errors.New("user id not found")
	ErrorUserIdInvalid  = errors.New("user id is of invalid type")
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, errorResponse{msg})
}

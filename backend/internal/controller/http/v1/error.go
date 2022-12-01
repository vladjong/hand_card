package v1

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, errorResponse{msg})
}

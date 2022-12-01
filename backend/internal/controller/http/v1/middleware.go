package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if len(header) == 0 {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.authUseCase.ParseToken(dto.TokenDto{Token: headerParts[1]})
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, ErrorUserIdNotFound
	}
	idInt, ok := id.(int)
	if !ok {
		return 0, ErrorUserIdInvalid
	}
	return idInt, nil
}

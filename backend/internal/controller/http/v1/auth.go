package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
)

func (h *handler) SignUp(c *gin.Context) {
	var userDto dto.SignUpDto
	if err := c.BindJSON(&userDto); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.authUseCase.CreateUser(userDto); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (h *handler) SignIn(c *gin.Context) {
	var input dto.SingInDto
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//
	// c.JSON(http.StatusOK, dto.Token{token})
}

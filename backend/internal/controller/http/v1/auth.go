package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
)

func (h *handler) SignUp(c *gin.Context) {
	var signUpDto dto.SignUpDto
	if err := c.BindJSON(&signUpDto); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.authUseCase.CreateUser(signUpDto); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (h *handler) SignIn(c *gin.Context) {
	var signInDto dto.SingInDto
	logrus.Infoln(signInDto)
	if err := c.BindJSON(&signInDto); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.authUseCase.GenerateToken(signInDto)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dto.TokenDto{
		Token: token.Name,
	})
}

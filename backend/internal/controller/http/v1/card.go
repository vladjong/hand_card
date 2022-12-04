package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladjong/hand_card/internal/controller/http/v1/dto"
)

func (h *handler) CreateCard(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	var cardDto dto.CardDto
	if err := c.BindJSON(&cardDto); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.cardUseCase.CreateCard(cardDto, userId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (h *handler) GetCards(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	var coordDto dto.Coordinate
	if err := c.BindJSON(&coordDto); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	cards, err := h.cardUseCase.GetCards(userId, coordDto)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, cards)
}

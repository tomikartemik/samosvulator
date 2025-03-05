package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"samosvulator/internal/service"
	"samosvulator/internal/utils"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "No authorization header")
		return
	}

	userId, err := service.ParseToken(header)

	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set("user_id", userId)
}

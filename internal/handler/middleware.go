package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"samosvulator/internal/service"
	"samosvulator/internal/utils"
	"strings"
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

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	userId, err := service.ParseToken(headerParts[1])

	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set("user_id", userId)
}

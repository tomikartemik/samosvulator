package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"samosvulator/internal/model"
)

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Println(message)
	c.AbortWithStatusJSON(statusCode, model.ErrorResponse{Message: message})
}

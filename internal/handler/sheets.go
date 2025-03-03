package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"samosvulator/internal/utils"
)

func (h *Handler) GetRecordsForAnalise(c *gin.Context) {
	records, err := h.services.GetRecordsForAnalise()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, records)
}

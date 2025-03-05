package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"samosvulator/internal/model"
	"samosvulator/internal/utils"
)

func (h *Handler) CreateRecord(c *gin.Context) {
	var record model.Record

	if err := c.ShouldBindJSON(&record); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.CreateRecord(record); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Запись успешно сохранена!"})
}

func (h *Handler) GetAllRecords(c *gin.Context) {
	records, err := h.services.GetAllRecords()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, records)
}

func (h *Handler) GetRecordByUserID(c *gin.Context) {
	id := c.Query("id")
	records, err := h.services.GetRecordsByUserID(id)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, records)
}

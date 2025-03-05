package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"samosvulator/internal/model"
	"samosvulator/internal/utils"
	"strconv"
)

func (h *Handler) CreateRecord(c *gin.Context) {
	var record model.Record

	userIDStr := c.GetString("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&record); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	record.UserID = userID

	if err := h.services.CreateRecord(record); err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
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

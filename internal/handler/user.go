package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"samosvulator/internal/model"
	"samosvulator/internal/utils"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.SignUp(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			utils.NewErrorResponse(c, http.StatusConflict, "Пользователь с таким никнеймом уже существует!")
			return
		}
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Пользователь успешно зарегистрирован!"})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input model.SignInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("handler " + input.Password)
	fmt.Println("hadnler hashed " + utils.GeneratePasswordHash(input.Password))

	user, err := h.services.SignIn(input)
	if err != nil {
		if err.Error() == "Пользователя с таким никнеймом не существует!" {
			utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		} else if err.Error() == "Неверный пароль!" {
			utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) ChangePasswordByMail(c *gin.Context) {
	username := c.Query("username")
	err := h.services.ChangePasswordByMail(username)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK!")
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var input model.ChangePasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userIDStr, exists := c.Get("user_id")
	if !exists {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "user_id not found in context")
		return
	}

	userID, ok := userIDStr.(int)
	if !ok {
		utils.NewErrorResponse(c, http.StatusUnauthorized, "invalid user_id type in context")
		return
	}

	err := h.services.ChangePassword(userID, input.Password, input.NewPassword)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NewErrorResponse(c, http.StatusNotFound, err.Error())
		}
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Password Changed Successfully!")
}

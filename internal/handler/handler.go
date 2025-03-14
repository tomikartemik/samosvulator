package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"samosvulator/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//UPLOADS
	////////////////////////////////////////////////////////////
	router.Static("/uploads", "./uploads")
	////////////////////////////////////////////////////////////

	user := router.Group("/user")
	{
		user.POST("/sign-up", h.SignUp)
		user.POST("/sign-in", h.SignIn)
		user.GET("/change-password", h.ChangePasswordByMail)
	}

	record := router.Group("/record")
	{
		record.GET("/all", h.GetAllRecords)
		record.GET("/analise", h.GetRecordsForAnalise)
	}

	authorized := router.Group("/authorized", h.UserIdentity)
	{
		authorized.POST("/create-record", h.CreateRecord)
		authorized.GET("/records-by-user-id", h.GetRecordByUserID)
		authorized.POST("/change-password", h.ChangePassword)
	}
	return router
}

package handler

import (
	"github.com/dtas-pm/send-task/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("web/**/*")
	router.Static("/style", "./web/style")
	router.Static("/js", "./web/js")

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/sign-in", h.getSignIn)
		auth.GET("/sign-up", h.getSignUp)
	}

	api := router.Group("/api", h.middlewareLogger)
	{
		api.GET("/profile", h.profile)
		api.GET("/disciplines", h.getAllDiscipline)
		api.POST("/disciplines", h.createDiscipline)
		api.GET("/students", h.getStudents)
		api.POST("/students", h.createStudent)
	}

	return router
}

package handler

import (
	"github.com/dtas-pm/send-task/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.Static("/style", "./web/style/*")
	router.Static("/js", "./web/js/*")

	router.Use(LiberalCORS)
	//router.Use(static.Serve("/api", static.LocalFile("web", true)))
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
		disciplines := api.Group("/disciplines")
		{

			disciplines.GET("/discipline", h.getAllDiscipline)
			//disciplines.POST("/discipline", h.createDiscipline)
		}
		api.GET("/disciplines", h.disciplines)
		api.POST("/disciplines", h.createDiscipline)
		api.GET("/students", h.getStudents)
		api.POST("/students", h.createStudent)
	}

	return router
}

// LiberalCORS is a very allowing CORS middleware.
func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}

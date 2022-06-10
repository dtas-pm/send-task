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

	//router.LoadHTMLGlob("web/template/*")
	//router.LoadHTMLFiles("web/template/sign-in.html", "web/template/sign-up.html", "web/template/profile.html")
	//router.Static("/template", "./web/template")
	router.Static("/style", "./web/style")
	router.Static("/js", "./web/js")
	router.Static("/api/style", "./web/style")
	router.Static("/api/js", "./web/js")
	router.Use(LiberalCORS)
	// router.Use(static.Serve("/api", static.LocalFile("web", true)))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/sign-in", h.getSignIn)
		auth.GET("/sign-up", h.getSignUp)
	}

	api := router.Group("/api", h.middlewareLogger)
	{
		adminRoute := api.Group("/admin", h.middlewareRoleAdmin)
		{
			adminRoute.GET("/profile", h.admin)
			students := adminRoute.Group("/students")
			{
				students.GET("/all", h.getStudents)
				students.POST("/", h.createStudent)
				students.POST("/:id", h.deleteStudent)
				students.POST("/update/:id", h.updateStudent)
			}
			adminRoute.GET("/students", h.students)

		}
		teacherRoute := api.Group("/teacher", h.middlewareRoleTeacher)
		{
			teacherRoute.GET("/profile", h.profile)
			disciplines := teacherRoute.Group("/disciplines")
			{

				disciplines.GET("/all", h.getAllDiscipline)
				disciplines.POST("/", h.createDiscipline)
				disciplines.POST("/:id", h.deleteDiscipline)
				//disciplines.POST("/discipline", h.createDiscipline)
			}
			teacherRoute.GET("/disciplines", h.disciplines)
			planDisciplines := teacherRoute.Group("/plan-disciplines")
			{

				planDisciplines.GET("/all", h.getAllPlanDiscipline)
				planDisciplines.POST("/", h.createPlanDiscipline)
				planDisciplines.POST("/:id", h.deletePlanDiscipline)
				planDisciplines.POST("/update/:id", h.updatePlanDiscipline)
				//disciplines.POST("/discipline", h.createDiscipline)
			}
			teacherRoute.GET("/plan-disciplines", h.planDisciplines)
			teacherRoute.GET("/students", h.getStudents)
			teacherRoute.POST("/students", h.createStudent)
		}

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

package handler

import (
	"github.com/dtas-pm/send-task"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	input := &send.User{
		Name:     c.PostForm("name"),
		UserName: c.PostForm("username"),
		Password: c.PostForm("password1"),
		Email:    c.PostForm("email"),
		Role:     "teacher",
	}

	_, err := h.services.Authorization.CreateUser(*input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//c.Request.URL.Path = "/sign-in"
	//h.getSignIn(c)
	c.Redirect(http.StatusSeeOther, "../auth/sign-in")
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	//var input signInInput
	input := signInInput{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}
	//if err := c.BindJSON(&input); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}
	role, token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie(authorizationHeader, token, 3600, "/api", "localhost", false, true)
	c.SetCookie(roleHandler, role, 3600, "/api", "localhost", false, true)
	if role == "admin" {
		c.Redirect(http.StatusSeeOther, "../api/admin/profile")
	} else {
		c.Redirect(http.StatusSeeOther, "../api/teacher/profile")
	}
}

func (h *Handler) getSignIn(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "./web/template/sign-in.html")
}

func (h *Handler) getSignUp(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "./web/template/sign-up.html")
}

//func (h *Handler) adminSignIn(c *gin.Context) {
//	//var input signInInput
//	input := signInInput{
//		Username: c.PostForm("username"),
//		Password: c.PostForm("password"),
//	}
//	//if err := c.BindJSON(&input); err != nil {
//	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
//	//	return
//	//}
//	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	c.SetCookie(authorizationHeader, token, 3600, "/api", "localhost", false, true)
//	c.Redirect(http.StatusSeeOther, "../api/profile")
//	//c.JSON(http.StatusOK, map[string]interface{}{
//	//	"token": token,
//	//})
//}

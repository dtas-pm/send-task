package handler

import (
	"github.com/dtas-pm/send-task"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	input := send.User{
		Name:     c.PostForm("name"),
		UserName: c.PostForm("username"),
		Password: c.PostForm("password"),
	}
	//if err := c.PostForm(&input); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}

	_, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "../api/profile")
	//c.HTML(http.StatusOK, "../web/template/sign-in.html", gin.H{
	//	"id": id,
	//})

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
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie(authorizationHeader, token, 3600, "/api", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "../api/profile")
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"token": token,
	//})
}

func (h *Handler) getSignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "sign-in.html", gin.H{})
}

func (h *Handler) getSignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "sign-up.html", gin.H{})
}

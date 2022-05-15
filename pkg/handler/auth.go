package handler

import (
	"github.com/dtas-pm/send-task"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context){
	var input send.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "web/template/index.html", gin.H{
		"id": id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}


func (h *Handler) getSignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "web/template/index.html", gin.H{})
}


func (h *Handler) getSignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "web/template/index.html", gin.H{})
}
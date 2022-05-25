package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getStudents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "students.html", gin.H{
		"id": userId,
	})
}

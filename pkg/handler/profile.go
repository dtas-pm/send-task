package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) profile(c *gin.Context) {
	// id, _ := c.Get(userCtx)
	c.HTML(http.StatusOK, "profile.html", gin.H{})
}

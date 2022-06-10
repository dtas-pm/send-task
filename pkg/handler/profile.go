package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) profile(c *gin.Context) {
	// id, _ := c.Get(userCtx)
	http.ServeFile(c.Writer, c.Request, "./web/template/profile.html")
}

func (h *Handler) admin(c *gin.Context) {
	// id, _ := c.Get(userCtx)
	http.ServeFile(c.Writer, c.Request, "./web/template/admin.html")
}

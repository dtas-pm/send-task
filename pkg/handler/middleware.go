package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
	roleHandler         = "Role"
	userCtx             = "userId"
)

func (h *Handler) middlewareLogger(c *gin.Context) {
	header, err := c.Cookie(authorizationHeader)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	// parse token
	userId, err := h.services.Authorization.ParseToken(header)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler) middlewareRoleAdmin(c *gin.Context) {
	header, err := c.Cookie(roleHandler)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	if header != "admin" {
		newErrorResponse(c, http.StatusUnauthorized, "no access")
		return
	}
}

func (h *Handler) middlewareRoleTeacher(c *gin.Context) {
	header, err := c.Cookie(roleHandler)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	if header != "teacher" {
		newErrorResponse(c, http.StatusUnauthorized, "no access")
		return
	}
}

//func getUserRole(c *gin.Context) (string, error) {
//	role, ok := c.Get(roleCtx)
//	if !ok {
//		newErrorResponse(c, http.StatusInternalServerError, "user role not found")
//		return "", errors.New("user role not found")
//	}
//
//	roleStr, ok := role.(string)
//	if !ok {
//		newErrorResponse(c, http.StatusInternalServerError, "user role is of invalid type")
//		return "", errors.New("user role not found")
//	}
//
//	return roleStr, nil
//}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

package handler

import (
	"github.com/dtas-pm/send-task"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createDiscipline(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input send.Discipline
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.DisciplineList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllDiscipline(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.DisciplineList.GetAllDiscipline(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "discipline.html", gin.H{
		"discipline": lists,
	})
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"lists": lists,
	//})
}

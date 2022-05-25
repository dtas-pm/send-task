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
	var input = send.Discipline{
		Name:  c.PostForm("new-name-discipline"),
		Group: c.PostForm("new-groups-discipline"),
		Event: send.Event{},
	}

	//if err := c.BindJSON(&input); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}

	_, err = h.services.DisciplineList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusFound, "/api/disciplines")
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
	//c.JSON(http.StatusOK, gin.H{
	//	"lists": lists,
	//})
}

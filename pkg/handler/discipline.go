package handler

import (
	"github.com/dtas-pm/send-task"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) disciplines(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "./web/template/discipline.html")

}

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
	//c.JSON(http.StatusOK, gin.H{})
	c.Redirect(http.StatusFound, "/api/teacher/disciplines")
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
	groups, err := h.services.GroupList.GetAllGroup()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"disciplines": lists,
		"groups":      groups,
	})

	//c.JSON(http.StatusOK, gin.H{
	//	"lists": lists,
	//})
}

func (h *Handler) deleteDiscipline(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//if err := c.BindJSON(&input); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}
	err = h.services.DisciplineList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

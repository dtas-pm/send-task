package handler

import (
	"github.com/dtas-pm/send-task"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) planDisciplines(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "./web/template/plan-discipline.html")

}

func (h *Handler) createPlanDiscipline(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input = send.PlanDiscipline{
		Name:      c.PostForm("new-name-plan-discipline"),
		Group:     c.PostForm("new-group-plan-discipline"),
		Event:     send.EventPD{},
		DateStart: time.Now(),
	}

	input.DateStart, err = time.Parse("2006-01-02", c.PostForm("new-date-plan-discipline"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.services.PlanDisciplineList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//c.JSON(http.StatusOK, gin.H{})
	c.Redirect(http.StatusFound, "/api/teacher/plan-disciplines")
}

func (h *Handler) getAllPlanDiscipline(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.PlanDisciplineList.GetAllPlanDiscipline(userId)
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
		"plan_disciplines": lists,
		"groups":           groups,
	})
}

func (h *Handler) deletePlanDiscipline(c *gin.Context) {
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
	err = h.services.PlanDisciplineList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) updatePlanDiscipline(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input send.PlanDiscipline
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.PlanDisciplineList.Update(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

package handler

import (
	"github.com/dtas-pm/send-task"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getStudents(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.StudentList.GetAllStudent()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "students.html", gin.H{
		"students": lists,
	})
}

func (h *Handler) createStudent(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	//lists, err := h.services.StudentList.GetAllStudent()
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//c.HTML(http.StatusOK, "students.html", gin.H{
	//	"students": lists,
	//})

	var input = send.Student{
		FullName:  c.PostForm("new-name-student"),
		Login:     c.PostForm("new-login-student"),
		Email:     []string{c.PostForm("new-email-student")},
		Group:     c.PostForm("new-group-student"),
		Institute: c.PostForm("new-institute-student"),
	}

	//if err := c.BindJSON(&input); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}
	_, err = h.services.StudentList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusFound, "/api/students")
}

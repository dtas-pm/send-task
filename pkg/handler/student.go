package handler

import (
	"github.com/dtas-pm/send-task"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	//c.HTML(http.StatusOK, "students.html", gin.H{
	//	"students": lists,
	//})
	c.JSON(http.StatusOK, gin.H{
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
	c.Redirect(http.StatusFound, "/api/admin/students")
}

func (h *Handler) students(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "./web/template/students-admin.html")

}

func (h *Handler) deleteStudent(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	//if err := c.BindJSON(&input); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}
	err = h.services.StudentList.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) updateStudent(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input send.Student
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.StudentList.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

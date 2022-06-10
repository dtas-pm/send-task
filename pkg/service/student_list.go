package service

import (
	"github.com/dtas-pm/send-task"
	"github.com/dtas-pm/send-task/pkg/repository"
)

type StudentListService struct {
	repo repository.StudentList
}

func NewStudentListService(repo repository.StudentList) *StudentListService {
	return &StudentListService{repo: repo}
}

func (s *StudentListService) GetAllStudent() ([]send.Student, error) {
	return s.repo.GetAllStudent()
}

func (s *StudentListService) Create(item send.Student) (int, error) {
	return s.repo.Create(item)
}

func (s *StudentListService) Delete(studentId int) error {
	return s.repo.Delete(studentId)
}

func (s *StudentListService) Update(studentId int, input send.Student) error {
	return s.repo.Update(studentId, input)
}

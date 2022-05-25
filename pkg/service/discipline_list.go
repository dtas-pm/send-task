package service

import (
	"github.com/dtas-pm/send-task"
	"github.com/dtas-pm/send-task/pkg/repository"
)

type DisciplineListService struct {
	repo repository.DisciplineList
}

func NewDisciplineListService(repo repository.DisciplineList) *DisciplineListService {
	return &DisciplineListService{repo: repo}
}

func (s *DisciplineListService) Create(userId int, item send.Discipline) (int, error) {
	return s.repo.Create(userId, item)
}

func (s *DisciplineListService) GetAllDiscipline(userId int) ([]send.Discipline, error) {
	return s.repo.GetAllDiscipline(userId)
}

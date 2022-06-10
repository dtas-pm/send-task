package service

import (
	"github.com/dtas-pm/send-task"
	"github.com/dtas-pm/send-task/pkg/repository"
)

type PlanDisciplineListService struct {
	repo repository.PlanDisciplineList
}

func NewPlanDisciplineListService(repo repository.PlanDisciplineList) *PlanDisciplineListService {
	return &PlanDisciplineListService{repo: repo}
}

func (s *PlanDisciplineListService) Create(userId int, item send.PlanDiscipline) (int, error) {
	return s.repo.Create(userId, item)
}

func (s *PlanDisciplineListService) GetAllPlanDiscipline(userId int) ([]send.PlanDiscipline, error) {
	return s.repo.GetAllPlanDiscipline(userId)
}

func (s *PlanDisciplineListService) Delete(userId, disciplineId int) error {
	return s.repo.Delete(userId, disciplineId)
}

func (s *PlanDisciplineListService) Update(userId, disciplineId int, item send.PlanDiscipline) error {
	return s.repo.Update(userId, disciplineId, item)
}

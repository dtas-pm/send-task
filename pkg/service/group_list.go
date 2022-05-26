package service

import (
	"github.com/dtas-pm/send-task"
	"github.com/dtas-pm/send-task/pkg/repository"
)

type GroupListService struct {
	repo repository.GroupList
}

func NewGroupListService(repo repository.GroupList) *GroupListService {
	return &GroupListService{repo: repo}
}

func (s *GroupListService) GetAllGroup() ([]send.Group, error) {
	return s.repo.GetAllGroup()
}

package service

import (
	"github.com/dtas-pm/send-task"
	"github.com/dtas-pm/send-task/pkg/repository"
)

type Authorization interface {
	CreateUser(user send.User) (int, error)
	GenerateToken(username, password string) (string, string, error)
	ParseToken(token string) (int, error)
}

type EndPoint interface {
}

type DisciplineList interface {
	Create(userId int, item send.Discipline) (int, error)
	GetAllDiscipline(userId int) ([]send.Discipline, error)
	Delete(userId, disciplineId int) error
}

type PlanDisciplineList interface {
	Create(userId int, item send.PlanDiscipline) (int, error)
	GetAllPlanDiscipline(userId int) ([]send.PlanDiscipline, error)
	Delete(userId, disciplineId int) error
	Update(userId, disciplineId int, item send.PlanDiscipline) error
}

type StudentList interface {
	GetAllStudent() ([]send.Student, error)
	Create(item send.Student) (int, error)
	Delete(studentId int) error
	Update(studentId int, input send.Student) error
}

type GroupList interface {
	GetAllGroup() ([]send.Group, error)
}

type Service struct {
	Authorization
	EndPoint
	DisciplineList
	PlanDisciplineList
	StudentList
	GroupList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:      NewAuthService(repos.Authorization),
		DisciplineList:     NewDisciplineListService(repos.DisciplineList),
		PlanDisciplineList: NewPlanDisciplineListService(repos.PlanDisciplineList),
		StudentList:        NewStudentListService(repos.StudentList),
		GroupList:          NewGroupListService(repos.GroupList),
	}
}

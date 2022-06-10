package repository

import (
	"github.com/dtas-pm/send-task"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user send.User) (int, error)
	GetUser(username, password string) (send.User, error)
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

type Repository struct {
	Authorization
	EndPoint
	DisciplineList
	PlanDisciplineList
	StudentList
	GroupList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:      NewAuthPostgres(db),
		DisciplineList:     NewDisciplineListPostgres(db),
		PlanDisciplineList: NewPlanDisciplineListPostgres(db),
		StudentList:        NewStudentListPostgres(db),
		GroupList:          NewGroupListPostgres(db),
	}
}

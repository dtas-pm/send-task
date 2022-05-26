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
}

type StudentList interface {
	GetAllStudent() ([]send.Student, error)
	Create(item send.Student) (int, error)
}

type GroupList interface {
	GetAllGroup() ([]send.Group, error)
}

type Repository struct {
	Authorization
	EndPoint
	DisciplineList
	StudentList
	GroupList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		DisciplineList: NewDisciplineListPostgres(db),
		StudentList:    NewStudentListPostgres(db),
		GroupList:      NewGroupListPostgres(db),
	}
}

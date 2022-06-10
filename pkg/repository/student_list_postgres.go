package repository

import (
	"fmt"
	"github.com/dtas-pm/send-task"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type StudentListPostgres struct {
	db *sqlx.DB
}

func NewStudentListPostgres(db *sqlx.DB) *StudentListPostgres {
	return &StudentListPostgres{db: db}
}

func (r *StudentListPostgres) GetAllStudent() ([]send.Student, error) {
	var lists []send.Student
	query := fmt.Sprintf("SELECT id, fullname, login, email, student_group, institute FROM %s",
		studentsTable)
	// err := r.db.Select(&lists, query)
	rows, err := r.db.Query(query)
	if err != nil {
		return []send.Student{}, err
	}
	for rows.Next() {
		var student send.Student
		err = rows.Scan(&student.Id, &student.FullName, &student.Login, (*pq.StringArray)(&student.Email), &student.Group, &student.Institute)
		if err != nil {
			return []send.Student{}, err
		}
		lists = append(lists, student)
	}

	return lists, err
}

func (r *StudentListPostgres) Create(item send.Student) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (fullname, login, email, student_group, institute) values ($1, $2, $3, $4, $5) RETURNING id", studentsTable)
	row := r.db.QueryRow(query, item.FullName, item.Login, pq.Array(item.Email), item.Group, item.Institute)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *StudentListPostgres) Delete(studentId int) error {
	query := fmt.Sprintf("DELETE FROM %s  WHERE id = $1",
		studentsTable)
	_, err := r.db.Exec(query, studentId)
	return err
}

func (r *StudentListPostgres) Update(studentId int, input send.Student) error {
	query := fmt.Sprintf("UPDATE %s SET fullname=$2, login=$3, email=$4, student_group=$5, institute=$6  WHERE id = $1",
		studentsTable)
	_, err := r.db.Exec(query, studentId, input.FullName, input.Login, pq.Array(input.Email), input.Group, input.Institute)
	return err
}

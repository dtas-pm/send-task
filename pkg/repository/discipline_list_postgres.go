package repository

import (
	"fmt"
	"github.com/dtas-pm/send-task"
	"github.com/jmoiron/sqlx"
)

type DisciplineListPostgres struct {
	db *sqlx.DB
}

func NewDisciplineListPostgres(db *sqlx.DB) *DisciplineListPostgres {
	return &DisciplineListPostgres{db: db}
}

func (r *DisciplineListPostgres) Create(userId int, item send.Discipline) (int, error) {
	fmt.Println(item.Event)
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (name, endpoints, groups) values ($1, $2, $3) RETURNING id", disciplineTable)
	row := tx.QueryRow(createListQuery, item.Name, item.Event, item.Group)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersDisciplineQuery := fmt.Sprintf("INSERT INTO %s (discipline_id, users_id) VALUES ($1, $2)", usersDisciplineTable)
	_, err = tx.Exec(createUsersDisciplineQuery, id, userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *DisciplineListPostgres) GetAllDiscipline(userId int) ([]send.Discipline, error) {
	var lists []send.Discipline
	query := fmt.Sprintf("SELECT tl.id, tl.name, tl.endpoints, tl.groups FROM %s tl INNER JOIN %s ul on tl.id = ul.discipline_id WHERE ul.users_id = $1",
		disciplineTable, usersDisciplineTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *DisciplineListPostgres) Delete(userId, disciplineId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.discipline_id AND ul.users_id=$1 AND ul.discipline_id=$2",
		disciplineTable, usersDisciplineTable)
	_, err := r.db.Exec(query, userId, disciplineId)

	return err
}

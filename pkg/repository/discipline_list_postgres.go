package repository

import (
	"fmt"
	"github.com/dtas-pm/send-task"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	row := tx.QueryRow(createListQuery, item.Name, item.Event, pq.Array(item.Groups))
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

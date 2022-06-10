package repository

import (
	"fmt"
	"github.com/dtas-pm/send-task"
	"github.com/jmoiron/sqlx"
)

type PlanDisciplineListPostgres struct {
	db *sqlx.DB
}

func NewPlanDisciplineListPostgres(db *sqlx.DB) *PlanDisciplineListPostgres {
	return &PlanDisciplineListPostgres{db: db}
}

func (r *PlanDisciplineListPostgres) Create(userId int, item send.PlanDiscipline) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (name, date_start, endpoints, groups) values ($1, $2, $3, $4) RETURNING id", planDisciplineTable)
	row := tx.QueryRow(createListQuery, item.Name, item.DateStart, item.Event, item.Group)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersDisciplineQuery := fmt.Sprintf("INSERT INTO %s (plan_discipline_id, users_id) VALUES ($1, $2)", usersPlanDisciplineTable)
	_, err = tx.Exec(createUsersDisciplineQuery, id, userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *PlanDisciplineListPostgres) GetAllPlanDiscipline(userId int) ([]send.PlanDiscipline, error) {
	var lists []send.PlanDiscipline
	query := fmt.Sprintf("SELECT tl.id, tl.name, tl.date_start, tl.endpoints, tl.groups FROM %s tl INNER JOIN %s ul on tl.id = ul.plan_discipline_id WHERE ul.users_id = $1",
		planDisciplineTable, usersPlanDisciplineTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *PlanDisciplineListPostgres) Delete(userId, disciplineId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.plan_discipline_id AND ul.users_id=$1 AND ul.plan_discipline_id=$2",
		planDisciplineTable, usersPlanDisciplineTable)
	_, err := r.db.Exec(query, userId, disciplineId)

	return err
}

func (r *PlanDisciplineListPostgres) Update(userId, disciplineId int, item send.PlanDiscipline) error {
	query := fmt.Sprintf("UPDATE %s tl SET name=$3, date_start=$4, endpoints=$5, groups=$6 FROM %s ul WHERE tl.id = ul.plan_discipline_id AND ul.plan_discipline_id=$1 AND ul.users_id=$2",
		planDisciplineTable, usersPlanDisciplineTable)
	_, err := r.db.Exec(query, disciplineId, userId, item.Name, item.DateStart, item.Event, item.Group)

	return err
}

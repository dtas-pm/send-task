package repository

import (
	"fmt"
	"github.com/dtas-pm/send-task"
	"github.com/jmoiron/sqlx"
)

type GroupListPostgres struct {
	db *sqlx.DB
}

func NewGroupListPostgres(db *sqlx.DB) *GroupListPostgres {
	return &GroupListPostgres{db: db}
}

func (r *GroupListPostgres) GetAllGroup() ([]send.Group, error) {
	var lists []send.Group
	query := fmt.Sprintf("SELECT name  FROM %s ORDER BY name", groupsTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

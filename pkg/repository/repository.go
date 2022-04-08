package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

}

type EndPoint interface {

}

type Repository struct {
	Authorization
	EndPoint
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}


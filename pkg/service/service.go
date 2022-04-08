package service

import "github.com/dtas-pm/send-task/pkg/repository"

type Authorization interface {

}

type EndPoint interface {

}

type Service struct {
	Authorization
	EndPoint
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}


package user

import "time"

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
		timeout: time.Duration(5) * time.Second,
	}
}
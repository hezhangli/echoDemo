package usecase

import (
	"context"
	"echoDemo/user/repository"
)

type usecase struct {
	repo repository.Repository
}

func MakeUseCase(repo repository.Repository) *usecase {
	return &usecase{repo: repo}
}

func NewUsecase(repo repository.Repository) *usecase {
	return &usecase{repo}
}

type Usecase interface {
	CreateUser(ctx context.Context, name, age string, gender bool) error
}

func (uc *usecase) CreateUser(ctx context.Context, name, age string, gender bool) error {
	err := uc.repo.CreateUser(ctx, name, age, gender)

	if err != nil {
		return err
	}
	return nil
}

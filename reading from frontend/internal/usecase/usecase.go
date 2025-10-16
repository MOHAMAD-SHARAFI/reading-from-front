package usecase

import (
	"project/internal/models"
	"project/internal/repository"
)

type Usecase struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (m *Usecase) CreateUser(user models.User) (u *models.User, err error) {
	return m.repo.CreateUser(user)
}

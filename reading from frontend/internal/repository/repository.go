package repository

import (
	"gorm.io/gorm"
	"project/internal/models"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (m *Repository) CreateUser(user models.User) (u *models.User, err error) {
	if err = m.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

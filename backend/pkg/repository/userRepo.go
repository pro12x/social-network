package repository

import "backend/pkg/entity"

type UserRepo interface {
	FindByID(id uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
	Update(user *entity.User) error
}

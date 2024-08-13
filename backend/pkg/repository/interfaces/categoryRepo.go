package interfaces

import "backend/pkg/entity"

type CategoryRepo interface {
	Create(category *entity.Category) error
	Update(category *entity.Category, id uint) error
	Delete(id uint) error
	FindByID(id uint) (*entity.Category, error)
	FindAll() ([]*entity.Category, error)
	FindByName(name string) (*entity.Category, error)
	CountAll() (uint, error)
}

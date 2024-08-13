package service

import "backend/pkg/entity"

type CategoryService interface {
	CreateCategory(category *entity.Category) error
	UpdateCategory(category *entity.Category, id uint) error
	DeleteCategory(id uint) error
	FindCategoryByID(id uint) (*entity.Category, error)
	FindAllCategories() ([]*entity.Category, error)
	FindCategoryByName(name string) (*entity.Category, error)
	CountAllCategories() (uint, error)
}

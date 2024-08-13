package impl

import (
	"backend/pkg/entity"
	"backend/pkg/repository/interfaces"
	"errors"
	"strings"
)

type CategoryServiceImpl struct {
	Repository interfaces.CategoryRepo
}

func (c *CategoryServiceImpl) CreateCategory(category *entity.Category) error {
	isExists, err := c.FindCategoryByName(category.Name)
	if err != nil {
		return errors.New("category already exists: " + err.Error())
	}

	if isExists != nil && strings.ToLower(isExists.Name) == strings.ToLower(category.Name) {
		return errors.New("category already exists")
	}
	return c.Repository.Create(category)
}

func (c *CategoryServiceImpl) UpdateCategory(category *entity.Category, id uint) error {
	isExists, err := c.Repository.FindByID(id)
	if err != nil {
		return errors.New("error updating category: " + err.Error())
	}

	if isExists == nil {
		return errors.New("cannot update non-existing category")
	}

	if category.ID != isExists.ID {
		return errors.New("cannot update category with different ID")
	}
	return c.Repository.Update(category, id)
}

func (c *CategoryServiceImpl) DeleteCategory(id uint) error {
	isExists, err := c.Repository.FindByID(id)
	if err != nil {
		return errors.New("error deleting category: " + err.Error())
	}

	if isExists == nil {
		return errors.New("cannot delete non-existing category")
	}

	return c.Repository.Delete(id)
}

func (c *CategoryServiceImpl) FindCategoryByID(id uint) (*entity.Category, error) {
	return c.Repository.FindByID(id)
}

func (c *CategoryServiceImpl) FindAllCategories() ([]*entity.Category, error) {
	return c.Repository.FindAll()
}

func (c *CategoryServiceImpl) FindCategoryByName(name string) (*entity.Category, error) {
	return c.Repository.FindByName(name)
}

func (c *CategoryServiceImpl) CountAllCategories() (uint, error) {
	return c.Repository.CountAll()
}

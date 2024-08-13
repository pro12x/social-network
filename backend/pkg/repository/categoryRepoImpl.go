package repository

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/entity"
	"database/sql"
	"errors"
)

type CategoryRepoImpl struct {
	db sqlite.Database
}

func NewCategoryRepoImpl(db sqlite.Database) *CategoryRepoImpl {
	return &CategoryRepoImpl{db: db}
}

func (c *CategoryRepoImpl) Create(category *entity.Category) error {
	_, err := c.db.GetDB().Exec(`INSERT INTO categories (name) VALUES (?)`, category.Name)
	if err != nil {
		return errors.New("error saving category")
	}
	return nil
}

func (c *CategoryRepoImpl) Update(category *entity.Category, id uint) error {
	_, err := c.db.GetDB().Exec(`UPDATE categories SET name = ? WHERE id = ?`, category.Name, id)
	if err != nil {
		return errors.New("error updating category")
	}
	return nil
}

func (c *CategoryRepoImpl) Delete(id uint) error {
	_, err := c.db.GetDB().Exec(`DELETE FROM categories WHERE id = ?`, id)
	if err != nil {
		return errors.New("error deleting category")
	}
	return nil
}

func (c *CategoryRepoImpl) FindByID(id uint) (*entity.Category, error) {
	category := new(entity.Category)
	err := c.db.GetDB().QueryRow(`SELECT * FROM categories WHERE id = ?`, id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepoImpl) FindAll() ([]*entity.Category, error) {
	rows, err := c.db.GetDB().Query(`SELECT * FROM categories`)
	if err != nil {
		return nil, errors.New("error fetching categories")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var categories []*entity.Category
	for rows.Next() {
		category := new(entity.Category)
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *CategoryRepoImpl) FindByName(name string) (*entity.Category, error) {
	category := new(entity.Category)
	err := c.db.GetDB().QueryRow(`SELECT * FROM categories WHERE name = ?`, name).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepoImpl) CountAll() (uint, error) {
	var count uint
	err := c.db.GetDB().QueryRow(`SELECT COUNT(*) FROM categories`).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

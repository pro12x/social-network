package interfaces

import "backend/pkg/entity"

type PostRepo interface {
	Create(post *entity.Post) error
	Update(post *entity.Post, id uint) error
	Delete(id uint) error
	FindByID(id uint) (*entity.Post, error)
	FindAll() ([]*entity.Post, error)
	CountAll() (uint, error)
	FindByUserID(userID uint, privacy string) ([]*entity.Post, error)
	FindPostByCategory(categoryID uint) ([]*entity.Post, error)
	// FindPostsByGroupID(groupID uint) ([]*entity.Post, error)
}

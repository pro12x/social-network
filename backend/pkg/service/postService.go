package service

import "backend/pkg/entity"

type PostService interface {
	CreatePost(post *entity.Post) error
	UpdatePost(post *entity.Post, id uint) error
	DeletePost(id uint) error
	FindPostByID(id uint) (*entity.Post, error)
	FindAllPosts() ([]*entity.Post, error)
	CountAllPosts() (uint, error)
	FindPostsByUserID(userID uint, privacy string) ([]*entity.Post, error)
	FindPostByCategory(categoryID uint) ([]*entity.Post, error)
	// FindPostsByGroupID(groupID uint) ([]*entity.Post, error)
}

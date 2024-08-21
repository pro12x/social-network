package interfaces

import "backend/pkg/entity"

type CommentRepo interface {
	Create(comment *entity.Comment) error
	Update(comment *entity.Comment, id uint) error
	Delete(id uint) error
	FindByID(id uint) (*entity.Comment, error)
	FindCommentsByPostID(postID uint) ([]*entity.Comment, error)
	CountAll() (uint, error)
}

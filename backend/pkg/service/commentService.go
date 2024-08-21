package service

import "backend/pkg/entity"

type CommentService interface {
	CreateComment(comment *entity.Comment) error
	UpdateComment(comment *entity.Comment, id uint) error
	DeleteComment(id uint) error
	FindCommentByID(id uint) (*entity.Comment, error)
	FindCommentsByPostID(postID uint) ([]*entity.Comment, error)
	CountAllComments() (uint, error)
}

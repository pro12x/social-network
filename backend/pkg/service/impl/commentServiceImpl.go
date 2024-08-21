package impl

import (
	"backend/pkg/entity"
	"backend/pkg/repository/interfaces"
	"backend/pkg/utils"
	"errors"
)

type CommentServiceImpl struct {
	Repository interfaces.CommentRepo
}

func (c *CommentServiceImpl) CreateComment(comment *entity.Comment) error {
	if !utils.CheckComment(comment) {
		return errors.New("invalid comment")
	}
	return c.Repository.Create(comment)
}

func (c *CommentServiceImpl) UpdateComment(comment *entity.Comment, id uint) error {
	isExists, err := c.Repository.FindByID(id)
	if err != nil {
		return errors.New("error updating comment: " + err.Error())
	}

	if isExists == nil {
		return errors.New("cannot update non-existing comment")
	}

	if comment.ID != isExists.ID {
		return errors.New("cannot update comment with different ID")
	}

	if !utils.CheckComment(comment) {
		return errors.New("invalid comment")
	}

	limit, err := c.Repository.CountAll()
	if err != nil {
		return errors.New("error updating comment: " + err.Error())
	}

	if limit < id {
		return errors.New("comment not found to update")
	}

	return c.Repository.Update(comment, id)
}

func (c *CommentServiceImpl) DeleteComment(id uint) error {
	isExists, err := c.Repository.FindByID(id)
	if err != nil {
		return errors.New("error deleting comment: " + err.Error())
	}

	if isExists == nil {
		return errors.New("cannot delete non-existing comment")
	}

	limit, err := c.Repository.CountAll()
	if err != nil {
		return errors.New("error deleting comment: " + err.Error())
	}

	if limit < id {
		return errors.New("comment not found to delete")
	}
	return c.Repository.Delete(id)
}

func (c *CommentServiceImpl) FindCommentByID(id uint) (*entity.Comment, error) {
	return c.Repository.FindByID(id)
}

func (c *CommentServiceImpl) FindCommentsByPostID(postID uint) ([]*entity.Comment, error) {
	if postID == 0 {
		return nil, errors.New("invalid post ID")
	}

	return c.Repository.FindCommentsByPostID(postID)
}

func (c *CommentServiceImpl) CountAllComments() (uint, error) {
	return c.Repository.CountAll()
}

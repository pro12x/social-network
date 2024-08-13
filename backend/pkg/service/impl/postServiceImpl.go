package impl

import (
	"backend/pkg/entity"
	"backend/pkg/globale"
	"backend/pkg/repository"
	"backend/pkg/repository/interfaces"
	"errors"
	"strings"
	"time"
)

type PostServiceImpl struct {
	Repository interfaces.PostRepo
}

func (p *PostServiceImpl) CreatePost(post *entity.Post) error {
	if post.UserID == 0 {
		return errors.New("user ID is required")
	}

	if strings.TrimSpace(post.Privacy) == "" {
		post.Privacy = "public"
	}

	return p.Repository.Create(post)
}

func (p *PostServiceImpl) UpdatePost(post *entity.Post, id uint) error {
	isExists, err := p.FindPostByID(id)
	if err != nil {
		return errors.New("error updating post")
	}

	if isExists == nil {
		return errors.New("you cannot update a non-existing post")
	}

	if post.ID != id {
		return errors.New("you cannot update a post with a different ID")
	}

	post.UpdatedAt = time.Now()
	return p.Repository.Update(post, id)
}

func (p *PostServiceImpl) DeletePost(id uint) error {
	isExists, err := p.FindPostByID(id)
	if err != nil {
		return errors.New("error deleting post")
	}

	if isExists == nil {
		return errors.New("you cannot delete a non-existing post")
	}

	return p.Repository.Delete(id)
}

func (p *PostServiceImpl) FindPostByID(id uint) (*entity.Post, error) {
	return p.Repository.FindByID(id)
}

func (p *PostServiceImpl) FindAllPosts() ([]*entity.Post, error) {
	return p.Repository.FindAll()
}

func (p *PostServiceImpl) CountAllPosts() (uint, error) {
	return p.Repository.CountAll()
}

func (p *PostServiceImpl) FindPostsByUserID(userID uint, privacy string) ([]*entity.Post, error) {
	return p.Repository.FindByUserID(userID, privacy)
}

func (p *PostServiceImpl) FindPostByCategory(categoryID uint) ([]*entity.Post, error) {
	catRepo := repository.NewCategoryRepoImpl(*globale.DB)
	category, err := catRepo.FindByID(categoryID)
	if err != nil {
		return nil, errors.New("cannot fetch category in post service")
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	return p.Repository.FindPostByCategory(categoryID)
}

package repository

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/entity"
	"backend/pkg/utils"
	"database/sql"
	"errors"
	"time"
)

type PostRepoImpl struct {
	db sqlite.Database
}

func NewPostRepoImpl(db sqlite.Database) *PostRepoImpl {
	return &PostRepoImpl{db}
}

func (p *PostRepoImpl) Create(post *entity.Post) error {
	_, err := p.db.GetDB().Exec(`INSERT INTO posts (title, content, image, privacy, user_id) VALUES (?, ?, ?, ?, ?)`, post.Title, post.Content, post.Image, post.Privacy, post.UserID)
	if err != nil {
		return errors.New("error saving post")
	}
	return nil
}

func (p *PostRepoImpl) Update(post *entity.Post, id uint) error {
	_, err := p.db.GetDB().Exec(`UPDATE posts SET title = ?, content = ?, image = ?, privacy = ?, updated_at = ? WHERE id = ?`, post.Title, post.Content, post.Image, post.Privacy, time.Now(), id)
	if err != nil {
		return errors.New("error updating post")
	}
	return nil
}

func (p *PostRepoImpl) Delete(id uint) error {
	_, err := p.db.GetDB().Exec(`DELETE FROM posts WHERE id = ?`, id)
	if err != nil {
		return errors.New("error deleting post")
	}
	return nil
}

func (p *PostRepoImpl) FindByID(id uint) (*entity.Post, error) {
	post := new(entity.Post)
	err := p.db.GetDB().QueryRow(`SELECT * FROM posts WHERE id = ?`, id).Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Privacy, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostRepoImpl) FindAll() ([]*entity.Post, error) {
	rows, err := p.db.GetDB().Query(`SELECT * FROM posts`)
	if err != nil {
		return nil, errors.New("cannot fetch posts")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerError.Println(err, utils.Reset)
			return
		}
	}(rows)

	var posts []*entity.Post
	for rows.Next() {
		post := new(entity.Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Privacy, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *PostRepoImpl) CountAll() (uint, error) {
	var count uint
	err := p.db.GetDB().QueryRow(`SELECT COUNT(*) FROM posts`).Scan(&count)
	if err != nil {
		return 0, errors.New("error counting posts")
	}
	return count, nil
}

func (p *PostRepoImpl) FindByUserID(userID uint, privacy string) ([]*entity.Post, error) {
	rows, err := p.db.GetDB().Query(`SELECT * FROM posts WHERE user_id = ? AND privacy = ?`, userID, privacy)
	if err != nil {
		return nil, errors.New("error fetching posts")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerError.Println(err, utils.Reset)
			return
		}
	}(rows)

	var posts []*entity.Post
	for rows.Next() {
		post := new(entity.Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Privacy, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, errors.New("error scanning posts")
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *PostRepoImpl) FindPostByCategory(categoryID uint) ([]*entity.Post, error) {
	rows, err := p.db.GetDB().Query(`SELECT p.* FROM posts p JOIN category_post cp ON p.id = cp.post_id WHERE cp.category_id = ?`, categoryID)
	if err != nil {
		return nil, errors.New("error fetching posts")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerError.Println(err, utils.Reset)
			return
		}
	}(rows)

	var posts []*entity.Post
	for rows.Next() {
		post := new(entity.Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Privacy, &post.UserID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, errors.New("error scanning posts")
		}
		posts = append(posts, post)
	}
	return posts, nil
}

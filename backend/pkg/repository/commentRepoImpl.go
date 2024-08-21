package repository

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/entity"
	"database/sql"
)

type CommentRepoImpl struct {
	db sqlite.Database
}

func NewCommentRepoImpl(db sqlite.Database) *CommentRepoImpl {
	return &CommentRepoImpl{db: db}
}

func (c *CommentRepoImpl) Create(comment *entity.Comment) error {
	_, err := c.db.GetDB().Exec(`INSERT INTO comments (content, image, post_id, user_id) VALUES (?, ?, ?, ?)`, comment.Content, comment.Image, comment.PostID, comment.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentRepoImpl) Update(comment *entity.Comment, id uint) error {
	_, err := c.db.GetDB().Exec(`UPDATE comments SET content = ?, image = ?, updated_at = ? WHERE id = ?`, comment.Content, comment.Image, comment.UpdatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentRepoImpl) Delete(id uint) error {
	_, err := c.db.GetDB().Exec(`DELETE FROM comments WHERE id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentRepoImpl) FindByID(id uint) (*entity.Comment, error) {
	comment := new(entity.Comment)
	err := c.db.GetDB().QueryRow(`SELECT * FROM comments WHERE id = ?`, id).Scan(&comment.ID, &comment.Content, &comment.Image, &comment.PostID, &comment.UserID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *CommentRepoImpl) FindCommentsByPostID(postID uint) ([]*entity.Comment, error) {
	rows, err := c.db.GetDB().Query(`SELECT * FROM comments WHERE post_id = ?`, postID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var comments []*entity.Comment
	for rows.Next() {
		comment := new(entity.Comment)
		err := rows.Scan(&comment.ID, &comment.Content, &comment.Image, &comment.PostID, &comment.UserID, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (c *CommentRepoImpl) CountAll() (uint, error) {
	var count uint
	err := c.db.GetDB().QueryRow(`SELECT COUNT(*) FROM comments`).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

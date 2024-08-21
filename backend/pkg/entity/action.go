package entity

type Action struct {
	ID        uint `json:"id" db:"id"`
	Like      bool `json:"like" db:"like"`
	Dislike   bool `json:"dislike" db:"dislike"`
	PostID    uint `json:"post_id" db:"post_id"`
	CommentID uint `json:"comment_id" db:"comment_id"`
	UserID    uint `json:"user_id" db:"user_id"`
}

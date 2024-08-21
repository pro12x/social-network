package entity

type Comment struct {
	ID        uint   `json:"id" db:"id"`
	Content   string `json:"content" db:"content"`
	Image     string `json:"image" db:"image"`
	PostID    uint   `json:"post_id" db:"post_id"`
	UserID    uint   `json:"user_id" db:"user_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

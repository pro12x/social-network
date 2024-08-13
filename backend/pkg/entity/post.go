package entity

import "time"

type Post struct {
	ID        uint      `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Image     string    `json:"image" db:"image"`
	Privacy   string    `json:"privacy" db:"privacy"`
	UserID    uint      `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

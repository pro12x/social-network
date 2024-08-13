package entity

type Category struct {
	ID   uint   `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

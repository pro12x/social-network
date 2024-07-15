package entity

type User struct {
	Id        uint   `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
	Avatar    string `json:"avatar" db:"avatar"`
	Nickname  string `json:"nickname" db:"nickname"`
	AboutMe   string `json:"aboutme" db:"aboutme"`
}

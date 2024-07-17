package dto

type UserDTO struct {
	ID          uint   `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	Firstname   string `json:"firstname" db:"firstname"`
	Lastname    string `json:"lastname" db:"lastname"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Avatar      string `json:"avatar" db:"avatar"`
	Nickname    string `json:"nickname" db:"nickname"`
	AboutMe     string `json:"about_me" db:"about_me"`
	IsPublic    bool   `json:"is_public" db:"is_public"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

type UserConnectionDTO struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

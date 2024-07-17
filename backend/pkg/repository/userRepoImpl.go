package repository

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/entity"
	"time"
)

type UserRepoImpl struct {
	db sqlite.Database
}

// FindByID is a method to find a user by ID
func (u *UserRepoImpl) FindByID(id uint) (*entity.User, error) {
	user := new(entity.User)
	err := u.db.GetDB().QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}

// FindByEmail is a method to find a user by email
func (u *UserRepoImpl) FindByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	err := u.db.GetDB().QueryRow(`SELECT * FROM users WHERE email = ?`, email).Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}

// Save is a method to save a user
func (u *UserRepoImpl) Save(user *entity.User) error {
	_, err := u.db.GetDB().Exec(`INSERT INTO users (email, password, firstname, lastname, avatar, nickname, about_me) VALUES (?, ?, ?, ?, ?, ?, ?)`, user.Email, user.Password, user.Firstname, user.Lastname, user.Avatar, user.Nickname, user.AboutMe)

	return err
}

func (u *UserRepoImpl) Update(user *entity.User) error {
	_, err := u.db.GetDB().Exec(`UPDATE users SET firstname = ?, lastname = ?, avatar = ?, nickname = ?, about_me = ?, updated_at = ? WHERE id = ?`, user.Firstname, user.Lastname, user.Avatar, user.Nickname, user.AboutMe, user.ID, time.Now())

	return err
}

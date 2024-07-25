package repository

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/entity"
	"backend/pkg/session"
	"database/sql"
	"errors"
	"log"
	"time"
)

type UserRepoImpl struct {
	db           sqlite.Database
	sessionStore *session.StoreSessions
}

func NewUserRepoImpl(db sqlite.Database) *UserRepoImpl {
	return &UserRepoImpl{
		db:           db,
		sessionStore: session.NewSessionStore(),
	}
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
	err := u.db.GetDB().QueryRow(`SELECT id, email, password, firstname, lastname, date_of_birth, avatar, nickname, about_me, is_public, created_at, updated_at FROM users WHERE email = ?`, email).Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No user found
		}
		return nil, err // Some error occurred
	}
	return user, nil
}

// Save is a method to save a user
func (u *UserRepoImpl) Save(user *entity.User) error {
	_, err := u.db.GetDB().Exec(`INSERT INTO users (email, password, firstname, lastname, date_of_birth, avatar, nickname, about_me) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, user.Email, user.Password, user.Firstname, user.Lastname, user.DateOfBirth, user.Avatar, user.Nickname, user.AboutMe)
	if err != nil {
		log.Println("Error saving user")
		return err
	}

	return nil
}

func (u *UserRepoImpl) Update(user *entity.User) error {
	_, err := u.db.GetDB().Exec(`UPDATE users SET firstname = ?, lastname = ?, avatar = ?, nickname = ?, about_me = ?, updated_at = ? WHERE id = ?`, user.Firstname, user.Lastname, user.Avatar, user.Nickname, user.AboutMe, user.ID, time.Now())

	return err
}

func (u *UserRepoImpl) Follow(followerID, followingID uint) error {
	_, err := u.db.GetDB().Exec(`INSERT INTO follows (follower_id, following_id, created_at) VALUES (?, ?, ?)`, followerID, followingID, time.Now())

	return err
}

func (u *UserRepoImpl) Unfollow(followerID, followingID uint) error {
	_, err := u.db.GetDB().Exec(`DELETE FROM follows WHERE follower_id = ? AND following_id = ?`, followerID, followingID)

	return err
}

func (u *UserRepoImpl) FindAllUsers() ([]*entity.User, error) {
	rows, err := u.db.GetDB().Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows")
			return
		}
	}(rows)

	var users []*entity.User
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepoImpl) GetFollowers(userID uint) ([]*entity.User, error) {
	rows, err := u.db.GetDB().Query(`SELECT u.* FROM users u JOIN follows f ON u.id = f.follower_id WHERE f.following_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows")
			return
		}
	}(rows)

	var users []*entity.User
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepoImpl) StoreSession(token string, userID uint) {
	u.sessionStore.StoreSession(token, userID)
}

func (u *UserRepoImpl) GetUserID(token string) (uint, bool) {
	return u.sessionStore.GetUserID(token)
}

func (u *UserRepoImpl) ClearSession(token string) {
	u.sessionStore.ClearSession(token)
}

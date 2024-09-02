package repository

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/entity"
	"backend/pkg/session"
	"backend/pkg/utils"
	"database/sql"
	"errors"
	"time"
)

type UserRepoImpl struct {
	db           sqlite.Database
	sessionStore *session.StoreSessions
}

func NewUserRepoImpl(db sqlite.Database) *UserRepoImpl {
	return &UserRepoImpl{db, session.NewSessionStore()}
}

// FindByID is a method to find a user by ID
func (u *UserRepoImpl) FindByID(id uint) (*entity.User, error) {
	user := new(entity.User)
	err := u.db.GetDB().QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)
	user.Password = ""

	return user, err
}

// FindByEmail is a method to find a user by email
func (u *UserRepoImpl) FindByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	err := u.db.GetDB().QueryRow(`SELECT id, email, password, firstname, lastname, date_of_birth, avatar, nickname, about_me, is_public, created_at, updated_at FROM users WHERE email = ?`, email).Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.LoggerInfo.Println(utils.Warn + "No user found" + utils.Reset)
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
		return errors.New("error saving user")
	}

	return nil
}

func (u *UserRepoImpl) Update(user *entity.User) error {
	_, err := u.db.GetDB().Exec(`UPDATE users SET firstname = ?, lastname = ?, avatar = ?, nickname = ?, about_me = ?, updated_at = ? WHERE id = ?`, user.Firstname, user.Lastname, user.Avatar, user.Nickname, user.AboutMe, user.ID, time.Now())

	return err
}

func (u *UserRepoImpl) CountUsers() (uint, error) {
	var count uint
	err := u.db.GetDB().QueryRow(`SELECT COUNT(*) FROM users`).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (u *UserRepoImpl) FindAllUsers() ([]*entity.User, error) {
	rows, err := u.db.GetDB().Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerError.Println("Error closing rows" + utils.Reset)
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
	query := `SELECT u.id, u.email, u.password, u.firstname, u.lastname, u.date_of_birth, u.avatar, u.nickname, u.about_me, u.is_public, u.created_at, u.updated_at FROM users u JOIN follows f ON u.id = f.follower_id WHERE f.followee_id = ? AND f.status = 'pending' order by f.created_at desc`
	rows, err := u.db.GetDB().Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerError.Println(err, utils.Reset)
			return
		}
	}(rows)

	users := make([]*entity.User, 0)
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		user.Password = ""
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepoImpl) GetFollowings(userID uint) ([]*entity.User, error) {
	query := `SELECT u.id, u.email, u.password, u.firstname, u.lastname, u.date_of_birth, u.avatar, u.nickname, u.about_me, u.is_public, u.created_at, u.updated_at FROM users u JOIN follows f ON u.id = f.followee_id WHERE f.follower_id = ? AND f.status = 'pending' order by f.created_at desc`
	rows, err := u.db.GetDB().Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerError.Println(err, utils.Reset)
			return
		}
	}(rows)

	users := make([]*entity.User, 0)
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		user.Password = ""
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepoImpl) GetFollowerCount(userID uint) (uint, error) {
	query := `SELECT COUNT(*) FROM follows WHERE followee_id = ?`
	row := u.db.GetDB().QueryRow(query, userID)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (u *UserRepoImpl) GetFriends(userID uint) ([]*entity.User, error) {
	query := `SELECT u.id, u.email, u.password, u.firstname, u.lastname, u.date_of_birth, u.avatar, u.nickname, u.about_me, u.is_public, u.created_at, u.updated_at FROM users u JOIN follows f ON u.id = f.followee_id WHERE f.follower_id = ? AND f.status = 'accepted'`
	rows, err := u.db.GetDB().Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerError.Println(utils.Error, err, utils.Reset)
			return
		}
	}(rows)

	users := make([]*entity.User, 0)
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.AboutMe, &user.IsPublic, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		user.Password = ""
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepoImpl) GetFriendsCount(userID uint) (uint, error) {
	query := `SELECT COUNT(*) FROM follows WHERE follower_id = ? AND status = 'accepted'`
	row := u.db.GetDB().QueryRow(query, userID)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (u *UserRepoImpl) GetFollowingCount(userID uint) (uint, error) {
	query := `SELECT COUNT(*) FROM follows WHERE follower_id = ?`
	row := u.db.GetDB().QueryRow(query, userID)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
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

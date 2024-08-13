package repository

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/entity"
	"backend/pkg/utils"
	"database/sql"
	"errors"
)

type FollowRepoImpl struct {
	db sqlite.Database
}

func NewFollowRepoImpl(db sqlite.Database) *FollowRepoImpl {
	return &FollowRepoImpl{db}
}

func (f *FollowRepoImpl) CreateFollow(follow *entity.Follow) error {
	query := `INSERT INTO follows (follower_id, followee_id) VALUES (?, ?)`
	_, err := f.db.GetDB().Exec(query, follow.FollowerID, follow.FolloweeID)
	return err
}

func (f *FollowRepoImpl) UpdateFollowStatus(id uint, status string) error {
	query := `UPDATE follows SET status = ? WHERE id = ?`
	_, err := f.db.GetDB().Exec(query, status, id)
	return err
}

func (f *FollowRepoImpl) DeleteFollow(followerID, followeeID uint) error {
	query := `DELETE FROM follows WHERE follower_id = ? AND followee_id = ?`
	_, err := f.db.GetDB().Exec(query, followerID, followeeID)
	return err
}

func (f *FollowRepoImpl) GetFollowers(userID uint) ([]*entity.User, error) {
	query := `SELECT u.id, u.email, u.password, u.firstname, u.lastname, u.date_of_birth, u.avatar, u.nickname, u.about_me, u.is_public, u.created_at, u.updated_at FROM users u JOIN follows f ON u.id = f.follower_id WHERE f.followee_id = ? AND f.status = 'accepted'`
	rows, err := f.db.GetDB().Query(query, userID)
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

func (f *FollowRepoImpl) GetPendingFollowRequest(id uint) ([]*entity.Follow, error) {
	query := `SELECT id, follower_id, followee_id, status, created_at FROM follows WHERE followee_id = ? AND status = 'pending'`
	rows, err := f.db.GetDB().Query(query, id)
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

	var follows []*entity.Follow
	for rows.Next() {
		follow := new(entity.Follow)
		err := rows.Scan(&follow.ID, &follow.FollowerID, &follow.FolloweeID, &follow.Status, &follow.CreatedAt)
		if err != nil {
			return nil, err
		}
		follows = append(follows, follow)
	}

	return follows, nil
}

func (f *FollowRepoImpl) GetFollowings(userID uint) ([]*entity.User, error) {
	query := `SELECT u.id, u.email, u.password, u.firstname, u.lastname, u.date_of_birth, u.avatar, u.nickname, u.about_me, u.is_public, u.created_at, u.updated_at FROM users u JOIN follows f ON u.id = f.followee_id WHERE f.follower_id = ? AND f.status = 'accepted'`
	rows, err := f.db.GetDB().Query(query, userID)
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

func (f *FollowRepoImpl) GetFollowerCount(userID uint) (uint, error) {
	query := `SELECT COUNT(*) FROM follows WHERE followee_id = ? AND status = 'accepted'`
	row := f.db.GetDB().QueryRow(query, userID)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (f *FollowRepoImpl) GetFollowingCount(userID uint) (uint, error) {
	query := `SELECT COUNT(*) FROM follows WHERE follower_id = ? AND status = 'accepted'`
	row := f.db.GetDB().QueryRow(query, userID)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (f *FollowRepoImpl) CountAllFollows() (uint, error) {
	query := `SELECT COUNT(*) FROM follows`
	row := f.db.GetDB().QueryRow(query)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (f *FollowRepoImpl) FindFollow(followerID, followeeID uint) (*entity.Follow, error) {
	follow := new(entity.Follow)
	err := f.db.GetDB().QueryRow(`SELECT id, follower_id, followee_id, status, created_at FROM follows WHERE (follower_id = ? AND followee_id = ?) OR (follower_id = ? AND followee_id = ?)`, followerID, followeeID, followeeID, followerID).Scan(&follow.ID, &follow.FollowerID, &follow.FolloweeID, &follow.Status, &follow.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No user found
		}
		return nil, err // Some error occurred
	}
	return follow, nil
}

func (f *FollowRepoImpl) FindByID(id uint) (*entity.Follow, error) {
	follow := new(entity.Follow)
	err := f.db.GetDB().QueryRow(`SELECT id, follower_id, followee_id, status, created_at FROM follows WHERE id = ?`, id).Scan(&follow.ID, &follow.FollowerID, &follow.FolloweeID, &follow.Status, &follow.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No user found
		}
		return nil, err // Some error occurred
	}
	return follow, nil
}

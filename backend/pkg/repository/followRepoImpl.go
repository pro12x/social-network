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

func (f *FollowRepoImpl) AreFollowing(followerID, followeeID uint) (bool, error) {
	query := `SELECT COUNT(*) FROM follows WHERE follower_id = ? AND followee_id = ? AND status = 'pending'`
	row := f.db.GetDB().QueryRow(query, followerID, followeeID)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (f *FollowRepoImpl) AreWeFriends(userID, friendID uint) (bool, error) {
	query := `SELECT COUNT(*) FROM follows WHERE (follower_id = ? AND followee_id = ?) OR (follower_id = ? AND followee_id = ?) AND status = 'accepted'`
	row := f.db.GetDB().QueryRow(query, userID, friendID, friendID, userID)

	var count uint
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

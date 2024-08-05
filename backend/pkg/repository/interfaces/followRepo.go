package interfaces

import "backend/pkg/entity"

type FollowRepo interface {
	CreateFollow(follow *entity.Follow) error
	UpdateFollowStatus(id uint, status string) error
	DeleteFollow(followerID, followeeID uint) error
	GetFollowers(userID uint) ([]*entity.User, error)
	GetPendingFollowRequest(id uint) ([]*entity.Follow, error)
	GetFollowings(userID uint) ([]*entity.User, error)
	GetFollowerCount(userID uint) (uint, error)
	GetFollowingCount(userID uint) (uint, error)
	CountAllFollows() (uint, error)
	FindFollow(followerID, followeeID uint) (*entity.Follow, error)
	FindByID(id uint) (*entity.Follow, error)
}

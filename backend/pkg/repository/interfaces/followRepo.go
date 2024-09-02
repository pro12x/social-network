package interfaces

import "backend/pkg/entity"

type FollowRepo interface {
	CreateFollow(follow *entity.Follow) error
	UpdateFollowStatus(id uint, status string) error
	DeleteFollow(followerID, followeeID uint) error
	GetPendingFollowRequest(id uint) ([]*entity.Follow, error)
	CountAllFollows() (uint, error)
	FindFollow(followerID, followeeID uint) (*entity.Follow, error)
	FindByID(id uint) (*entity.Follow, error)
	AreFollowing(followerID, followeeID uint) (bool, error)
	AreWeFriends(userID, friendID uint) (bool, error)
}

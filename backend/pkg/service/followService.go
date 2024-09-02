package service

import (
	"backend/pkg/entity"
)

type FollowService interface {
	FollowUser(followerID, followeeID uint) error
	UnfollowUser(followerID, followeeID uint) error
	AcceptFollowRequest(id uint) error
	DeclineFollowRequest(id uint) error
	GetPendingFollowRequest(userID uint) ([]*entity.Follow, error)
	CountAllFollows() (uint, error)
	FindFollow(followerID, followeeID uint) (*entity.Follow, error)
	FindByID(id uint) (*entity.Follow, error)
	AreFollowing(followerID, followeeID uint) (bool, error)
	AreWeFriends(userID, friendID uint) (bool, error)
}

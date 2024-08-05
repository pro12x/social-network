package service

import (
	"backend/pkg/dto"
	"backend/pkg/entity"
)

type FollowService interface {
	FollowUser(followerID, followeeID uint) error
	UnfollowUser(followerID, followeeID uint) error
	AcceptFollowRequest(id uint) error
	DeclineFollowRequest(id uint) error
	GetPendingFollowRequest(userID uint) ([]*entity.Follow, error)
	GetFollowers(userID uint) ([]*dto.UserDTO, error)
	GetFollowings(userID uint) ([]*dto.UserDTO, error)
	GetFollowerCount(userID uint) (uint, error)
	GetFollowingCount(userID uint) (uint, error)
	CountAllFollows() (uint, error)
	FindFollow(followerID, followeeID uint) (*entity.Follow, error)
	FindByID(id uint) (*entity.Follow, error)
}

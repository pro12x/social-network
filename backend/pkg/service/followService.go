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
	GetFollowerCount(userID uint) (int, error)
	GetFollowingCount(userID uint) (int, error)
}

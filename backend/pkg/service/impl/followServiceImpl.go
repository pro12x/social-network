package impl

import (
	"backend/pkg/dto"
	"backend/pkg/entity"
	"backend/pkg/mapper"
	"backend/pkg/repository/interfaces"
	"errors"
)

type FollowServiceImpl struct {
	Repository interfaces.FollowRepo
}

func (f *FollowServiceImpl) FollowUser(followerID, followeeID uint) error {
	// Check if the follow already exists
	isExisted, err := f.Repository.FindFollow(followerID, followeeID)
	if err != nil {
		return err
	}

	if isExisted != nil {
		return errors.New("you already followed this user")
	}

	follow := &entity.Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
	}
	return f.Repository.CreateFollow(follow)
}

func (f *FollowServiceImpl) UnfollowUser(followerID, followeeID uint) error {
	return f.Repository.DeleteFollow(followerID, followeeID)
}

func (f *FollowServiceImpl) AcceptFollowRequest(id uint) error {
	return f.Repository.UpdateFollowStatus(id, "accepted")
}

func (f *FollowServiceImpl) DeclineFollowRequest(id uint) error {
	return f.Repository.UpdateFollowStatus(id, "declined")
}

func (f *FollowServiceImpl) GetPendingFollowRequest(userID uint) ([]*entity.Follow, error) {
	follows, err := f.Repository.GetPendingFollowRequest(userID)
	if err != nil {
		return nil, err
	}
	var followDTOs []*entity.Follow
	for _, follow := range follows {
		followDTOs = append(followDTOs, follow)
	}
	return followDTOs, nil
}

func (f *FollowServiceImpl) GetFollowers(userID uint) ([]*dto.UserDTO, error) {
	users, err := f.Repository.GetFollowers(userID)
	if err != nil {
		return nil, err
	}
	var userDTOs []*dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, mapper.UserToDTO(user))
	}

	return userDTOs, nil
}

func (f *FollowServiceImpl) GetFollowings(userID uint) ([]*dto.UserDTO, error) {
	users, err := f.Repository.GetFollowings(userID)
	if err != nil {
		return nil, err
	}
	var userDTOs []*dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, mapper.UserToDTO(user))
	}
	return userDTOs, nil
}

func (f *FollowServiceImpl) GetFollowerCount(userID uint) (int, error) {
	return f.Repository.GetFollowerCount(userID)
}

func (f *FollowServiceImpl) GetFollowingCount(userID uint) (int, error) {
	return f.Repository.GetFollowingCount(userID)
}

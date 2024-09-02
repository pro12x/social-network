package impl

import (
	"backend/pkg/entity"
	"backend/pkg/repository/interfaces"
	"errors"
)

type FollowServiceImpl struct {
	Repository interfaces.FollowRepo
}

func (f *FollowServiceImpl) FollowUser(followerID, followeeID uint) error {
	// Check if the follow already exists
	isExists, err := f.Repository.FindFollow(followerID, followeeID)
	if err != nil {
		return errors.New("error occurred while checking if follow exists")
	}

	// Check if follow is existing
	if isExists != nil {
		return errors.New("you already followed this user")
	}

	follow := &entity.Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
	}
	return f.Repository.CreateFollow(follow)
}

func (f *FollowServiceImpl) UnfollowUser(followerID, followeeID uint) error {
	isExists, err := f.Repository.FindFollow(followerID, followeeID)
	if err != nil {
		return err
	}

	// Check if the follow exists
	if isExists == nil {
		return errors.New("requested follow not found")
	}

	// Check if the follow is pending
	if isExists.Status == "pending" {
		return errors.New("cannot unfollow a pending follow request")
	}

	return f.Repository.DeleteFollow(followerID, followeeID)
}

func (f *FollowServiceImpl) AcceptFollowRequest(id uint) error {
	isExists, err := f.Repository.FindByID(id)
	if err != nil {
		return err
	}

	if isExists == nil {
		return errors.New("follow request not found")
	}

	if isExists.Status == "accepted" {
		return errors.New("follow request already accepted")
	}

	return f.Repository.UpdateFollowStatus(id, "accepted")
}

func (f *FollowServiceImpl) DeclineFollowRequest(id uint) error {
	isExists, err := f.Repository.FindByID(id)
	if err != nil {
		return err
	}

	if isExists == nil {
		return errors.New("follow request not found")
	}

	if isExists.Status == "rejected" {
		return errors.New("follow request already rejected")
	}

	err = f.Repository.DeleteFollow(isExists.FollowerID, isExists.FolloweeID)
	if err != nil {
		return err
	}

	return f.Repository.UpdateFollowStatus(id, "rejected")
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

func (f *FollowServiceImpl) CountAllFollows() (uint, error) {
	return f.Repository.CountAllFollows()
}

func (f *FollowServiceImpl) FindFollow(followerID, followeeID uint) (*entity.Follow, error) {
	return f.Repository.FindFollow(followerID, followeeID)
}

func (f *FollowServiceImpl) FindByID(id uint) (*entity.Follow, error) {
	return f.Repository.FindByID(id)
}

func (f *FollowServiceImpl) AreFollowing(followerID, followeeID uint) (bool, error) {

	return f.Repository.AreFollowing(followerID, followeeID)
}

func (f *FollowServiceImpl) AreWeFriends(userID, friendID uint) (bool, error) {
	return f.Repository.AreWeFriends(userID, friendID)
}

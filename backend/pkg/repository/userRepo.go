package repository

import "backend/pkg/entity"

type UserRepo interface {
	FindByID(id uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
	Update(user *entity.User) error
	Follow(followerID, followingID uint) error
	Unfollow(followerID, followingID uint) error
	FindAllUsers() ([]*entity.User, error)
	GetFollowers(userID uint) ([]*entity.User, error)
	StoreSession(token string, userID uint)
	GetUserID(token string) (uint, bool)
	ClearSession(token string)
}

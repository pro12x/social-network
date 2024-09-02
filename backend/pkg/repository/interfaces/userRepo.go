package interfaces

import "backend/pkg/entity"

type UserRepo interface {
	FindByID(id uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Save(user *entity.User) error
	Update(user *entity.User) error
	CountUsers() (uint, error)
	FindAllUsers() ([]*entity.User, error)
	GetFollowers(userID uint) ([]*entity.User, error)
	GetFollowings(userID uint) ([]*entity.User, error)
	GetFriends(userID uint) ([]*entity.User, error)
	GetFriendsCount(userID uint) (uint, error)
	GetFollowerCount(userID uint) (uint, error)
	GetFollowingCount(userID uint) (uint, error)
	StoreSession(token string, userID uint)
	GetUserID(token string) (uint, bool)
	ClearSession(token string)
}

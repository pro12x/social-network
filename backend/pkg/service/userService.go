package service

import "backend/pkg/dto"

type UserServcie interface {
	GetUserById(id uint) (*dto.UserDTO, error)
	CreateUser(user *dto.UserDTO) error
	Connection(email, password string) (*dto.UserDTO, error)
	UpdateProfile(id uint, userDTO *dto.UserDTO) error
	CountUsers() (uint, error)
	GetProfile(id uint) (*dto.UserDTO, error)
	CreateSession(user *dto.UserDTO) (string, error)
	IsUserOnline(token string) (bool, error)
	GetAllUsers() ([]*dto.UserDTO, error)
	GetFollowers(userID uint) ([]*dto.UserDTO, error)
	GetFollowings(userID uint) ([]*dto.UserDTO, error)
	GetFriends(userID uint) ([]*dto.UserDTO, error)
	GetFriendsCount(userID uint) (uint, error)
	GetFollowerCount(userID uint) (uint, error)
	GetFollowingCount(userID uint) (uint, error)
}

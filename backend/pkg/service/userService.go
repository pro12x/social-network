package service

import "backend/pkg/dto"

type UserServcie interface {
	GetUserById(id uint) (*dto.UserDTO, error)
	CreateUser(user *dto.UserDTO) error
	Connection(email, password string) (*dto.UserDTO, error)
	UpdateProfile(id uint, userDTO *dto.UserDTO) error
}

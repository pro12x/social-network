package impl

import (
	"backend/pkg/dto"
	"backend/pkg/mapper"
	"backend/pkg/repository"
	"backend/pkg/utils"
	"errors"
)

type UserServiceImpl struct {
	repository repository.UserRepo
}

func (s *UserServiceImpl) GetUserById(id uint) (*dto.UserDTO, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return mapper.UserToDTO(user), nil
}

func (s *UserServiceImpl) CreateUser(user *dto.UserDTO) error {
	isExisted, err := s.repository.FindByEmail(user.Email)
	if isExisted != nil || err == nil {
		return errors.New("user already existed")
	}

	hashedPassword, err := utils.Encrypt(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repository.Save(mapper.DTOToUser(user))
}

func (s *UserServiceImpl) Connection(email, password string) (*dto.UserDTO, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = utils.Compare(user.Password, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return mapper.UserToDTO(user), nil
}

func (s *UserServiceImpl) UpdateProfile(id uint, userDTO *dto.UserDTO) error {
	user := mapper.DTOToUser(userDTO)
	user.ID = id
	return s.repository.Update(user)
}

func (s *UserServiceImpl) GetProfile(id uint) (*dto.UserDTO, error) {
	user, err := s.repository.FindByID(id)

	return mapper.UserToDTO(user), err
}

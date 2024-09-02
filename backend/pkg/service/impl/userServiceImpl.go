package impl

import (
	"backend/pkg/dto"
	"backend/pkg/mapper"
	"backend/pkg/repository/interfaces"
	"backend/pkg/session"
	"backend/pkg/utils"
	"errors"
	"strings"
)

type UserServiceImpl struct {
	Repository interfaces.UserRepo
}

func (s *UserServiceImpl) GetUserById(id uint) (*dto.UserDTO, error) {
	user, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return mapper.UserToDTO(user), nil
}

func (s *UserServiceImpl) CreateUser(user *dto.UserDTO) error {
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" || strings.TrimSpace(user.Firstname) == "" || strings.TrimSpace(user.Lastname) == "" || strings.TrimSpace(user.DateOfBirth) == "" {
		return errors.New("missing required fields")
	}

	isExisted, err := s.Repository.FindByEmail(user.Email)
	if err != nil {
		return err
	}

	if isExisted != nil {
		return errors.New("user already existed")
	}

	hashedPassword, err := utils.Encrypt(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.Repository.Save(mapper.DTOToUser(user))
}

func (s *UserServiceImpl) Connection(email, password string) (*dto.UserDTO, error) {
	user, err := s.Repository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = utils.Compare(user.Password, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	user.Password = ""

	return mapper.UserToDTO(user), nil
}

func (s *UserServiceImpl) GetAllUsers() ([]*dto.UserDTO, error) {
	users, err := s.Repository.FindAllUsers()
	if err != nil {
		return nil, err
	}

	var userDTOs []*dto.UserDTO
	for _, user := range users {
		user.Password = ""
		userDTOs = append(userDTOs, mapper.UserToDTO(user))
	}

	return userDTOs, nil
}

func (s *UserServiceImpl) UpdateProfile(id uint, userDTO *dto.UserDTO) error {
	isExixts, err := s.Repository.FindByID(id)
	if err != nil {
		return err
	}

	if isExixts == nil {
		return errors.New("user not found")
	}

	user := mapper.DTOToUser(userDTO)
	user.ID = id
	return s.Repository.Update(user)
}

func (s *UserServiceImpl) GetProfile(id uint) (*dto.UserDTO, error) {
	user, err := s.Repository.FindByID(id)

	return mapper.UserToDTO(user), err
}

func (s *UserServiceImpl) CountUsers() (uint, error) {
	return s.Repository.CountUsers()
}

func (s *UserServiceImpl) GetFollowers(userID uint) ([]*dto.UserDTO, error) {
	users, err := s.Repository.GetFollowers(userID)
	if err != nil {
		return nil, err
	}
	var userDTOs []*dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, mapper.UserToDTO(user))
	}

	return userDTOs, nil
}

func (s *UserServiceImpl) GetFollowings(userID uint) ([]*dto.UserDTO, error) {
	users, err := s.Repository.GetFollowings(userID)
	if err != nil {
		return nil, err
	}
	var userDTOs []*dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, mapper.UserToDTO(user))
	}
	return userDTOs, nil
}

func (s *UserServiceImpl) GetFriends(userID uint) ([]*dto.UserDTO, error) {
	users, err := s.Repository.GetFriends(userID)
	if err != nil {
		return nil, err
	}
	var userDTOs []*dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, mapper.UserToDTO(user))
	}
	return userDTOs, nil
}

func (s *UserServiceImpl) GetFriendsCount(userID uint) (uint, error) {
	return s.Repository.GetFriendsCount(userID)
}

func (s *UserServiceImpl) GetFollowerCount(userID uint) (uint, error) {
	return s.Repository.GetFollowerCount(userID)
}

func (s *UserServiceImpl) GetFollowingCount(userID uint) (uint, error) {
	return s.Repository.GetFollowingCount(userID)
}

func (s *UserServiceImpl) CreateSession(user *dto.UserDTO) (string, error) {
	return session.CreateSession(*user)
}

func (s *UserServiceImpl) Logout(token string) error {
	err := session.DeleteSession(token)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) IsUserOnline(token string) (bool, error) {
	_, err := session.GetSession(token)
	if err != nil {
		return false, err
	}
	return true, nil
}

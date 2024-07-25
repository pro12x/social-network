package impl

import (
	"backend/pkg/dto"
	"backend/pkg/mapper"
	"backend/pkg/repository"
	"backend/pkg/session"
	"backend/pkg/utils"
	"errors"
)

type UserServiceImpl struct {
	Repository repository.UserRepo
}

func (s *UserServiceImpl) GetUserById(id uint) (*dto.UserDTO, error) {
	user, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return mapper.UserToDTO(user), nil
}

func (s *UserServiceImpl) CreateUser(user *dto.UserDTO) error {
	if user.Email == "" || user.Password == "" || user.Firstname == "" || user.Lastname == "" || user.DateOfBirth == "" {
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

	return mapper.UserToDTO(user), nil
}

func (s *UserServiceImpl) GetAllUsers() ([]*dto.UserDTO, error) {
	users, err := s.Repository.FindAllUsers()
	if err != nil {
		return nil, err
	}

	userDTOs := make([]*dto.UserDTO, len(users))
	for _, user := range users {
		userDTOs = append(userDTOs, mapper.UserToDTO(user))
	}

	return userDTOs, nil
}

func (s *UserServiceImpl) UpdateProfile(id uint, userDTO *dto.UserDTO) error {
	user := mapper.DTOToUser(userDTO)
	user.ID = id
	return s.Repository.Update(user)
}

func (s *UserServiceImpl) GetProfile(id uint) (*dto.UserDTO, error) {
	user, err := s.Repository.FindByID(id)

	return mapper.UserToDTO(user), err
}

func (s *UserServiceImpl) Follow(followerID, followingID uint) error {
	return s.Repository.Follow(followerID, followingID)
}

func (s *UserServiceImpl) Unfollow(followerID, followingID uint) error {
	return s.Repository.Unfollow(followerID, followingID)
}

func (s *UserServiceImpl) GetFollowers(userID uint) ([]*dto.UserDTO, error) {
	users, err := s.Repository.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	userDTOs := make([]*dto.UserDTO, len(users))
	for _, user := range users {
		userDTOs = append(userDTOs, mapper.UserToDTO(user))
	}

	return userDTOs, nil
}

/*func (s *UserServiceImpl) CreateSession(user *dto.UserDTO) (string, error) {
	token, err := utils.GenerateToken()
	if err != nil {
		return "", err
	}
	s.Repository.StoreSession(token, user.ID)
	return token, nil
}*/

func (s *UserServiceImpl) CreateSession(user *dto.UserDTO) (string, error) {
	return session.CreateSession(user.ID)
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

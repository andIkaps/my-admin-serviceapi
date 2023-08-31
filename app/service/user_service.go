package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
	"myadmin/app/request"
	"myadmin/security"
)

type UserService interface {
	List() ([]entity.User, error)
	Insert(req request.RegisterRequest) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	FindById(ID string) (entity.User, error)
	Update(req request.UserUpdateRequest, ID string) (entity.User, error)
	Delete(ID string) (bool, error)
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userServiceImpl{
		repository: r,
	}
}

// @Summary		: List User
// @Description	: Get users from repository
// @Author			: azrielfatur
func (s *userServiceImpl) List() ([]entity.User, error) {
	users, err := s.repository.FindAll()

	if err != nil {
		return users, err
	}

	return users, nil
}

// @Summary		: Insert user
// @Description	: Insert user to repository
// @Author			: azrielfatur
func (s *userServiceImpl) Insert(req request.RegisterRequest) (entity.User, error) {
	hash, err := security.HashPassword(req.Password)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: hash,
	}

	newUser, err := s.repository.Insert(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// @Summary		: Find user
// @Description	: Find user by username
// @Author			: azrielfatur
func (s *userServiceImpl) FindByUsername(username string) (entity.User, error) {
	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary		: Find user
// @Description	: Find user by id
// @Author			: azrielfatur
func (s *userServiceImpl) FindById(ID string) (entity.User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary		: Udate user
// @Description	: Update user to repository
// @Author			: azrielfatur
func (s *userServiceImpl) Update(req request.UserUpdateRequest, ID string) (entity.User, error) {

	user := entity.User{
		Username:  req.Username,
		Name:      req.Name,
		Email:     req.Email,
		UpdatedBy: req.UpdatedBy,
	}

	updateUser, err := s.repository.Update(user, ID)

	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

// @Summary		: Delete user
// @Description	: Delete user to repository
// @Author			: azrielfatur
func (s *userServiceImpl) Delete(ID string) (bool, error) {
	user, err := s.repository.Delete(ID)

	if err != nil {
		return false, nil
	}

	return user, nil
}

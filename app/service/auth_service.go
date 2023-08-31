package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
	"myadmin/app/request"
	"myadmin/security"
)

type AuthService interface {
	Login(req request.LoginRequest) (entity.User, error)
}

type authServiceImpl struct {
	repository repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return &authServiceImpl{
		repository: r,
	}
}

func (s *authServiceImpl) Login(req request.LoginRequest) (entity.User, error) {
	reqUsername := req.Username
	reqPassword := req.Password

	user, err := s.repository.FindByUsername(reqUsername)

	if err != nil {
		return user, err
	}

	checkPassword := security.CheckPassword(user.Password, reqPassword)

	if checkPassword == nil {
		return user, nil
	}

	return user, checkPassword
}

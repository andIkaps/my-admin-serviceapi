package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
	"myadmin/app/request"
)

type UserRoleService interface {
	FindRole(req request.UserRole) (entity.UserRole, error)
	AttachRole(req request.UserHasRole) ([]entity.UserRole, error)
	DetachRole(req request.UserHasRole) (bool, error)
}

type userRoleServiceImpl struct {
	repository repository.UserRoleRepository
}

func NewUserRoleService(r repository.UserRoleRepository) UserRoleService {
	return &userRoleServiceImpl{
		repository: r,
	}
}

//	@Summary		: Find role
//	@Description	: Find user has role
//	@Author			: azrielfatur
func (s *userRoleServiceImpl) FindRole(req request.UserRole) (entity.UserRole, error) {
	userRole := entity.UserRole{
		UserID: req.UserID,
		RoleID: req.RoleID,
	}
	userRole, err := s.repository.Find(userRole)

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

//	@Summary		: Attach role
//	@Description	: Attach user has role
//	@Author			: azrielfatur
func (s *userRoleServiceImpl) AttachRole(req request.UserHasRole) ([]entity.UserRole, error) {
	attach, err := s.repository.Attach(req.Roles)

	if err != nil {
		return attach, err
	}

	return attach, nil

	// return entity.UserRole{}, nil
}

//	@Summary		: Detach role
//	@Description	: Detach user has role
//	@Author			: azrielfatur
func (s *userRoleServiceImpl) DetachRole(req request.UserHasRole) (bool, error) {
	detach, err := s.repository.Detach(req.Roles)

	if err != nil {
		return false, err
	}

	return detach, nil
}

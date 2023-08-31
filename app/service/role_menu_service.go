package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
	"myadmin/app/request"
)

type RoleMenuService interface {
	FindMenu(RoleID string) (entity.UserRole, error)
	AttachMenu(req request.RoleHasMenu) ([]entity.RoleMenu, error)
	DetachMenu(req request.RoleHasMenu, ID string) (bool, error)
}

type roleMenuServiceImpl struct {
	repository repository.RoleMenuRepository
}

func NewRoleMenuService(r repository.RoleMenuRepository) RoleMenuService {
	return &roleMenuServiceImpl{
		repository: r,
	}
}

// @Summary		: Find role
// @Description	: Find user has role
// @Author			: azrielfatur
func (s *roleMenuServiceImpl) FindMenu(RoleID string) (entity.UserRole, error) {
	userRole, err := s.repository.Find(RoleID)

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

// @Summary		: Attach menu
// @Description	: Attach user has menu
// @Author			: azrielfatur
func (s *roleMenuServiceImpl) AttachMenu(req request.RoleHasMenu) ([]entity.RoleMenu, error) {
	attach, err := s.repository.Attach(req.Menus)

	if err != nil {
		return attach, err
	}

	return attach, nil
}

// @Summary		: Detach menu
// @Description	: Detach role has menus
// @Author			: azrielfatur
func (s *roleMenuServiceImpl) DetachMenu(req request.RoleHasMenu, ID string) (bool, error) {
	attach, err := s.repository.Detach(req.Menus, ID)

	if err != nil {
		return attach, err
	}

	return attach, nil
}

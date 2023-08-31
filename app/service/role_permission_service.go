package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
	"myadmin/app/request"
)

type RolePermissionService interface {
	AttachPermission(req request.RoleHasPermission) ([]entity.RolePermission, error)
	DetachPermission(req request.RoleHasPermission) (bool, error)
}

type rolePermissionServiceImpl struct {
	repository repository.RolePermissionRepository
}

func NewRolePermissionService(r repository.RolePermissionRepository) RolePermissionService {
	return &rolePermissionServiceImpl{
		repository: r,
	}
}

// @Summary		: Attach menu
// @Description	: Attach user has menu
// @Author			: azrielfatur
func (s *rolePermissionServiceImpl) AttachPermission(req request.RoleHasPermission) ([]entity.RolePermission, error) {
	attach, err := s.repository.Attach(req.Permission)

	if err != nil {
		return attach, err
	}

	return attach, nil
}

// @Summary		: Detach menu
// @Description	: Detach role has menus
// @Author			: azrielfatur
func (s *rolePermissionServiceImpl) DetachPermission(req request.RoleHasPermission) (bool, error) {
	attach, err := s.repository.Detach(req.Permission)

	if err != nil {
		return attach, err
	}

	return attach, nil
}

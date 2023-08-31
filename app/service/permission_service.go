package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
	"myadmin/app/request"
)

type PermissionService interface {
	List() ([]entity.Permission, error)
	Insert(req request.PermissionRequest) (entity.Permission, error)
	FindById(ID string) (entity.Permission, error)
	Update(req request.PermissionRequest, ID string) (entity.Permission, error)
	Delete(ID string) (bool, error)
}

type permissionServiceImpl struct {
	repository repository.PermissionRepository
}

func NewPermissionService(r repository.PermissionRepository) PermissionService {
	return &permissionServiceImpl{
		repository: r,
	}
}

//	@Summary		: List permission
//	@Description	: Get permissions from repository
//	@Author			: andikaps
func (s *permissionServiceImpl) List() ([]entity.Permission, error) {
	permissions, err := s.repository.FindAll()

	if err != nil {
		return permissions, err
	}

	return permissions, nil
}

//	@Summary		: Insert permission
//	@Description	: insert permission to repository
//	@Author			: andikaps
func (s *permissionServiceImpl) Insert(req request.PermissionRequest) (entity.Permission, error) {

	permission := entity.Permission{
		Name:        req.Name,
		Description: req.Description,
	}

	newPermission, err := s.repository.Insert(permission)

	if err != nil {
		return newPermission, err
	}

	return newPermission, nil
}

//	@Summary		: Find permission
//	@Description	: Find permission by id from repository
//	@Author			: andikaps
func (s *permissionServiceImpl) FindById(ID string) (entity.Permission, error) {
	permission, err := s.repository.FindByID(ID)

	if err != nil {
		return permission, err
	}

	return permission, nil
}

//	@Summary		: Update permission
//	@Description	: Update permission by id from repository
//	@Author			: andikaps
func (s *permissionServiceImpl) Update(req request.PermissionRequest, ID string) (entity.Permission, error) {

	permission := entity.Permission{
		Name:        req.Name,
		Description: req.Description,
	}

	updatePermission, err := s.repository.Update(permission, ID)

	if err != nil {
		return updatePermission, err
	}

	return updatePermission, nil
}

//	@Summary		: Delete permission
//	@Description	: Delete permission to repository
//	@Author			: andikaps
func (s *permissionServiceImpl) Delete(ID string) (bool, error) {
	permission, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return permission, nil
}

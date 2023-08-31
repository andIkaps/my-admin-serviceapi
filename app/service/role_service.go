package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
	"myadmin/app/request"
)

type RoleService interface {
	List() ([]entity.Role, error)
	Insert(req request.RoleRequest) (entity.Role, error)
	FindById(ID string) (entity.Role, error)
	Update(req request.RoleRequest, ID string) (entity.Role, error)
	Delete(ID string) (bool, error)
}

type roleServiceImpl struct {
	repository repository.RoleRepository
}

func NewRoleService(r repository.RoleRepository) RoleService {
	return &roleServiceImpl{
		repository: r,
	}
}

//	@Summary		: List role
//	@Description	: Get roles from repository
//	@Author			: azrielfatur
func (s *roleServiceImpl) List() ([]entity.Role, error) {
	roles, err := s.repository.FindAll()

	if err != nil {
		return roles, err
	}

	return roles, nil
}

//	@Summary		: Insert role
//	@Description	: insert role to repository
//	@Author			: azrielfatur
func (s *roleServiceImpl) Insert(req request.RoleRequest) (entity.Role, error) {

	role := entity.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	newRole, err := s.repository.Insert(role)

	if err != nil {
		return newRole, err
	}

	return newRole, nil
}

//	@Summary		: Find role
//	@Description	: Find role by id from repository
//	@Author			: azrielfatur
func (s *roleServiceImpl) FindById(ID string) (entity.Role, error) {
	role, err := s.repository.FindById(ID)

	if err != nil {
		return role, err
	}

	return role, nil
}

//	@Summary		: Update role
//	@Description	: Update role by id from repository
//	@Author			: azrielfatur
func (s *roleServiceImpl) Update(req request.RoleRequest, ID string) (entity.Role, error) {

	role := entity.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	updateRole, err := s.repository.Update(role, ID)

	if err != nil {
		return updateRole, err
	}

	return updateRole, nil
}

//	@Summary		: Delete role
//	@Description	: Delete role to repository
//	@Author			: azrielfatur
func (s *roleServiceImpl) Delete(ID string) (bool, error) {
	role, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return role, nil
}

package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
)

type PrivilegeService interface {
	List() ([]entity.Privilege, error)
	FindById(ID string) (entity.Privilege, error)
}

type privilegeServiceImpl struct {
	repository repository.PrivilegeRepository
}

func NewPrivilegeService(r repository.PrivilegeRepository) PrivilegeService {
	return &privilegeServiceImpl{
		repository: r,
	}
}

//	@Summary		: List privilege
//	@Description	: Get privileges from repository
//	@Author			: azrielfatur
func (s *privilegeServiceImpl) List() ([]entity.Privilege, error) {
	privileges, err := s.repository.FindAll()

	if err != nil {
		return privileges, err
	}

	return privileges, nil
}

//	@Summary		: Find permission
//	@Description	: Find permission by id from repository
//	@Author			: azrielfatur
func (s *privilegeServiceImpl) FindById(ID string) (entity.Privilege, error) {
	privilege, err := s.repository.FindByID(ID)

	if err != nil {
		return privilege, err
	}

	return privilege, nil
}

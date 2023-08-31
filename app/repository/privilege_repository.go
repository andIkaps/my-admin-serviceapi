package repository

import (
	"myadmin/app/entity"
	"myadmin/config"
)

type PrivilegeRepository interface {
	FindAll() ([]entity.Privilege, error)
	FindByID(ID string) (entity.Privilege, error)
}

type privilegeRepositoryImpl struct {
	config config.Database
}

func NewPrivilegeRepository(database config.Database) PrivilegeRepository {
	return &privilegeRepositoryImpl{
		config: database,
	}
}

//	@Summary		: Get privileges
//	@Description	: -
//	@Author			: andikaps
func (r *privilegeRepositoryImpl) FindAll() ([]entity.Privilege, error) {
	var privileges []entity.Privilege

	err := r.config.DB.Debug().Find(&privileges).Error

	if err != nil {
		return privileges, err
	}

	return privileges, nil
}

//	@Summary		: Get privilege
//	@Description	: Find privilege by ID
//	@Author			: andikaps
func (r *privilegeRepositoryImpl) FindByID(ID string) (entity.Privilege, error) {
	var privilege entity.Privilege

	err := r.config.DB.Debug().Where("id = ?", ID).First(&privilege).Error

	if err != nil {
		return privilege, err
	}

	return privilege, nil
}

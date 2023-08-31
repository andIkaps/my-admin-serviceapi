package repository

import (
	"myadmin/app/entity"
	"myadmin/config"
)

type RoleMenuRepository interface {
	Find(ID string) (entity.UserRole, error)
	Attach(roleMenu []entity.RoleMenu) ([]entity.RoleMenu, error)
	Detach(roleMenu []entity.RoleMenu, ID string) (bool, error)
}

type roleMenuRepositoryImpl struct {
	config config.Database
}

func NewRoleMenuRepository(database config.Database) RoleMenuRepository {
	return &roleMenuRepositoryImpl{
		config: database,
	}
}

//	@Summary		: Find role has menu
//	@Description	: Find role has menu to dataCreatedAt time.Time

func (r *roleMenuRepositoryImpl) Find(RoleID string) (entity.UserRole, error) {
	var userRole entity.UserRole

	err := r.config.DB.Where(entity.UserRole{RoleID: RoleID}).First(&userRole).Error

	if err != nil {
		return userRole, err
	}

	return userRole, nil
}

//	@Summary		: Assign role has menu
//	@Description	: Assign role has menu to dataCreatedAt time.Time

func (r *roleMenuRepositoryImpl) Attach(roleMenu []entity.RoleMenu) ([]entity.RoleMenu, error) {
	err := r.config.DB.Create(&roleMenu).Error

	if err != nil {
		return roleMenu, err
	}

	return roleMenu, nil
}

//	@Summary		: Un-Assign role has menu
//	@Description	: Un-Assign role has menu to dataCreatedAt time.Time

func (r *roleMenuRepositoryImpl) Detach(roleMenu []entity.RoleMenu, ID string) (bool, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Delete(&roleMenu).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

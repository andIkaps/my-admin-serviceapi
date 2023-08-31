package repository

import (
	"myadmin/app/entity"
	"myadmin/config"
)

type PermissionRepository interface {
	Insert(permission entity.Permission) (entity.Permission, error)
	FindAll() ([]entity.Permission, error)
	FindByID(ID string) (entity.Permission, error)
	Update(permission entity.Permission, ID string) (entity.Permission, error)
	Delete(ID string) (bool, error)
}

type permissionRepositoryImpl struct {
	config config.Database
}

func NewPermissionRepository(database config.Database) PermissionRepository {
	return &permissionRepositoryImpl{
		config: database,
	}
}

//	@Summary		: Insert permission
//	@Description	: Insert permission to dataCreatedAt time.Time

func (r *permissionRepositoryImpl) Insert(permission entity.Permission) (entity.Permission, error) {
	err := r.config.DB.Debug().Create(&permission).Error

	if err != nil {
		return permission, err
	}

	return permission, nil
}

//	@Summary		: Get permissions
//	@Description	: -
//	@Author			: andikaps
func (r *permissionRepositoryImpl) FindAll() ([]entity.Permission, error) {
	var permissions []entity.Permission

	err := r.config.DB.Debug().Find(&permissions).Error

	if err != nil {
		return permissions, err
	}

	return permissions, nil
}

//	@Summary		: Get permission
//	@Description	: Find permission by ID
//	@Author			: andikaps
func (r *permissionRepositoryImpl) FindByID(ID string) (entity.Permission, error) {
	var permission entity.Permission

	err := r.config.DB.Debug().Where("id = ?", ID).First(&permission).Error

	if err != nil {
		return permission, err
	}

	return permission, nil
}

//	@Summary		: Update permission
//	@Description	: Update permission by ID
//	@Author			: andikaps
func (r *permissionRepositoryImpl) Update(permission entity.Permission, ID string) (entity.Permission, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&permission).Error

	if err != nil {
		return permission, err
	}

	return permission, nil
}

//	@Summary		: Delete permission
//	@Description	: Delete permission temporary
//	@Author			: andikaps
func (r *permissionRepositoryImpl) Delete(ID string) (bool, error) {
	var permission entity.Permission
	err := r.config.DB.Debug().Where("id = ?", ID).Delete(&permission).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

package repository

import (
	"myadmin/app/entity"
	"myadmin/config"
)

type RolePermissionRepository interface {
	Attach(rolePermission []entity.RolePermission) ([]entity.RolePermission, error)
	Detach(rolePermission []entity.RolePermission) (bool, error)
}

type rolePermissionRepositoryImpl struct {
	config config.Database
}

func NewRolePermissionRepository(database config.Database) RolePermissionRepository {
	return &rolePermissionRepositoryImpl{
		config: database,
	}
}

//	@Summary		: Assign role has menu
//	@Description	: Assign role has menu to dataCreatedAt time.Time

func (r *rolePermissionRepositoryImpl) Attach(rolePermission []entity.RolePermission) ([]entity.RolePermission, error) {
	err := r.config.DB.Debug().Create(&rolePermission).Error

	if err != nil {
		return rolePermission, err
	}

	return rolePermission, nil
}

//	@Summary		: Un-Assign role has menu
//	@Description	: Un-Assign role has menu to dataCreatedAt time.Time

func (r *rolePermissionRepositoryImpl) Detach(rolePermission []entity.RolePermission) (bool, error) {
	permissionID := make([]string, 0)
	roleID := make([]string, 0)

	for _, col := range rolePermission {
		permissionID = append(permissionID, col.PermissionID)
		roleID = append(roleID, col.RoleID)
	}

	err := r.config.DB.Debug().
		Unscoped().
		Where("permission_id IN ?", permissionID).
		Where("role_id IN ?", roleID).
		Delete(&rolePermission).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

package repository

import (
	"myadmin/app/entity"
	"myadmin/config"
)

type RoleRepository interface {
	Insert(role entity.Role) (entity.Role, error)
	FindAll() ([]entity.Role, error)
	FindById(ID string) (entity.Role, error)
	Update(role entity.Role, ID string) (entity.Role, error)
	Delete(ID string) (bool, error)
}

type roleRepositoryImpl struct {
	config config.Database
}

func NewRoleRepository(database config.Database) RoleRepository {
	return &roleRepositoryImpl{
		config: database,
	}
}

//	@Summary		: Insert role
//	@Description	: Insert role to dataCreatedAt time.Time

func (r *roleRepositoryImpl) Insert(role entity.Role) (entity.Role, error) {
	err := r.config.DB.Create(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

// @Summary		: Get roles
// @Description	: -
// @Author			: andikaps
func (r *roleRepositoryImpl) FindAll() ([]entity.Role, error) {
	var roles []entity.Role

	err := r.config.DB.
		Debug().
		Preload("Menus").
		Preload("Permissions").
		Find(&roles).
		Error

	if err != nil {
		return roles, err
	}

	return roles, nil
}

// @Summary		: Get role
// @Description	: Find role by ID
// @Author			: andikaps
func (r *roleRepositoryImpl) FindById(ID string) (entity.Role, error) {
	var user entity.Role

	err := r.config.DB.Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary		: Update role
// @Description	: Update role by ID
// @Author			: andikaps
func (r *roleRepositoryImpl) Update(role entity.Role, ID string) (entity.Role, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

// @Summary		: Delete role
// @Description	: Delete role temporary
// @Author			: andikaps
func (r *roleRepositoryImpl) Delete(ID string) (bool, error) {
	var role entity.Role

	err := r.config.DB.Where("id = ?", ID).Delete(&role).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

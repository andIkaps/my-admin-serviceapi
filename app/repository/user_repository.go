package repository

import (
	"myadmin/app/entity"
	"myadmin/config"
)

type UserRepository interface {
	Insert(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindByUsername(username string) (entity.User, error)
	FindById(ID string) (entity.User, error)
	Update(user entity.User, ID string) (entity.User, error)
	Delete(ID string) (bool, error)
	ForceDelete(ID string) (bool, error)
}

type userRepositoryImpl struct {
	config config.Database
}

func NewUserRepository(database config.Database) UserRepository {
	return &userRepositoryImpl{
		config: database,
	}
}

//	@Summary		: Insert user
//	@Description	: Insert user to dataCreatedAt time.Time

func (r *userRepositoryImpl) Insert(user entity.User) (entity.User, error) {
	err := r.config.DB.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary		: Get users
// @Description	: -
// @Author			: andikaps
func (r *userRepositoryImpl) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := r.config.DB.Preload("Roles").Order("created_at DESC").Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

// @Summary		: Get user
// @Description	: Find user by Username
// @Author			: andikaps
func (r *userRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	var user entity.User

	r.config.DB.Where("username = ?", username).First(&user)

	return user, nil
}

// @Summary		: Get user
// @Description	: Find user by ID
// @Author			: andikaps
func (r *userRepositoryImpl) FindById(ID string) (entity.User, error) {
	var user entity.User

	err := r.config.DB.Where("id = ?", ID).Preload("Roles.Menus").Preload("Roles.Permissions").First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary		: Update user
// @Description	: Update user
// @Author			: andikaps
func (r *userRepositoryImpl) Update(user entity.User, ID string) (entity.User, error) {
	err := r.config.DB.Where("id = ?", ID).Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// @Summary		: Delete user
// @Description	: Delete user temporary
// @Author			: andikaps
func (r *userRepositoryImpl) Delete(ID string) (bool, error) {
	var user entity.User

	err := r.config.DB.Where("id = ?", ID).Delete(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// @Summary		: Delete user
// @Description	: Delete user permanent
// @Author			: andikaps
func (r *userRepositoryImpl) ForceDelete(ID string) (bool, error) {
	var user entity.User

	err := r.config.DB.Unscoped().Where("id = ?", ID).Delete(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

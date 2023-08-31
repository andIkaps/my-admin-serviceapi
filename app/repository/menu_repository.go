package repository

import (
	"myadmin/app/entity"
	"myadmin/config"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MenuRepository interface {
	Insert(menu entity.Menu) (entity.Menu, error)
	FindAll() ([]entity.Menu, error)
	FindByID(ID string) (entity.Menu, error)
	Update(menu entity.Menu, ID string) (entity.Menu, error)
	Delete(ID string) (bool, error)
}

type menuRepositoryImpl struct {
	config config.Database
}

func NewMenuRepository(database config.Database) MenuRepository {
	return &menuRepositoryImpl{
		config: database,
	}
}

//	@Summary		: Insert menu
//	@Description	: Insert menu to dataCreatedAt time.Time

func (r *menuRepositoryImpl) Insert(menu entity.Menu) (entity.Menu, error) {
	err := r.config.DB.Debug().Create(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

//	@Summary		: Get menus
//	@Description	: -
//	@Author			: andikaps
func (r *menuRepositoryImpl) FindAll() ([]entity.Menu, error) {
	var menus []entity.Menu

	err := r.config.DB.Debug().
		Order("ord ASC").
		Preload("Childs.Childs."+clause.Associations, func(db *gorm.DB) *gorm.DB {
			return db.Order("ord ASC")
		}).
		Where("parent_id", nil).
		Find(&menus).Error

	if err != nil {
		return menus, err
	}

	return menus, nil
}

//	@Summary		: Get menu
//	@Description	: find menu by ID
//	@Author			: andikaps
func (r *menuRepositoryImpl) FindByID(ID string) (entity.Menu, error) {
	var menu entity.Menu

	err := r.config.DB.Debug().Where("id = ?", ID).First(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

//	@Summary		: Update menu
//	@Description	: Update menu by ID
//	@Author			: andikaps
func (r *menuRepositoryImpl) Update(menu entity.Menu, ID string) (entity.Menu, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

//	@Summary		: Delete menu
//	@Description	: Delete menu temporary
//	@Author			: andikaps
func (r *menuRepositoryImpl) Delete(ID string) (bool, error) {
	var menu entity.Menu

	err := r.config.DB.Debug().Where("id = ?", ID).Delete(&menu).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

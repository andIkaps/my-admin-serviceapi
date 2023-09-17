package repository

import (
	"myadmin/app/entity"
	"myadmin/config"
)

type MenuSectionRepository interface {
	Insert(menuSection entity.MenuSection) (entity.MenuSection, error)
	FindAll() ([]entity.MenuSection, error)
	FindByID(ID string) (entity.MenuSection, error)
	Update(menuSection entity.MenuSection, ID string) (entity.MenuSection, error)
	Delete(ID string) (bool, error)
}

type menuSectionRepositoryImpl struct {
	config config.Database
}

func NewMenuSectionRepository(database config.Database) MenuSectionRepository {
	return &menuSectionRepositoryImpl{
		config: database,
	}
}

// @Summary : Inset Menu Section
// @Description : Inset menu to database
func (r *menuSectionRepositoryImpl) Insert(menuSection entity.MenuSection) (entity.MenuSection, error) {
	err := r.config.DB.Debug().Create(&menuSection).Error

	if err != nil {
		return menuSection, err
	}

	return menuSection, nil
}

// @Summary : Get menu section
// @Description : get all data in menu section
func (r *menuSectionRepositoryImpl) FindAll() ([]entity.MenuSection, error) {
	var menuSections []entity.MenuSection

	err := r.config.DB.Debug().
		Order("ord ASC").
		Preload("MenuItem.Childs.Childs").
		Find(&menuSections).
		Error
	if err != nil {
		return menuSections, err
	}

	return menuSections, nil
}

// @Summary : Get menu section by id
// @Description : get data menu section by id
func (r *menuSectionRepositoryImpl) FindByID(ID string) (entity.MenuSection, error) {
	var menuSection entity.MenuSection

	err := r.config.DB.Debug().Where("id = ?", ID).First(&menuSection).Error
	if err != nil {
		return menuSection, err
	}

	return menuSection, nil
}

// @Summary		: Update menu section
// @Description	: Update menu section by ID
func (r *menuSectionRepositoryImpl) Update(menuSection entity.MenuSection, ID string) (entity.MenuSection, error) {
	err := r.config.DB.Debug().Where("id = ?", ID).Updates(&menuSection).Error

	if err != nil {
		return menuSection, err
	}

	return menuSection, nil
}

// @Summary		: Delete menu section
// @Description	: Delete menu section temporary
func (r *menuSectionRepositoryImpl) Delete(ID string) (bool, error) {
	var menuSection entity.MenuSection

	err := r.config.DB.Debug().Where("id = ?", ID).Delete(&menuSection).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

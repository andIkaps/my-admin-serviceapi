package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
	"myadmin/app/request"
)

type MenuSectionService interface {
	List() ([]entity.MenuSection, error)
	Insert(req request.MenuSectionRequest) (entity.MenuSection, error)
	FindById(ID string) (entity.MenuSection, error)
	Update(req request.MenuSectionRequest, ID string) (entity.MenuSection, error)
	Delete(ID string) (bool, error)
}

type menuSectionServiceImpl struct {
	repository repository.MenuSectionRepository
}

func NewMenuSectionService(r repository.MenuSectionRepository) MenuSectionService {
	return &menuSectionServiceImpl{
		repository: r,
	}
}

// @Summary		: List menu section
// @Description	: Get menus from repository
// @Author			: andikaps
func (s *menuSectionServiceImpl) List() ([]entity.MenuSection, error) {
	menuSections, err := s.repository.FindAll()

	if err != nil {
		return menuSections, err
	}

	return menuSections, nil
}

// @Summary		: Insert menu
// @Description	: Get menu to repository
// @Author			: andikaps
func (s *menuSectionServiceImpl) Insert(req request.MenuSectionRequest) (entity.MenuSection, error) {
	menuSection := entity.MenuSection{
		Name:      req.Name,
		Ord:       req.Ord,
		CreatedBy: req.CreatedBy,
	}

	newMenuSection, err := s.repository.Insert(menuSection)

	if err != nil {
		return newMenuSection, err
	}

	return newMenuSection, nil
}

// @Summary		: Find menu section
// @Description	: Find menu section by id from repository
func (s *menuSectionServiceImpl) FindById(ID string) (entity.MenuSection, error) {
	menuSection, err := s.repository.FindByID(ID)

	if err != nil {
		return menuSection, err
	}

	return menuSection, nil
}

// @Summary		: Update menu section
// @Description	: Update menu section by id from repository
// @Author			: andikaps
func (s *menuSectionServiceImpl) Update(req request.MenuSectionRequest, ID string) (entity.MenuSection, error) {
	menuSection := entity.MenuSection{
		Name:      req.Name,
		Ord:       req.Ord,
		UpdatedBy: req.UpdatedBy,
	}

	updatedMenuSection, err := s.repository.Update(menuSection, ID)

	if err != nil {
		return updatedMenuSection, err
	}

	return updatedMenuSection, nil
}

// @Summary		: Delete menu section
// @Description	: Delete menu section to repository
// @Author		: andikaps
func (s *menuSectionServiceImpl) Delete(ID string) (bool, error) {
	_, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return true, nil
}

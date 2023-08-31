package service

import (
	"myadmin/app/entity"
	"myadmin/app/repository"
	"myadmin/app/request"
)

type MenuService interface {
	List() ([]entity.Menu, error)
	Insert(req request.MenuRequest) (entity.Menu, error)
	FindById(ID string) (entity.Menu, error)
	Update(req request.MenuRequest, ID string) (entity.Menu, error)
	Delete(ID string) (bool, error)
}

type menuServiceImpl struct {
	repository repository.MenuRepository
}

func NewMenuService(r repository.MenuRepository) MenuService {
	return &menuServiceImpl{
		repository: r,
	}
}

//	@Summary		: List menu
//	@Description	: Get menus from repository
//	@Author			: andikaps
func (s *menuServiceImpl) List() ([]entity.Menu, error) {
	menus, err := s.repository.FindAll()

	if err != nil {
		return menus, err
	}

	return menus, nil
}

//	@Summary		: Insert role
//	@Description	: Get role to repository
//	@Author			: andikaps
func (s *menuServiceImpl) Insert(req request.MenuRequest) (entity.Menu, error) {

	menu := entity.Menu{
		Name:     req.Name,
		Url:      req.Url,
		Icon:     req.Icon,
		Ord:      req.Ord,
		ParentID: req.ParentID,
	}

	newMenu, err := s.repository.Insert(menu)

	if err != nil {
		return newMenu, err
	}

	return newMenu, nil
}

//	@Summary		: Find role
//	@Description	: Find role by id from repository
//	@Author			: andikaps
func (s *menuServiceImpl) FindById(ID string) (entity.Menu, error) {
	menu, err := s.repository.FindByID(ID)

	if err != nil {
		return menu, err
	}

	return menu, nil
}

//	@Summary		: Update role
//	@Description	: Update role by id from repository
//	@Author			: andikaps
func (s *menuServiceImpl) Update(req request.MenuRequest, ID string) (entity.Menu, error) {

	menu := entity.Menu{
		Name:     req.Name,
		Url:      req.Url,
		Icon:     req.Icon,
		Ord:      req.Ord,
		ParentID: req.ParentID,
	}

	updateMenu, err := s.repository.Update(menu, ID)

	if err != nil {
		return updateMenu, err
	}

	return updateMenu, nil
}

//	@Summary		: Delete role
//	@Description	: Delete role to repository
//	@Author			: andikaps
func (s *menuServiceImpl) Delete(ID string) (bool, error) {
	menu, err := s.repository.Delete(ID)

	if err != nil {
		return false, err
	}

	return menu, nil
}

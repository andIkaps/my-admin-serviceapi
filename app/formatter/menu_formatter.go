package formatter

import (
	"myadmin/app/entity"
)

type MenuSectionFormatter struct {
	Name     string
	Ord      int
	MenuItem []MenuFormatter
}

type MenuFormatter struct {
	Name   string
	Url    string
	Icon   string
	Ord    int
	Childs []MenuFormatter
}

func FormatMenus(Menus []entity.Menu) []MenuFormatter {
	menusFormatter := []MenuFormatter{}

	for _, menu := range Menus {
		menuFormatter := MenuFormatter{
			Name: menu.Name,
			Url:  menu.Url,
			Icon: menu.Icon,
			Ord:  menu.Ord,
		}

		if len(menu.Childs) < 1 {
			menuFormatter.Childs = []MenuFormatter{}
		}

		if len(menu.Childs) > 0 {
			childs := FormatMenus(menu.Childs)
			menuFormatter.Childs = append(menuFormatter.Childs, childs...)
		}

		menusFormatter = append(menusFormatter, menuFormatter)
	}

	return menusFormatter
}

func FormatMenuSections(MenuSections []entity.MenuSection) []MenuSectionFormatter {
	menuSectionsFormatter := []MenuSectionFormatter{}

	// Menu Section
	for _, menuSection := range MenuSections {
		menusFormatter := []MenuFormatter{}

		// Menu Item
		for _, menu := range menuSection.MenuItem {
			menuFormatter := MenuFormatter{
				Name: menu.Name,
				Url:  menu.Url,
				Icon: menu.Icon,
				Ord:  menu.Ord,
			}

			if len(menu.Childs) > 0 {
				// Menu item childs
				menuChilds := FormatMenus(menu.Childs)

				menuFormatter.Childs = append(menuFormatter.Childs, menuChilds...)
			}

			if len(menu.Childs) < 1 {
				menuFormatter.Childs = []MenuFormatter{}
			}

			menusFormatter = append(menusFormatter, menuFormatter)
		}

		menuSec := MenuSectionFormatter{
			Name:     menuSection.Name,
			Ord:      menuSection.Ord,
			MenuItem: menusFormatter,
		}

		menuSectionsFormatter = append(menuSectionsFormatter, menuSec)
	}

	return menuSectionsFormatter
}

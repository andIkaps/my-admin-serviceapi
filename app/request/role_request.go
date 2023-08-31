package request

import "myadmin/app/entity"

type RoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

type RoleMenu struct {
	RoleID    int    `json:"role_id"`
	MenuID    int    `json:"menu_id"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type RoleHasMenu struct {
	Menus []entity.RoleMenu `json:"menus"`
}

type RoleHasPermission struct {
	Permission []entity.RolePermission `json:"permissions"`
}

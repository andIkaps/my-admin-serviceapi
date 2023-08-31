package request

import "myadmin/app/entity"

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	// Password  string `json:"password" validate:"required,min=8"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type UserRole struct {
	UserID    string `json:"user_id"`
	RoleID    string `json:"role_id"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type UserHasRole struct {
	Roles []entity.UserRole `json:"roles"`
}

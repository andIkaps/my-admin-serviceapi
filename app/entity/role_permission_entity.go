package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RolePermission struct {
	ID           string `gorm:"primary_key; not null"`
	RoleID       string `gorm:"not null;" json:"role_id"`
	PermissionID string `gorm:"not null;" json:"permission_id"`
}

func (d *RolePermission) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()

	return
}

type TablerRolePermission interface {
	TableName() string
}

func (RolePermission) TableName() string {
	return "tr_role_permissions"
}

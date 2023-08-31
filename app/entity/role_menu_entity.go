package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleMenu struct {
	ID        string `gorm:"primary_key; not null"`
	RoleID    string `gorm:"not null; index" json:"role_id"`
	MenuID    string `gorm:"not null; index" json:"menu_id"`
	CreatedAt time.Time
	CreatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt time.Time
	UpdatedBy string         `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (d *RoleMenu) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()

	return
}

type TablerRoleMenu interface {
	TableName() string
}

func (RoleMenu) TableName() string {
	return "tr_role_menus"
}

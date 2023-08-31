package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          string       `gorm:"primary_key; not null"`
	Name        string       `gorm:"size:100; not null"`
	Description string       `gorm:"default: NULL"`
	Menus       []Menu       `gorm:"many2many:tr_role_menus;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permissions []Permission `gorm:"many2many:tr_role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time
	CreatedBy   string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt   time.Time
	UpdatedBy   string         `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt   gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (d *Role) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()

	return
}

type TablerRole interface {
	TableName() string
}

func (Role) TableName() string {
	return "ms_roles"
}

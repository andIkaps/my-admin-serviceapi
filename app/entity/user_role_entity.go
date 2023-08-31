package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	ID        string `gorm:"primary_key; not null"`
	UserID    string `gorm:"not null; index" json:"user_id"`
	RoleID    string `gorm:"not_null; index" json:"role_id"`
	CreatedAt time.Time
	CreatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt time.Time
	UpdatedBy string         `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (d *UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()

	return
}

type TablerUserRole interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "tr_user_roles"
}

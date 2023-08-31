package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID          string `gorm:"primary_key; not null"`
	Name        string `gorm:"size:100; not null"`
	Description string `gorm:"default: NULL"`
	CreatedAt   time.Time
	CreatedBy   string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt   time.Time
	UpdatedBy   string         `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt   gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (d *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()

	return
}

type TablerPermission interface {
	TableName() string
}

func (Permission) TableName() string {
	return "ms_permissions"
}

package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Privilege struct {
	ID          string `gorm:"primary_key; not null"`
	Name        string `gorm:"size:100; not null"`
	Description string `gorm:"default: NULL"`
	CreatedAt   time.Time
	CreatedBy   string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt   time.Time
	UpdatedBy   string         `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt   gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (d *Privilege) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()

	return
}

type TablerPrivilege interface {
	TableName() string
}

func (Privilege) TableName() string {
	return "ms_privileges"
}

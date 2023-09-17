package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuSection struct {
	ID        string `gorm:"primary_key; not null"`
	Name      string `gorm:"size:50; not null"`
	Ord       int    `gorm:"default: NULL;default_order:ord desc"`
	MenuItem  []Menu `gorm:"foreignKey:MenuSectionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	CreatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt time.Time
	UpdatedBy string         `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (d *MenuSection) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()

	return
}

type TablerMenuSection interface {
	TableName() string
}

func (MenuSection) TableName() string {
	return "ms_menu_section"
}

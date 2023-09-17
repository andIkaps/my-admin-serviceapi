package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name          string    `gorm:"size:100; not null"`
	Url           string    `gorm:"size:100; not null"`
	Icon          string    `gorm:"default: NULL"`
	Ord           int       `gorm:"default: NULL;default_order:ord desc"`
	ParentID      string    `gorm:"default: NULL"`
	MenuSectionID string    `gorm:"size:60; default: NULL"`
	Childs        []Menu    `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time
	CreatedBy     string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt     time.Time
	UpdatedBy     string         `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt     gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (d *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.New()

	return
}

type TablerMenu interface {
	TableName() string
}

func (Menu) TableName() string {
	return "ms_menus"
}

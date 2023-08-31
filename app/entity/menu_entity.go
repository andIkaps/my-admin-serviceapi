package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID        string `gorm:"primary_key; not null"`
	Name      string `gorm:"size:100; not null"`
	Url       string `gorm:"size:100; not null"`
	Icon      string `gorm:"default: NULL"`
	Ord       int    `gorm:"default: NULL"`
	ParentID  string `gorm:"default: NULL"`
	Childs    []Menu `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	CreatedBy string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt time.Time
	UpdatedBy string         `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (d *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()

	return
}

type TablerMenu interface {
	TableName() string
}

func (Menu) TableName() string {
	return "ms_menus"
}

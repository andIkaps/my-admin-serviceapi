package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string     `gorm:"primary_key; not null"`
	Username    string     `gorm:"size:100; not null"`
	Name        string     `gorm:"size:100; not null"`
	Email       string     `gorm:"size:100; not null"`
	Password    string     `gorm:"size:60; not null"`
	Roles       []Role     `gorm:"many2many:tr_user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PrivilegeID string     `gorm:"default: NULL"`
	Privilege   *Privilege `gorm:"not null"`
	CreatedAt   time.Time
	CreatedBy   string `gorm:"size:100; not null; default: SYSTEM"`
	UpdatedAt   time.Time
	UpdatedBy   string         `gorm:"size:100; not null; default: SYSTEM"`
	DeletedAt   gorm.DeletedAt `json:"-" swaggerignore:"true"`
}

func (d *User) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()

	return
}

type TablerUser interface {
	TableName() string
}

func (User) TableName() string {
	return "ms_users"
}

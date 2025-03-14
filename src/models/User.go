package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/torrez/pkg"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string    `gorm:"type:varchar(100); not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone     string    `gorm:"type:varchar(100)"`
	Rol       Role      `gorm:"type:varchar(100);default:Usuario;not null"`
	Slug      string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
}

type Role string

const (
	Administrador Role = "Administrador"
	Usuario       Role = "Usuario"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.Slug = pkg.GenerateSlug(u.Name)
	return
}

package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Email        string         `gorm:"column:email;unique;not null"`
	PasswordHash string         `gorm:"column:password_hash;not null"`
	IsSuperAdmin bool           `gorm:"not null; default:false"`
	IsAdmin      bool           `gorm:"not null; default:false"`
	//Roles        []Role         `gorm:"many2many:user_roles;"`
	IsBlocked bool `gorm:"not null; default:false"`
}

//type Role struct {
//	gorm.Model
//	Name        string `gorm:"uniqueIndex;not null"`
//	Description string
//	Permissions []Permission `gorm:"many2many:role_permissions;"`
//}

//type Permission struct {
//	gorm.Model
//	Name        string `gorm:"uniqueIndex;not null"`
//	Description string
//}
//
//type UserRole struct {
//	UserID uuid.UUID `gorm:"type:uuid;not null"`
//	RoleID uint      `gorm:"primaryKey"`
//}
//
//type RolePermission struct {
//	RoleID       uint `gorm:"primaryKey"`
//	PermissionID uint `gorm:"primaryKey"`
//}

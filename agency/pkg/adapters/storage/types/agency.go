package types

import "gorm.io/gorm"

type Agency struct {
	gorm.Model
	Name    string `gorm:"not null;"`
	OwnerID uint
}

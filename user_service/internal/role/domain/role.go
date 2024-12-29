package domain

import "time"

type RoleId uint

type Role struct {
	ID        RoleId
	Name      string
	Type      string
	CreatedAt time.Time
}

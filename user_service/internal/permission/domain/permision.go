package domain

import "time"

type PermisionID uint

type Permision struct {
	ID        PermisionID
	Name      string
	Type      string
	Scope     string
	CreatedAt time.Time
}

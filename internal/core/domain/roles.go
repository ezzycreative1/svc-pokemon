package domain

import (
	"time"
)

type Roles struct {
	ID        int64     `json:"id"`
	Name      string    `json:"title" validate:"required"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewRoles(id int64, name string, status int8) *Roles {
	return &Roles{
		ID:        id,
		Name:      name,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

package domain

import (
	"time"
)

type Users struct {
	ID        int64     `json:"id"`
	Name      string    `json:"title" validate:"required"`
	RoleID    int64     `json:"role_id"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUsers(id int64, name string, roleID int64, status int8) *Users {
	return &Users{
		ID:        id,
		Name:      name,
		RoleID:    roleID,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

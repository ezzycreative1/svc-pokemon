package domain

import (
	"time"
)

type RoleRequest struct {
	Name string `json:"name" validate:"required"`
}

type Roles struct {
	ID        int64     `json:"id" gorm:"primaryKey,autoIncrement"`
	Name      string    `json:"name"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Roles) TableName() string {
	return "Roles"
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

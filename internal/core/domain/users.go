package domain

import (
	"time"
)

type StoreUserRequest struct {
	Fullname      string `json:"fullname" validate:"required"`
	Role          int64  `json:"role" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Password      string `json:"password" validate:"required"`
	CheckPassword string `json:"check_password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Users struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleID    int64     `json:"role_id"`
	Name      string    `json:"title"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
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

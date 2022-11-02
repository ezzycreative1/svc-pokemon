package domain

import (
	"time"
)

type StoreUserRequest struct {
	Fullname        string `json:"full_name" validate:"required"`
	Role            int64  `json:"role" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

type UpdateUserRequest struct {
	Fullname string `json:"full_name" validate:"required"`
	Role     int64  `json:"role" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	IsActive int8   `json:"is_Active"`
}

type SingleUserResponse struct {
	ID        int64     `json:"id"`
	Role      int64     `json:"role"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	IsActive  int8      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleID    int64     `json:"role_id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  int8      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Users) TableName() string {
	return "Users"
}

func NewUsers(id int64, roleID int64, fullname string, email string, password string, isActive int8) *Users {
	return &Users{
		ID:        id,
		RoleID:    roleID,
		FullName:  fullname,
		Email:     email,
		Password:  password,
		IsActive:  isActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

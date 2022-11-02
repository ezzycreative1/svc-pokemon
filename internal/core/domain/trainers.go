package domain

import "time"

type Trainers struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name"`
	Gender      string    `json:"email"`
	CountryCode string    `json:"password"`
	IsActive    int8      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Trainers) TableName() string {
	return "Trainers"
}

func NewTrainers(id int64, name string, gender string, countryCode string, isActive int8) *Trainers {
	return &Trainers{
		ID:          id,
		Name:        name,
		Gender:      gender,
		CountryCode: countryCode,
		IsActive:    isActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

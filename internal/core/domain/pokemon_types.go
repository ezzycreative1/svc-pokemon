package domain

import (
	"time"
)

type PokemonTypes struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (PokemonTypes) TableName() string {
	return "pokemon_type"
}

func NewPokemonTypes(id int64, name string) *PokemonTypes {
	return &PokemonTypes{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

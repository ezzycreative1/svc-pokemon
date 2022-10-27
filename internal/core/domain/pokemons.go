package domain

import (
	"time"
)

type Pokemons struct {
	ID        int64     `json:"id"`
	Name      string    `json:"title" validate:"required"`
	UserID    int64     `json:"user_id"`
	Stock     int64     `json:"stock" validate:"required"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPokemons(id int64, name string, userID int64, stock int64, status int8) *Pokemons {
	return &Pokemons{
		ID:        id,
		Name:      name,
		UserID:    userID,
		Stock:     stock,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type PokemonExternal struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PokemonMetaExternal struct {
}

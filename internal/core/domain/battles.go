package domain

import (
	"time"
)

type Battles struct {
	ID        int64     `json:"id"`
	PokemonID int64     `json:"pokemon_id"`
	UserID    int64     `json:"user_id"`
	Score     int64     `json:"score" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBattles(id int64, pokemonID int64, userID int64, score int64) *Battles {
	return &Battles{
		ID:        id,
		PokemonID: pokemonID,
		UserID:    userID,
		Score:     score,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

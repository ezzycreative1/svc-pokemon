package domain

import (
	"time"
)

type Teams struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TrainerID int64     `json:"trainer_id"`
	PokemonID int64     `json:"pokemon_id"`
	BattleID  int64     `json:"battle_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Teams) TableName() string {
	return "Teams"
}

func NewTeams(id int64, trainerID int64, pokemonID int64, battleID int64) *Teams {
	return &Teams{
		ID:        id,
		TrainerID: trainerID,
		PokemonID: pokemonID,
		BattleID:  battleID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

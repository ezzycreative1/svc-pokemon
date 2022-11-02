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

type Type struct {
	DamageRelations struct {
		DoubleDamageFrom []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"double_damage_from"`
		DoubleDamageTo []interface{} `json:"double_damage_to"`
		HalfDamageFrom []interface{} `json:"half_damage_from"`
		HalfDamageTo   []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"half_damage_to"`
		NoDamageFrom []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"no_damage_from"`
		NoDamageTo []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"no_damage_to"`
	} `json:"damage_relations"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID              int `json:"id"`
	MoveDamageClass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move_damage_class"`
	Moves []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"moves"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Pokemon []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		Slot int `json:"slot"`
	} `json:"pokemon"`
}

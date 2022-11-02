package domain

import (
	"time"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Battles struct {
	ID        int64     `json:"id"`
	Team1     int64     `json:"team_1"`
	Team2     int64     `json:"team_2"`
	Location  Location  `json:"location"`
	WinnerID  int64     `json:"winner_id"`
	Score     int64     `json:"score" validate:"required"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func NewBattles(id int64, team1 int64, team2 int64, location Location, winnerID int64, score int64) *Battles {
	return &Battles{
		ID:        id,
		Team1:     team1,
		Team2:     team2,
		Location:  location,
		WinnerID:  winnerID,
		Score:     score,
		StartTime: time.Now(),
		EndTime:   time.Now(),
	}
}

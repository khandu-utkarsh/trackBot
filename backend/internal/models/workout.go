package models

import (
	"time"
)

type Workout struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Duration    int       `json:"duration"` // in minutes
	Exercises   []Exercise `json:"exercises"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Exercise struct {
	ID          int64     `json:"id"`
	WorkoutID   int64     `json:"workout_id"`
	Name        string    `json:"name"`
	Sets        int       `json:"sets"`
	Reps        int       `json:"reps"`
	Weight      float64   `json:"weight"`
	Notes       string    `json:"notes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
} 
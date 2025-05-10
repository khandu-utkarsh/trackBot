package main

import (
	"workout_app_backend/services/workoutAppServices/internal/utils"
)

func LoadConfig() {
	utils.LoadEnv()
}

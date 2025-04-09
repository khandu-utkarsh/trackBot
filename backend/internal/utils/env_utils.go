package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("Unable to get current file path")
	}

	fmt.Println("Current file path:", file)
	fmt.Println("Current directory:", filepath.Dir(file))

	envPath := filepath.Join(filepath.Dir(file), "..", "..", ".env")

	fmt.Println("Reading env from: ", envPath)
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("No .env file found at path: %s\n", envPath)
	} else {
		log.Printf(".env file loaded from: %s\n", envPath)
	}
}

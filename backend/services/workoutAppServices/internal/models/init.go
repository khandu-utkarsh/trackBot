package models

import (
	"log"
	"os"
)

var modelsLogger *log.Logger

func init() {
	modelsLogger = log.New(os.Stdout, "Models: ", log.LstdFlags)
}

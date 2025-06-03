package routes

import (
	"log"
	"os"
)

var routesLogger *log.Logger

func init() {
	routesLogger = log.New(os.Stdout, "Routes: ", log.LstdFlags)
}

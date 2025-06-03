package middleware

import (
	"log"
	"os"
)

var middlewareLogger *log.Logger

func init() {
	middlewareLogger = log.New(os.Stdout, "Middleware: ", log.LstdFlags)
}

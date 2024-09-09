package initialize

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenvvariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}
}

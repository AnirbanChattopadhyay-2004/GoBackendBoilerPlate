package main

import (
	"example/server/initializer"
	"example/server/models"
	"log"
)

func init() {
	initialize.Loadenvvariable()
	initialize.Connect()
}
func main() {
	log.Println("Running migrations...",initialize.DB)
    err := initialize.DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    } else {
        log.Println("Migration successful!")
    }
}
package initialize

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB  *gorm.DB
func Connect() {
	// Connect to the database
	
	  var Err error
	  dsn := os.Getenv("Db_url")
	  fmt.Println(dsn)
	  DB, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	  fmt.Println(DB)
	  if Err != nil {
		log.Fatal("error is ",Err)
	  } 
	  
}
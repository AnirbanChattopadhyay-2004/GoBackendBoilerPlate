package controls

import (
	initialize "example/server/initializer"
	"example/server/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddValDb(c *gin.Context) {
	// Add your logic here
	fmt.Println("Post")
	var body struct{
		Name string 
	}
	c.BindJSON(&body)
	// Add your logic here
	posts := models.User{Name:body.Name}
	results := initialize.DB.Create(&posts)
	if results.Error != nil {
		fmt.Println(results.Error)
	}
	c.IndentedJSON(http.StatusAccepted,posts)

}
func GetValDb(c *gin.Context) {
	// Add your logic here
	Name := c.Param("name")
	// Add your logic here
	fmt.Println(Name)
	// creating the variable here and with the help of pointer the value will be inserted here.
	var user models.User
	result := initialize.DB.Where("Name = ?",Name).First(&user)
	if result.Error != nil {
		log.Fatal("error is ",result.Error)
	}
	c.IndentedJSON(http.StatusAccepted,user)
}
func GetAllValDb(c  *gin.Context) {
	var users []models.User
	result := initialize.DB.Find(&users)
	if result.Error != nil{
		log.Fatal("Error ",result.Error)
	}
	c.IndentedJSON(http.StatusAccepted,users)
}
func GetUpdate(c  *gin.Context) {
	// Get updated records count with `RowsAffected`
	result := initialize.DB.Where("Id = ?", "1").Updates(models.User{ Name: "Anirban"})

	if result.Error != nil{
		log.Fatal("Error ",result.Error)
	}
	c.IndentedJSON(http.StatusAccepted,"Updated")
}
// do the delete one after same as above
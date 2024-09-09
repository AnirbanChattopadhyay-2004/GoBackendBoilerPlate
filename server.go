package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example/server/initializer"
	"example/server/controllers"
)

type value struct{
	ID string `json:"id"`
	Name string `json:"name"`
}
var values = []value{
	{ID: "1", Name: "anirban"},
	{ID:"2",Name:"surjo"},
}
func getVal(c  *gin.Context) {
	c.IndentedJSON(http.StatusOK,values)
}
func addVal(c  *gin.Context) {
	var newvalue value
	err := c.BindJSON(&newvalue)
	if(err != nil) {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	// Add the new value to the slice
	values = append(values,newvalue)
	c.IndentedJSON(http.StatusAccepted,newvalue)
}
func getValId(c  *gin.Context) {
	id := c.Param("id")
	// Find the value with the given ID
	var arr []string
	for _,v := range  values {
		if(v.ID  == id) {
			arr	= append(arr, v.Name)
		}

	}
	c.IndentedJSON(http.StatusAccepted,arr)

}
func init() {
	initialize.Loadenvvariable()
	initialize.Connect()
}
func main(){
	router := gin.Default()
	router.GET("/values", getVal)
	router.POST("/add",addVal)
	router.GET("/values/:id",getValId)
	router.POST("/db",controls.AddValDb)
	router.GET("/db/:name",controls.GetValDb)
	router.GET("/db/getallvals",controls.GetAllValDb)
	router.GET("/db/update",controls.GetUpdate)
	router.Run("localhost:8080")
}
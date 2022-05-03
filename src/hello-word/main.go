package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	FirstName string `xml:"firstName, attr"`
	LastName  string `xml:"lastName, attr"`
}

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func ParamsHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name,
	})
}

func XMLHandler(c *gin.Context) {
	c.XML(http.StatusOK, Person{
		FirstName: "Andrei",
		LastName:  "Santos",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.GET("/params/:name", ParamsHandler)
	router.GET("/xml", XMLHandler)
	router.Run()
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func HomePage (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func PostHomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Post Hello World",
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name") // name=anthony
	age := c.Query("age")	//age=25
	c.JSON(http.StatusOK, gin.H{
		"name" : name,
		"age" : age,
	})
}

func PathParams(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"name" : name,
		"age" : age,
	})
}

func PostRetrieve(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(value),
	})
}

func main() {
	fmt.Println("Yo!")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.POST("/retrieve", PostRetrieve)
	r.GET("/query", QueryStrings)	// /query?name=anthony?age=25
	r.GET("/path/:name/:age", PathParams)	// /path/anthony/25/
	// PUT, DEL, PATCH
	r.Run()
}
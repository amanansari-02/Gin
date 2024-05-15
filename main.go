package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getData)
	router.POST("/", getPostData)
	router.GET("/nameAndAge", getNameAndAge)
	router.GET("/getData/:name", getDataById)

	router.Run(":8080")
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Gin successfully connected",
	})
}

func getPostData(c *gin.Context) {
	body := c.Request.Body
	val, _ := ioutil.ReadAll(body)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Display body data",
		"data":    val,
	})
}

func getNameAndAge(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"message": "Display user data",
		"name":    name,
		"age":     age,
	})
}

func getDataById(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Data by id",
		"data": gin.H{
			"name": name,
		},
	})
}

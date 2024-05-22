package controller

import (
	"Gin/src/api/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetData(c *gin.Context) {
	dataModel := model.GetData{
		Name:     "Aman",
		Email:    "aman@gmail.com",
		Password: "12345",
		Age:      20,
		MobileNo: 9156033705,
	}
	data, _ := json.Marshal(dataModel)     //Convert into JSON format
	logrus.Println("data: ", string(data)) // Log data
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   string(data),
	})
}

func GetNameAndAge(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	logrus.Println("name = ", name, ", age = ", age)
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

package controller

import (
	"Gin/src/api/model"
	"Gin/src/config"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c *gin.Context) {
	var data model.GetData //Modal Name
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} // Check any error comes to call model

	client := config.GetMongoClient()                         // Take a database
	collection := client.Database("ginDB").Collection("user") // Database name and collection name

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Check api less than 10 sec time

	var existingUser model.GetData
	newErr := collection.FindOne(ctx, bson.M{"email": data.Email}).Decode(&existingUser) // Check email already exists
	if newErr == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Email already exists"})
		return
	} // if email already exists then display

	_, err := collection.InsertOne(ctx, data) // Insert data
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} // any error comes to create a user

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusCreated,
		"message": "User craeted successfully",
		"data":    data,
	}) // success message
}

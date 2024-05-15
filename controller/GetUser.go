package controller

import (
	"context"
	"net/http"
	"strconv"
	model "user-api/Model"
	"user-api/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	filter := bson.D{{Key: "id", Value: idInt}}

	var result model.User

	err := database.Collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, result)
}

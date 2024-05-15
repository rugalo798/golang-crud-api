package controller

import (
	"context"
	"net/http"
	"strconv"
	"user-api/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	filter := bson.D{{Key: "id", Value: idInt}}

	result, err := database.Collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)
}

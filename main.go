package main

import (
	"user-api/controller"
	"user-api/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// Conexão com banco de dados e iniciando o servidor
	database.InitDB()
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// Criando as rotas e passando os controllers
	router.POST("/users", controller.CreateUser)
	router.GET("/users", controller.GetUsers)
	router.DELETE("/users/:id", controller.DeleteUser)
	router.GET("/users/:id", controller.GetUser)

	// Rodando o servidor na porta 8080
	router.Run(":8080")

}

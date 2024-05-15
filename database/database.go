package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Pegando o client e a collection
var (
	Client     *mongo.Client
	Collection *mongo.Collection
)

func InitDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Algum erro aconteceu: %s", err.Error())
	}

	URL := os.Getenv("URL")

	clientOptions := options.Client().ApplyURI(URL)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	Collection = Client.Database("api").Collection("users")
}

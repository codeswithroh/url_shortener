package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"url-shortener/internal/handler"
	"url-shortener/internal/services"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
	dotenvError := godotenv.Load()

	if dotenvError != nil {
		fmt.Println("Failed to load .env file")
		return
	}

	client, errors := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if errors != nil {
		fmt.Println("Failed to connect to MongoDB:", errors)
		return
	}

	services.InitMongoDB(client)

	fmt.Println("Server starting on port 8080")

	http.HandleFunc("/", handler.Handler)
	http.HandleFunc("/shorten", handler.ShortURLHandler)
	http.HandleFunc("/redirect/", handler.RedirectToLongURLHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Println("Server failed to start for this: ",err)
	}
}
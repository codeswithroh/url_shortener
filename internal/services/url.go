package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
	"url-shortener/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// var urlDB = make(map[string]model.URL)
var collection *mongo.Collection

func InitMongoDB(client *mongo.Client) {
	collection = client.Database("url-shortener").Collection("urls")
}

func GenerateShortURL(longURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(longURL))
	data := hasher.Sum(nil)
     
	hash := hex.EncodeToString(data)
	
	return hash[:8]
}

func CreateShortURL(longURL string) string {
	shortURL := GenerateShortURL(longURL)
	id := shortURL

	url := model.URL {
		ID: id,
		LongURL: longURL,
		ShortURL: shortURL,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	_, err := collection.InsertOne(context.Background(), url)

	if err != nil {
		panic(err)
	}

	return shortURL
}

func GetLongURL(shortURL string) (model.URL,error) {
	var url model.URL

	fmt.Println("shortURL", shortURL)

	err := collection.FindOne(context.Background(), bson.M{"shortUrl": shortURL}).Decode(&url)

	if err != nil {
		return model.URL{}, err
	}

	if url.ExpiresAt.Before(time.Now()) {
		return model.URL{}, fmt.Errorf("url is expired")
	}

	return url, nil
}
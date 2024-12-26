package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"url-shortener/internal/services"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"This is a url shortener service")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	shortURL := services.CreateShortURL(data.URL)

	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: os.Getenv("REDIRECT_URL") + shortURL,
	}

	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(response)
}

func RedirectToLongURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	
	url, err := services.GetLongURL(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.LongURL, http.StatusMovedPermanently)
}
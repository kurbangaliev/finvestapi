package handlers

import (
	"encoding/json"
	"finvestapi/internal/db"
	"finvestapi/internal/models"
	"net/http"
)

// HandleLikeNews PUT /news/like/ /news/dislike/
func HandleLikeNews(w http.ResponseWriter, r *http.Request) {
	var item models.NewsLike
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err := db.LikeNews(item)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News LikeNews updated"))
}

// HandleViewNews PUT /news/view/
func HandleViewNews(w http.ResponseWriter, r *http.Request) {
	var item models.NewsViewing
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := db.ViewNews(item)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News updated"))
}

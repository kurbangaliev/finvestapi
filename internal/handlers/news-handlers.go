package handlers

import (
	"encoding/json"
	"finvestapi/internal/db"
	"finvestapi/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HandleAllNews GET /news/
func HandleAllNews(w http.ResponseWriter, r *http.Request) {
	items, err := db.LoadAllNews()
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"error": "` + err.Error() + `"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// HandleNewsGroup GET /news/group/{id}
func HandleNewsGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	items, err := db.LoadNewsGroup(id)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"error": "` + err.Error() + `"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// HandleEnableNews GET /enableNews/
func HandleEnableNews(w http.ResponseWriter, r *http.Request) {
	items, err := db.LoadEnableNews()
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"error": "` + err.Error() + `"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

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

// HandleGetLikeNews GET /news/{id}
func HandleGetLikeNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	count, err := db.GetLikesView(id)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"error": "` + err.Error() + `"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}

// HandleViewNews PUT /news/view/
func HandleViewNews(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleViewNews")
	var item models.NewsViewing
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := db.ViewNews(item)
	if err != nil {
		http.Error(w, "Failed to create NewsViewing item", http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotModified)
		w.Write([]byte("News Viewing not updated"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News Viewing updated"))
	log.Printf("Backend Server. News Viewing updated. Items: NewsId=%d UserId=%d", item.NewsId, item.UserId)
}

// HandleGetLikeNews GET /news/analytics/{id}
func HandleGetNewsAnalytics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	item, err := db.GetNewsAnalytics(id)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"error": "` + err.Error() + `"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// HandleGetNewsUnread GET /news/loadUnread/{id}
func HandleGetNewsUnread(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	item, err := db.GetUnreadNews(id)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"error": "` + err.Error() + `"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// HandleAddNews POST /news/
func HandleAddNews(w http.ResponseWriter, r *http.Request) {
	var item models.News
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	item, err := db.SaveNews(item)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News added"))
}

// HandleEditNews PUT /news/{id}
func HandleEditNews(w http.ResponseWriter, r *http.Request) {
	var item models.News
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	item, err := db.UpdateNews(item)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News updated"))
}

package handlers

import (
	"encoding/json"
	"finvestapi/internal/db"
	"log"
	"net/http"
)

/* =================== Objects =================== */

// HandleGetObjects GET /news
func HandleGetObjects[T comparable](w http.ResponseWriter, r *http.Request) {
	items, err := db.SelectAll[T]()
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"error": "` + err.Error() + `"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// HandleAddObject POST /news
func HandleAddObject[T comparable](w http.ResponseWriter, r *http.Request) {
	log.Println("HandleAddObject")

	var item T
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	log.Printf("Add Item: %v", item)
	item, err := db.SaveObject[T](item)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// HandleEditObject PUT /news/{id}
func HandleEditObject[T comparable](w http.ResponseWriter, r *http.Request) {
	log.Println("HandleEditObject")

	var item T
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err := db.UpdateObject(item)
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News updated"))
}

// HandleDeleteObject DELETE /news/{id}
func HandleDeleteObject[T comparable](w http.ResponseWriter, r *http.Request) {
	log.Println("HandleDeleteObject")

	var item T
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := db.DeleteObject(item)
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News deleted"))
}

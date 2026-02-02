package handlers

import (
	"encoding/json"
	"finvestapi/internal/db"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

// HandleGetObject GET /news/{id}
func HandleGetObject[T comparable](w http.ResponseWriter, r *http.Request) {
	var item T

	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	strBody := `{"id": ` + strconv.Itoa(id) + `}`
	reader := strings.NewReader(strBody)
	if err := json.NewDecoder(reader).Decode(&item); err != nil {
		//bodyStr, _ := io.ReadAll(r.Body)
		log.Printf("Get Item: %s", strBody)
		http.Error(w, "Invalid JSON. "+string(strBody), http.StatusBadRequest)
		return
	}

	log.Printf("Get Item: %v\n", item)
	item, err = db.Select[T](item)
	if err != nil {
		http.Error(w, "Failed to get item", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
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
	log.Printf("Edit Item: %v\n", item)
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
		bodyStr, _ := io.ReadAll(r.Body)
		log.Println("Delete Item:", string(bodyStr))
		http.Error(w, "Invalid JSON. "+string(bodyStr), http.StatusBadRequest)
		return
	}

	err := db.DeleteObject(item)
	if err != nil {
		log.Printf("Failed to delete item: %s\n", err.Error())
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Item deleted"}`))
}

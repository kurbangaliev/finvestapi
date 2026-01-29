package main

import (
	"finvestapi/internal/db"
	"finvestapi/internal/handlers"
	"finvestapi/internal/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Println("Backend server is starting...")
	log.Println("Database auto migrate...")
	err := db.AutoMigrate()
	if err != nil {
		log.Println(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "https://cabinet.finvest.kz", "http://192.168.111.11:8080", "http://192.168.111.103"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		Debug:            false,
	})

	mainRouter := mux.NewRouter()

	// --- Version 1 API ---
	apiV1 := mainRouter.PathPrefix("/api/v1").Subrouter()
	//news handlers
	apiV1.HandleFunc("/news/", handlers.HandleAddObject[models.News]).Methods("POST")
	apiV1.HandleFunc("/news/{id}", handlers.HandleEditObject[models.News]).Methods("PUT")
	apiV1.HandleFunc("/news/", handlers.HandleAllNews).Methods("GET")
	apiV1.HandleFunc("/news/{id}", handlers.HandleDeleteObject[models.News]).Methods("DELETE")
	//LikeNews and Dislike
	apiV1.HandleFunc("/news/like/", handlers.HandleLikeNews).Methods("PUT")
	apiV1.HandleFunc("/news/dislike/", handlers.HandleLikeNews).Methods("PUT")
	apiV1.HandleFunc("/news/view/", handlers.HandleViewNews).Methods("PUT")
	////manager handlers
	//r.HandleFunc("/managers/", handlers.HandleAddManager).Methods("POST")
	//r.HandleFunc("/managers/{id}", handlers.HandleEditManagers).Methods("PUT")
	//r.HandleFunc("/managers/", handlers.HandleGetManagers).Methods("GET")
	//r.HandleFunc("/managers/{id}", handlers.HandleDeleteManager).Methods("DELETE")
	//security handlers
	mainRouter.HandleFunc("/login", handlers.HandleLogin).Methods("POST")

	log.Println("🚀 Backend server started on http://localhost:8081")

	handler := c.Handler(mainRouter)
	log.Fatal(http.ListenAndServe(":8081", handler))
}

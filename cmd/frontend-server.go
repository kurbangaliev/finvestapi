package main

import (
	"finvestapi/internal/db"
	"finvestapi/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	staticDir = "/assets/"
)

func main() {
	fmt.Println("Server is starting...")
	fmt.Println("Database auto migrate...")
	err := db.AutoMigrate()
	if err != nil {
		log.Println(err)
	}

	err = db.CreateDefaultUser()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Frontend server handling requests...")
	r := mux.NewRouter()

	staticPath := http.Dir("./web/assets")
	fs := http.FileServer(staticPath)

	prometheus.MustRegister(handlers.HttpCounter)

	// Public Routes
	r.HandleFunc("/news", handlers.ShowNews).Methods("GET")

	//Prometey handler
	r.Handle("/metrics", promhttp.Handler())
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, fs)).Methods("GET")

	// Authentification
	r.HandleFunc("/", handlers.ShowLoginPage).Methods("GET")
	r.HandleFunc("/logout", handlers.HandlerLogout).Methods("POST")

	// Private Routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(handlers.JWTAuth)
	// Admin page handlers
	protected.HandleFunc("/newsAdd", handlers.NewsAddPage).Methods("GET")
	protected.HandleFunc("/newsBrowser", handlers.NewsBrowserPage).Methods("GET")

	//log.Println("🚀 Frontend Server started on http://localhost:8080")
	//log.Fatal(http.ListenAndServe(":8080", r))

	log.Println("🚀 Frontend Server started on https://localhost:7443")
	log.Fatal(http.ListenAndServeTLS(":7443",
		"certs/finvestapi.crt",
		"certs/finvestapi.key",
		r))
}

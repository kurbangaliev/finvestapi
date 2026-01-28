package handlers

import (
	"finvestapi/internal/db"
	"finvestapi/internal/models"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var HttpCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_request_total",
	Help: "Total numbers of HTTP Requests"},
	[]string{"path"},
)

func ShowNews(writer http.ResponseWriter, request *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/news.html")
	if err != nil {
		fmt.Printf("Error parsing news.html: %v \n", err)
	}

	news, err := db.SelectAll[models.News]()
	if err != nil {
		http.Error(writer, "Error load news", http.StatusInternalServerError)
		log.Fatal(err)
	}

	err = tmpl.Execute(writer, news)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func ShowSustainableDevelopment(writer http.ResponseWriter, request *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/sustainableDevelopment.html")
	if err != nil {
		fmt.Printf("Error parsing sustainableDevelopment.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func ShowServicesPage(writer http.ResponseWriter, request *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/services.html")
	if err != nil {
		fmt.Printf("Error parsing services.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func NewsAddPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/newsAdd.html")
	if err != nil {
		fmt.Printf("Error parsing newsAdd.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func NewsBrowserPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/newsBrowser.html")
	if err != nil {
		fmt.Printf("Error parsing newsBrowser.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func ShowLoginPage(writer http.ResponseWriter, r *http.Request) {
	HttpCounter.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	tmpl, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		fmt.Printf("Error parsing login.html: %v \n", err)
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

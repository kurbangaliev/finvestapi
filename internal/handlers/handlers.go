package handlers

import (
	"finvestapi/internal/db"
	"finvestapi/internal/models"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

type GenericFunc[T comparable] func() ([]T, error)
type DataFuncParams func(map[string]string) (any, error)

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

func HandlerTemplate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ShowTemplatePage(templatePage string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("ShowTemplatePage [%s] \n]...", templatePage)
		HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
		tmpl, err := template.ParseFiles(templatePage)
		if err != nil {
			fmt.Printf("Error parsing services.html: %v \n", err)
		}

		err = tmpl.Execute(writer, nil)
		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}
		handler.ServeHTTP(writer, request)
	})
}

func ShowTemplatePageGeneric[T comparable](templatePage string, handler http.Handler, dataFunc GenericFunc[T]) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("ShowTemplatePageGeneric [%s]...", templatePage)
		HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
		tmpl, err := template.ParseFiles(templatePage)
		if err != nil {
			fmt.Printf("Error parsing services.html: %v \n", err)
		}

		data, err := dataFunc()
		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}

		err = tmpl.Execute(writer, data)
		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}
		handler.ServeHTTP(writer, request)
	})
}

func ShowTemplatePageParams(templatePage string, handler http.Handler, dataFunc DataFuncParams) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("ShowTemplatePageParams [%s]...", templatePage)
		HttpCounter.With(prometheus.Labels{"path": request.URL.Path}).Inc()
		vars := mux.Vars(request)
		tmpl, err := template.ParseFiles(templatePage)
		if err != nil {
			fmt.Printf("Error parsing services.html: %v \n", err)
		}

		data, err := dataFunc(vars)

		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}

		err = tmpl.Execute(writer, data)
		if err != nil {
			http.Error(writer, "Error rendering template page "+templatePage, http.StatusInternalServerError)
			log.Printf("Error executing template: [%s] %v", templatePage, err)
		}
		handler.ServeHTTP(writer, request)
	})
}

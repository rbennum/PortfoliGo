package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rbennum/PortfoliGo/data"
)

func main() {
	http.HandleFunc("/", homeHandler)
	for _, project := range data.Projects {
		http.HandleFunc(project.URL, projectHandler(project))
	}

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.ParseFiles(filepath.Join("templates", "index.html")),
	)
	tmpl.Execute(w, data.Projects)
}

func projectHandler(project data.Project) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(
			template.ParseFiles(filepath.Join("templates", "portfolio.html")),
		)
		tmpl.Execute(w, project)
	}
}
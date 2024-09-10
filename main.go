package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rbennum/PortfoliGo/data"
)

func main() {
	appPort := os.Getenv("APP_PORT")
	http.HandleFunc("/", homeHandler)
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))

	log.Printf("Starting server on :%s", appPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", appPort), nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.ParseFiles(
			filepath.Join("templates", "index.html"),
			filepath.Join("templates", "header.html"),
			filepath.Join("templates", "footer.html"),
			filepath.Join("templates", "main.html"),
			filepath.Join("templates", "tech_list.html"),
			filepath.Join("templates", "past_projects.html"),
		),
	)
	tmpl.Execute(w, data.PageDataItem)
}

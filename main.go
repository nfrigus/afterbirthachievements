package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

// We store all templates on first launch for efficiency.
var templates map[string]*template.Template

// initializeTemplates populates `templates` for use in our handlers. The logic is that each
// of our templates is composed by base.html and some other HTML template.
func initializeTemplates() (err error) {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templateFiles, err := filepath.Glob("tmpl/*.html")
	if err != nil {
		return
	}
	for _, t := range templateFiles {
		if t != "tmpl/base.html" {
			templates[t] = template.Must(template.ParseFiles("tmpl/base.html", t))
		}
	}
	return
}

// renderContent parses the content (given as a template) and puts it into our base template.
// The control of the input data is handled by the handlers in handlers.go
func renderContent(t string, r *http.Request, w http.ResponseWriter, data interface{}) {
	// Set up various security headers
	w.Header().Set("Content-Security-Policy", "default-src 'self'; img-src 'self' https://steamcdn-a.akamaihd.net")
	w.Header().Set("Referrer-Policy", "no-referrer")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Xss-Protection", "1; mode=block")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// Combine templates and write response
	err := templates[t].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err)
	}
}

// initialize sets up all global variables relevant for the application
func initialize() {
	initializeTemplates()
	readConfig()
	readAllCategories()
	readAllAchievements()
}

func main() {
	initialize()
	staticHandler := http.FileServer(http.Dir("tmpl"))
	http.Handle("/css/", staticHandler)
	http.Handle("/img/", staticHandler)
	http.Handle("/js/", staticHandler)

	router := mux.NewRouter()
	router.HandleFunc("/", landingHandler)
	router.HandleFunc("/{username:[a-z0-9_-]+}", achievementHandler)
	http.Handle("/", router)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.WebserverPort), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

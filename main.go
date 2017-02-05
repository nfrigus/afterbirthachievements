package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// initialize sets up all global variables relevant for the application
func initialize() {
	readConfig()
	readAllCategories()
	readAllAchievements()
}

func main() {
	initialize()
	staticHandler := http.FileServer(http.Dir("tmpl"))
	http.Handle("/css/", staticHandler)
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

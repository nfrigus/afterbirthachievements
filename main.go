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
	router := mux.NewRouter()
	staticHandler := http.FileServer(http.Dir(""))
	http.Handle("/style.css", staticHandler)
	router.HandleFunc("/", landingHandler)
	router.HandleFunc("/{username:[a-z0-9_-]+}", achievementHandler)
	http.Handle("/", router)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.WebserverPort), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

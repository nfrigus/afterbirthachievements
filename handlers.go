package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// landingHandler handles requests to "/"
func landingHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/landing.html")
	t.Execute(w, nil)
}

// landingHandler handles requests of the form "/{username}".
func achievementHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	userID, err := getUserID(username)
	log.Println(userID)
	if err != nil {
		t, _ := template.ParseFiles("tmpl/usernotfound.html")
		t.Execute(w, nil)
		return
	}
	unearned, err := unearnedAchievements(userID)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	categorized := categorizeAchievements(unearned)
	t, _ := template.ParseFiles("tmpl/achievements.html")
	t.Execute(w, categorized)
}

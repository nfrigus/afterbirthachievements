package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

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

	// Did the user supply their SteamID?
	userID, err := strconv.Atoi(username)
	if err != nil || len(username) != 17 {
		// Nope, that must have been their profile name. Let's try
		// that instead.
		userID, err = getUserID(username)
	}

	if err != nil {
		log.Println(err)
		t, _ := template.ParseFiles("tmpl/usernotfound.html")
		t.Execute(w, nil)
		return
	}
	unearned, err := unearnedAchievements(userID)
	if err != nil {
		log.Println(err)
		t, _ := template.ParseFiles("tmpl/usernotfound.html")
		t.Execute(w, nil)
		return
	}
	categorized := categorizeAchievements(unearned)
	t, _ := template.ParseFiles("tmpl/achievements.html")
	t.Execute(w, categorized)
}

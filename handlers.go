package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// landingHandler handles requests to "/"
func landingHandler(w http.ResponseWriter, r *http.Request) {
	renderContent("tmpl/landing.html", r, w, nil)
}

// achievementHandler handles requests of the form "/{username}".
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
		renderContent("tmpl/usernotfound.html", r, w, nil)
		return
	}
	unearned, err := unearnedAchievements(userID)
	if err != nil {
		log.Println(err)
		renderContent("tmpl/usernotfound.html", r, w, nil)
		return
	}
	categorized := categorizeAchievements(unearned)
	renderContent("tmpl/achievements.html", r, w, categorized)
}

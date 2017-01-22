package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"
)

const achievementAPIURL = "https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?appid=250900&key=%s&steamid=%d"
const userIDAPIURL = "http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s"

type SteamAchievementResponse struct {
	Playerstats struct {
		SteamID      string `json:"steamID"`
		GameName     string `json:"gameName"`
		Achievements []struct {
			Name     string `json:"name"`
			Achieved int    `json:"achieved"`
		} `json:"achievements"`
		Stats []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"stats"`
	} `json:"playerstats"`
}

func readSteamStats(steamID int) (steamAchievementResponse SteamAchievementResponse, err error) {
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(fmt.Sprintf(achievementAPIURL, config.SteamAPIKey, steamID))
	if err != nil {
		return
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&steamAchievementResponse)
	return
}

func unearnedAchievements(steamID int) (unearnedAchievements []Achievement, err error) {
	steamAchievementResponse, err := readSteamStats(steamID)
	if err != nil {
		return
	}
	var earnedAchievements []int
	// The achievements as obtained from Steam are not sorted, so we do that first
	// to make them easier to manipulate.
	for _, achievement := range steamAchievementResponse.Playerstats.Achievements {
		achievementID, _ := strconv.Atoi(achievement.Name)
		earnedAchievements = append(earnedAchievements, achievementID)
	}
	sort.Ints(earnedAchievements)
	earnedAchievements = append(earnedAchievements, 340)

	// Find the complement of the slice
	previousAchievement := 0
	for _, thisAchievement := range earnedAchievements {
		for i := previousAchievement + 1; i < thisAchievement; i++ {
			achievement := getAchievementByID(i)
			unearnedAchievements = append(unearnedAchievements, achievement)
		}
		previousAchievement = thisAchievement
	}
	return
}

type SteamPlayerIDResponse struct {
	Response struct {
		Steamid string `json:"steamid"`
		Success int    `json:"success"`
	} `json:"response"`
}

func getUserID(username string) (userID int, err error) {
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(fmt.Sprintf(userIDAPIURL, config.SteamAPIKey, username))
	if err != nil {
		return
	}
	defer response.Body.Close()
	steamResponse := SteamPlayerIDResponse{}
	err = json.NewDecoder(response.Body).Decode(&steamResponse)
	return strconv.Atoi(steamResponse.Response.Steamid)
}

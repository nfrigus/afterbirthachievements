package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Achievement represents the information pertaining to a given Afterbirth+ achievement
type Achievement struct {
	Name        string
	Category    int
	Description string
	Icon        string
	UnlockedBy  string
}

// Global variables to cache information about achievements and their categories.
var allCategories map[int]string
var allAchievements map[int]Achievement

// readAllAchievements reads from disk data about all achievement categories and stores the information
// in the global variable `allCategories`.
func readAllCategories() {
	allCategories = make(map[int]string)
	file, _ := os.Open("categories.csv")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	firstLine := true
	for scanner.Scan() {
		if firstLine {
			firstLine = false
			continue
		}
		line := strings.Split(scanner.Text(), ";")
		id, _ := strconv.Atoi(line[0])
		name := line[1]
		allCategories[id] = name
	}
}

// readAllAchievements reads from disk data about all achievements and stores the information
// in the global variable `allAchievements`.
func readAllAchievements() {
	allAchievements = make(map[int]Achievement)
	file, _ := os.Open("achievements.csv")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	firstLine := true
	for scanner.Scan() {
		if firstLine {
			firstLine = false
			continue
		}
		line := strings.Split(scanner.Text(), ";")
		id, _ := strconv.Atoi(line[0])
		category, _ := strconv.Atoi(line[1])
		description := line[2]
		name := line[3]
		icon := line[4]
		unlockedBy := line[5]
		achievement := Achievement{Name: name, Category: category, Description: description, Icon: icon, UnlockedBy: unlockedBy}
		allAchievements[id] = achievement
	}
}

// getAchievementByID finds the achievement with a given id (as an integer between 1 and 339)
func getAchievementByID(id int) Achievement {
	return allAchievements[id]
}

// categorizeAchievements takes a slice of achievements and puts them in their appropriate categories.
func categorizeAchievements(achievements []Achievement) map[string][]Achievement {
	result := make(map[string][]Achievement)
	for _, achievement := range achievements {
		cat := allCategories[achievement.Category]
		result[cat] = append(result[cat], achievement)
	}
	return result
}

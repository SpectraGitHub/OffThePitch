package handlers

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/dbaxs"
	"Integrasjonsprosjekt/internal"
	"encoding/json"
	"log"
	"net/http"
)

// Function to create a new league
func CreateLeagueHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "POST" && r.Method != "OPTIONS" {
		println("Invalid request. Expected POST, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "OPTIONS" {
		return
	}

	// Get user ID from token in header
	userid := GetUserId(w, r)
	if userid == 0 {
		http.Error(w, "User token not verified", http.StatusMethodNotAllowed)
		return
	}

	var league constants.LeagueInfo
	var player constants.LeaguePlayer
	player.UserId = userid
	// Add the player creating the league to the new league
	league.Players = append(league.Players, player)

	// Decode the leaguename into the league class
	err := json.NewDecoder(r.Body).Decode(&league)
	if err != nil {
		log.Fatal(err)
	}

	// If the leaguename is too short, return an error
	if len(league.LeagueName) < 5 {
		// Send error
		http.Error(w, "League name too short", http.StatusBadRequest)
		println("League name too short")
		return
	}

	// Creating the league in the database
	Error, league := dbaxs.CreateLeague(league)
	if Error != constants.No_error {
		println(internal.PrintError(Error))
		// Return error code to frontend
		http.Error(w, "Internal fail", 404)
	}

	// Write status created and the info to the new league as a response
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, league)
}

// Function for a user to join a new league
func JoinLeagueHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "POST" && r.Method != "OPTIONS" {
		println("Invalid request. Expected POST, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "OPTIONS" {
		return
	}

	var league constants.LeagueInfo
	var user constants.LeaguePlayer

	// Get user ID from token in header
	user.UserId = GetUserId(w, r)
	if user.UserId == 0 {
		return
	}

	// Add player to a new league
	league.Players = append(league.Players, user)

	// Decode the league code into league class
	err := json.NewDecoder(r.Body).Decode(&league)
	if err != nil {
		log.Fatal(err)
	}

	// If the league code does not exist, return error
	if !dbaxs.LeagueCodeIsTaken(league.LeagueCode) {
		http.Error(w, "League code does not exist", 404)
		println("League with the given code does not exist: ", league.LeagueCode)
		return
	}

	// Join league in database, returns league with all its players
	errorcode, league := dbaxs.JoinLeague(league)

	if errorcode != constants.No_error {
		println(internal.PrintError(errorcode))
	}

	// Write status created and the full league with all players as response
	w.WriteHeader(http.StatusCreated)
	writeResponse(w, league)
}

// Function to get all the leagues containing a specified player
func GetLeaguesHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "GET" && r.Method != "OPTIONS" {
		println("Invalid request. Expected GET, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "OPTIONS" {
		return
	}

	var user constants.UserInfo

	// Get user ID from token in header
	user.UserID = GetUserId(w, r)

	// Get a list from the database of all the leagues the given player is a part of
	leagues, errorcode := dbaxs.GetAllLeagues(user.UserID)

	if errorcode != constants.No_error {
		writeResponse(w, internal.PrintError(errorcode))
	}

	// Write all leagues as a response
	writeResponse(w, leagues)
}

// Function to leave a league
func LeaveLeagueHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var league constants.LeagueInfo

	if r.Method != "POST" && r.Method != "OPTIONS" {
		println("Invalid request. Expected POST, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "OPTIONS" {
		return
	}

	// Decode league ID into new league object
	err := json.NewDecoder(r.Body).Decode(&league)
	if err != nil {
		log.Fatal(err)
	}

	var user constants.LeaguePlayer
	// Get user ID from token in header
	user.UserId = GetUserId(w, r)
	if user.UserId == 0 {
		return
	}

	// Remove the userID from the given league
	errorcode := dbaxs.LeaveLeague(user.UserId, league.LeagueID)

	if errorcode != constants.No_error {
		println(internal.PrintError(errorcode))
	}

	// Write a response saying the league was left
	writeResponse(w, "League left.")
}

package handlers

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/dbaxs"
	"Integrasjonsprosjekt/internal"
	"encoding/json"
	"log"
	"net/http"
	"sort"
)

// Function to save a player's prediction for medal winners of the tournament
func SaveMedalPredictionHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var preds constants.MedalPredictions

	var err constants.Errorcode
	if r.Method != "POST" && r.Method != "OPTIONS" {
		println("Invalid request. Expected POST, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use POST with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "POST" {
		err, preds = saveMedalPredictions(w, r)
	} else {
		return
	}

	if err != constants.No_error {
		http.Error(w, internal.PrintError(err), 404)
	} else {
		// Write the medal predictions as a response
		writeResponse(w, preds)
	}
}

// Function to handle saving predictions for a user in database
func saveMedalPredictions(w http.ResponseWriter, r *http.Request) (constants.Errorcode, constants.MedalPredictions) {
	var preds constants.MedalPredictions
	err := json.NewDecoder(r.Body).Decode(&preds)
	if err != nil {
		log.Fatal(err)
	}

	// Get user ID from token in header
	preds.UserID = GetUserId(w, r)

	if preds.UserID == 0 {
		return constants.Unauthorized, preds
	}

	// If any of the predictions are not set, return error
	if preds.Gold == 0 || preds.Silver == 0 || preds.Bronze == 0 {
		return constants.WrongInput, preds
	}

	// Save predictions to database and return the results
	return dbaxs.SaveMedalPredictions(preds), preds
}

// Function to get a users medal predictions
func GetMedalPredictionHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var preds constants.MedalPredictions

	var err constants.Errorcode
	if r.Method != "GET" && r.Method != "OPTIONS" && r.Method != "POST" {
		println("Invalid request. Expected GET, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use POST with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "GET" {
		// If the method is get, the user is requesting their own predictions
		preds, err = getMedalPredictions(w, r)
	} else if r.Method == "POST" {
		// If the method is post, it might be a user requesting someone else's predictions
		// This should only be possible if the time is passed the start of matchday 1
		if internal.GetTime() < constants.MATCHDAYDIVIDERS[1] {
			http.Error(w, "The predictions of other players can not be accessed before predictions are locked in", http.StatusForbidden)
			return
		}
		preds, err = getMedalPredictions(w, r)
	} else {
		return
	}

	if err != constants.No_error {
		http.Error(w, internal.PrintError(err), 404)
	} else {
		// Write the requested predictions as a response
		writeResponse(w, preds)
	}
}

// Function to handle getting medal predictions from database
func getMedalPredictions(w http.ResponseWriter, r *http.Request) (constants.MedalPredictions, constants.Errorcode) {
	var preds constants.MedalPredictions

	if r.Method == "GET" {
		// If the method is GET, the user is requesting their own predictions and the userID can be found in the token in the header
		preds.UserID = GetUserId(w, r)
	} else {
		// If not, the body will contain the user ID, and can be decoded into the medalpredictions object
		err := json.NewDecoder(r.Body).Decode(&preds)
		if err != nil {
			println("ERROR in getPredictions: " + err.Error())
		}
	}

	// If the user is not found return an error
	if preds.UserID == 0 {
		return preds, constants.Unauthorized
	}

	// Get predictions from the database and return the results
	return dbaxs.GetMedalPredictions(preds.UserID)
}

// Function to get the team names (used in order to set medal predictions for the users)
func GetTeamNamesHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var teams []constants.TeamInfo

	var err constants.Errorcode
	if r.Method != "GET" && r.Method != "OPTIONS" {
		println("Invalid request. Expected GET, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "GET" {
		// If the method is get, retrieve the teams from the database
		teams, err = dbaxs.GetTeams()
	} else {
		return
	}

	// Sort the teams alphabetically 
	sort.SliceStable(teams, func(i, j int) bool {
		return teams[i].Name < teams[j].Name
	})

	if err != constants.No_error {
		http.Error(w, internal.PrintError(err), 404)
	} else {
		// Write a list of the teams as a response
		writeResponse(w, teams)
	}
}

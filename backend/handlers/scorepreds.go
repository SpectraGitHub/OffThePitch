package handlers

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/dbaxs"
	"Integrasjonsprosjekt/internal"
	"encoding/json"
	"net/http"
)

// Function to save a players score predictions
func SavePredictionsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "POST" && r.Method != "OPTIONS" {
		println("Invalid request. Expected POST, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use POST with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "POST" {
		saveGamePredictions(w, r)
	}
}

// Function to handle a request to get a user's predicitons
func GetPredictionsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "GET" && r.Method != "OPTIONS" && r.Method != "POST" {
		println("Invalid request. Expected GET, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "GET" || r.Method == "POST" {
		getPredictions(w, r)
	}
}

// Function to save a users predictions to the database
func saveGamePredictions(w http.ResponseWriter, r *http.Request) {
	var preds constants.Predictions
	// Decode the predictions into the new predictions object
	err := json.NewDecoder(r.Body).Decode(&preds)
	if err != nil {
		println("ERROR in saveGamePredictions: " + err.Error())
	}

	// Retrieve the user ID from the token in the header
	preds.UserID = GetUserId(w, r)
	if preds.UserID == 0 {
		return
	}
	// Save the predictions to the database
	dbaxs.SavePredictions(preds)
}

// Function to get a specified user's predictions
func getPredictions(w http.ResponseWriter, r *http.Request) {
	var preds constants.Predictions
	if r.Method == "GET" {
		// If the method is get, the user is requesting their own predictions
		preds.UserID = GetUserId(w, r)
	} else {
		// If not, someone is trying to access another player's predictions
		// This should only be possible after the first matchday has started
		if internal.GetTime() > constants.MATCHDAYDIVIDERS[1] {
			err := json.NewDecoder(r.Body).Decode(&preds)
			if err != nil {
				println("ERROR in getPredictions: " + err.Error())
			}
		}
	}
	if preds.UserID == 0 {
		return
	}

	// Get the predictions from the database
	preds = dbaxs.GetPredictions(preds.UserID)
	// Sort the predictions after the time the matches are played
	preds = internal.SortPredictions(preds)

	if r.Method == "POST" {
		// If someone is accessing another player's predictions, sort out all where their round is yet to start
		preds = filterPreds(preds)
	}

	// Write the requested predictions as a response
	writeResponse(w, preds)
}


// Filters out all predictions that have not been "locked in" due to the matchday starting
func filterPreds(preds constants.Predictions) constants.Predictions {

	for i, pred := range preds.Predictions {
		for _, match := range internal.Fixtures.Matches {
			if pred.MatchID == match.MatchID {
				if !match.Status.RoundStarted {
					preds.Predictions = preds.Predictions[:i]
					return preds
				}
			}
		}
	}

	return preds
}

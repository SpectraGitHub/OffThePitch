package handlers

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/dbaxs"
	"Integrasjonsprosjekt/internal"
	"encoding/json"
	"log"
	"net/http"
)

// Function to handle a request to save top scorers for a user
func SaveTopScorerPredictionHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var err constants.Errorcode
	var preds constants.TopScorerPredictions

	if r.Method != "POST" && r.Method != "OPTIONS" {
		println("Invalid request. Expected POST or GET, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use POST or GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "POST" {
		preds, err = saveTopScorerPredictions(w, r)
	} else if r.Method == "OPTIONS" {
		return
	}

	if err == constants.No_error {
		// Wite the requested topscorers as a response
		writeResponse(w, preds)
	} else {
		println(internal.PrintError(err))
		http.Error(w, internal.PrintError(err), 404)
	}
}

// Function to save top scorers for a user
func saveTopScorerPredictions(w http.ResponseWriter, r *http.Request) (constants.TopScorerPredictions, constants.Errorcode) {
	var preds constants.TopScorerPredictions
	// Decode the topscorers into an empty topscorerpredictions object
	err := json.NewDecoder(r.Body).Decode(&preds)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the userID from the token in the header
	preds.UserID = GetUserId(w, r)

	if preds.UserID == 0 {
		return preds, constants.Unauthorized
	}

	// Check that three topscorers are chosen
	if len(preds.Topscorers) != 3 {
		return preds, constants.WrongInput
	}

	// Save the topscorers to the database and return the response
	return preds, dbaxs.SaveTopScorerPredictions(preds)
}

// Function to handle requests to retrieve information about a player's chosen topscorers
func GetTopScorerPredictionHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var err constants.Errorcode
	var preds constants.TopScorerPredictions

	if r.Method != "GET" && r.Method != "OPTIONS" && r.Method != "POST" {
		println("Invalid request. Expected GET, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "GET" {
		// If the method is get, a player is trying to request their own topscorers
		preds, err = getTopScorerPredictions(w, r)
	} else if r.Method == "POST" {
		// If not, a player is requesting someone else's topscorers
		// This should only be possible after the first matchday has started
		if internal.GetTime() < constants.MATCHDAYDIVIDERS[1] {
			http.Error(w, "The predictions of other players can not be accessed before predictions are locked in", http.StatusForbidden)
			return
		}
		preds, err = getTopScorerPredictions(w, r)
	} else {
		return
	}

	if err == constants.No_error || err == constants.NoDataSet {
		// Write the predictions as a response
		writeResponse(w, preds)
	} else {
		println(internal.PrintError(err))
		http.Error(w, internal.PrintError(err), 404)
	}

}

// Fuction to retrieve topscorer predictions for a specified user
func getTopScorerPredictions(w http.ResponseWriter, r *http.Request) (constants.TopScorerPredictions, constants.Errorcode) {
	var preds constants.TopScorerPredictions
	if r.Method == "GET" {
		// If the method is get, a player is trying to request their own topscorers
		preds.UserID = GetUserId(w, r)
	} else {
		// If not the userID can be decoded from the body
		err := json.NewDecoder(r.Body).Decode(&preds)
		if err != nil {
			println("ERROR in getPredictions: " + err.Error())
		}
	}

	if preds.UserID == 0 {
		return preds, constants.Unauthorized
	}

	// Get the topscorers for the specified player from the database and return the response
	return dbaxs.GetTopscorers(preds.UserID)
}

// Function to get all players (Used to know what players are available to choose from when choosing topscorers)
func GetPlayersHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var Teams constants.AllPlayers
	var err constants.Errorcode

	if r.Method != "GET" && r.Method != "OPTIONS" {
		println("Invalid request. Expected GET, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "OPTIONS" {
		return
	} else {
		Teams, err = getAllPlayers()
	}
	if err != constants.No_error {
		http.Error(w, internal.PrintError(err), 404)
	} else {
		// Write all the teams containing each player as a response
		writeResponse(w, Teams)
	}

}

// Function to get all players from the database
func getAllPlayers() (constants.AllPlayers, constants.Errorcode) {
	var err constants.Errorcode
	var players []constants.Player
	var teams []constants.TeamInfo
	var allPlayers constants.AllPlayers

	// Retrieve all the players from the database
	players, err = dbaxs.GetAllPlayers()
	if err != constants.No_error {
		println(internal.PrintError(err))
		return allPlayers, err
	}

	// Retrieve all the teams from the database
	teams, err = dbaxs.GetTeams()
	if err != constants.No_error {
		println(internal.PrintError(err))
		return allPlayers, err
	}

	// Append each player to their corresponding team
	var teamBuf constants.Team
	for i := range teams {
		for y := range players {
			if teams[i].TeamID == players[y].TeamID {
				teamBuf.Players = append(teamBuf.Players, players[y])
			}
		}
		teamBuf.TeamID = teams[i].TeamID
		teamBuf.TeamName = internal.TeamMap[teamBuf.TeamID]
		allPlayers.Teams = append(allPlayers.Teams, teamBuf)
		teamBuf.Players = nil
	}

	// Sort the players alphabetically
	allPlayers = internal.SortAllPlayers(allPlayers)

	// Return the players
	return allPlayers, constants.No_error
}

package dbaxs

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"log"
)

const SQL_ADD_MEDAL_PREDICTION = `INSERT INTO medalje_valg (bruker_id, bronse, solv, gull) VALUES (?, ?, ?, ?);`
const SQL_GET_MEDAL_PREDICTIONS = `SELECT bronse, solv, gull FROM medalje_valg WHERE bruker_id = ?`
const SQL_CLEAR_MEDAL_PREDICTIONS = `DELETE FROM medalje_valg WHERE bruker_id = ?`
const SQL_GET_TEAMS = `SELECT * FROM lagnavn`

// Get medal predictions for the given user
func GetMedalPredictions(userID int) (constants.MedalPredictions, constants.Errorcode) {
	var preds constants.MedalPredictions

	if !checkUserID(userID) { // If user doesn't exist return errorcode
		return preds, constants.InvalidUsername
	}

	preds, err := getMedalPredictions(userID) // Fetch the users medalpredictions

	return preds, err
}

// Save medalpredictions for the user
func SaveMedalPredictions(pred constants.MedalPredictions) constants.Errorcode {
	if !checkUserID(pred.UserID) { // Check that user exists
		return constants.InvalidUsername
	}
	clearMedalPredictions(pred.UserID) // Clear old medalpredictions

	// Insert meddal predictions to DB
	err := DB.QueryRow(SQL_ADD_MEDAL_PREDICTION, pred.UserID, pred.Bronze, pred.Silver, pred.Gold).Err()
	if err != nil {
		log.Println("ERROR:", err.Error())
		return constants.InternalFail
	}

	return constants.No_error
}

// Clear users medal predictions from DB
func clearMedalPredictions(userID int) error {
	return DB.QueryRow(SQL_CLEAR_MEDAL_PREDICTIONS, userID).Err()
}

// Fetch a user's medal predictions
func getMedalPredictions(userID int) (pred constants.MedalPredictions, errorCode constants.Errorcode) {
	pred.UserID = userID
	tx, err := DB.Begin()
	if err != nil {
		log.Printf("ERROR, Failed to begin transaction, error: %v", err.Error())
		return pred, constants.InternalFail
	}
	// Fetch users medal predictions from DB
	tx.QueryRow(SQL_GET_MEDAL_PREDICTIONS, userID).Scan(&pred.Bronze, &pred.Silver, &pred.Gold)

	err = tx.Commit()
	if err != nil { // Log error
		log.Printf("ERROR, unable to commit getMedalPredictions transaction, error: %v", err.Error())
		return pred, constants.No_error
	}

	return pred, constants.No_error
}

// Fetch all the teams
func GetTeams() ([]constants.TeamInfo, constants.Errorcode) {
	teams, err := getAllTeams() // Fetch all the teams

	for i := range teams { // Update the TeamMap with the teams in a global hashmap
		internal.TeamMap[teams[i].TeamID] = teams[i].Name
	}

	if err != constants.No_error {
		println("Error getting teams from database: ", internal.PrintError(err))
	}

	return teams, err
}

// Fetch all the teams from the DB
func getAllTeams() ([]constants.TeamInfo, constants.Errorcode) {
	var teams []constants.TeamInfo
	var newTeam constants.TeamInfo

	rows, err := DB.Query(SQL_GET_TEAMS) // Get all teams from DB
	if err != nil {
		log.Println(err.Error())
		return teams, constants.InternalFail
	}

	for rows.Next() { // Iterate through rows and scan team information
		if err := rows.Scan(&newTeam.TeamID, &newTeam.Name); err != nil {
			log.Fatal(err.Error())
		}
		teams = append(teams, newTeam) // Add each team information to a list
	}
	rows.Close()

	return teams, constants.No_error // Return list of teams
}

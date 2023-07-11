package dbaxs

/*
*	Database functions connected to submitting and getting topscorer predictions
 */

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"database/sql"
	"log"
)

const SQL_ADD_TOPSCORER_PREDICTION = `INSERT INTO toppscorer (bruker_id, spiller_id) VALUES (?, ?);`
const SQL_GET_TOPSCORER_PREDICTIONS = `
	SELECT 
		spiller.spiller_id, spiller.fornavn, spiller.etternavn, spiller.lag_id, lagnavn.lag_navn, spiller.maal 
	FROM toppscorer 
		INNER JOIN spiller on toppscorer.spiller_id = spiller.spiller_id 
		INNER JOIN lagnavn on spiller.lag_id = lagnavn.lag_id 
	WHERE bruker_id = ?;
	`

const SQL_CLEAR_TOPSCORER_PREDICTIONS = `DELETE FROM toppscorer WHERE bruker_id = ?`
const SQL_GET_ALL_PLAYERS = `SELECT spiller_id, fornavn, etternavn, lag_id FROM spiller;`

const SQL_ADD_GOAL = "UPDATE spiller SET maal = maal + 1 WHERE spiller_id = ?;"

/*
 * Fetches the users topscorer predictions and returns them along with an errorcode
 */
func GetTopscorers(userID int) (predictions constants.TopScorerPredictions, errorCode constants.Errorcode) {
	var prediction constants.TopScorer
	var last_name string

	tx, err := DB.Begin()
	if err != nil {
		log.Printf("ERROR, Failed to begin transaction, error: %v", err.Error())
		return predictions, constants.InternalFail
	}

	// Fetch the users topscorer rows
	rows, err := tx.Query(SQL_GET_TOPSCORER_PREDICTIONS, userID)
	if err != nil {
		log.Printf("ERROR(dbaxs.go:362): %v", err.Error())
		return predictions, constants.InternalFail
	}
	for rows.Next() { // Iterate each topscorer and scan player
		if err := rows.Scan(&prediction.Player_id, &prediction.Player_name, &last_name, &prediction.Team_id, &prediction.Team_name, &prediction.Goals); err != nil {
			tx.Rollback()
			return predictions, constants.InternalFail
		}
		if prediction.Player_name != "" { // If player has first name, add space for last name (not all players have first name)
			prediction.Player_name += " "
		}
		prediction.Player_name += last_name
		predictions.Topscorers = append(predictions.Topscorers, prediction) // Append player to list of topscorers
	}
	err = tx.Commit() // Commit and close connection
	if err != nil {
		log.Printf("ERROR, unable to commit GetTopscorers transaction, error: %v", err.Error())
		return predictions, constants.InternalFail
	}

	predictions.UserID = userID
	return predictions, constants.No_error // Return the topscorer predictions for user
}

// Save topscorers for specified user
func SaveTopScorerPredictions(pred constants.TopScorerPredictions) constants.Errorcode {
	if !checkUserID(pred.UserID) { // Return if user doesn't exist
		return constants.InvalidUsername
	}

	clearTopScorerPredictions(pred.UserID) // Clear all the old topscorers
	tx, err := DB.Begin()
	if err != nil {
		log.Printf("ERROR, Failed to begin transaction, error: %v", err.Error())
	}
	for i := range pred.Topscorers {
		// Add each topscorer to topscorer table in DB
		_, err = tx.Exec(SQL_ADD_TOPSCORER_PREDICTION, pred.UserID, pred.Topscorers[i].Player_id)
		if err != nil {
			log.Println("ERROR:", err.Error())
			tx.Rollback()
			return constants.InternalFail
		}
	}
	tx.Commit() // Commit all transactions and close connection
	return constants.No_error
}

// Clear ALL topscorers for user
func clearTopScorerPredictions(userID int) error {
	return DB.QueryRow(SQL_CLEAR_TOPSCORER_PREDICTIONS, userID).Err()
}

// Get all players from DB
func getAllPlayers() ([]constants.Player, constants.Errorcode) {
	var player constants.Player
	var players []constants.Player

	rows, err := DB.Query(SQL_GET_ALL_PLAYERS) // Fetch all the players in DB (football players, not users)
	if err != nil {
		log.Println(err.Error())
	}
	for rows.Next() { // Iterate each player row, and scan into buffer
		if err := rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.TeamID); err != nil {
			log.Fatal(err.Error())
		}
		players = append(players, player) // Append buffer to players
	}
	rows.Close() // Close connection

	if len(players) == 0 { // If no players fetched, return error
		return players, constants.NoDataSet
	}

	return players, constants.No_error
}

// Get all world cup players
func GetAllPlayers() ([]constants.Player, constants.Errorcode) {
	var players []constants.Player
	var err = constants.No_error

	players, err = getAllPlayers() // Fetch players from DB

	if err != constants.No_error {
		println(internal.PrintError(err))
	}

	return players, err
}

// Add goals scored by player to DB
func addGoal(tx *sql.Tx, playerid int) constants.Errorcode {
	println("Add goal")
	_, err := tx.Exec(SQL_ADD_GOAL, playerid) // Update player row with goal scored

	if err != nil {
		println("ERROR in dbaxs.addGoal", err.Error())
		return constants.InternalFail
	}

	return constants.No_error
}

package dbaxs

import (
	"Integrasjonsprosjekt/constants"
	"database/sql"
	"log"
)

const SQL_GET_PREDICTIONS = `SELECT kamp_id, hjemme_score, borte_score, poeng FROM kamp_valg WHERE bruker_id = ?`
const SQL_ADD_PREDICTION = `REPLACE INTO kamp_valg (bruker_id, kamp_id, hjemme_score, borte_score, poeng) 
							VALUES (?, ?, ?, ?, 0);`
const SQL_CLEAR_PREDICTIONS = `DELETE FROM kamp_valg WHERE bruker_id = ?`

// Get all the predictions for the specified user
func GetPredictions(userID int) constants.Predictions {
	var preds []constants.Prediction
	var prediction constants.Prediction
	var predictionReturn constants.Predictions

	tx, err := DB.Begin()
	if err != nil {
		log.Printf("ERROR, Failed to begin transaction, error: %v", err.Error())
	}

	// Fetch all the rows from the given user
	rows, err := tx.Query(SQL_GET_PREDICTIONS, userID)
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return predictionReturn
	}

	// iterate through the rows and scan each prediction
	for rows.Next() {
		if err := rows.Scan(&prediction.MatchID, &prediction.HomeScore, &prediction.AwayScore, &prediction.Points); err != nil {
			log.Printf("Failed to fetch prediction, errror: %v", err.Error())
			tx.Rollback()
			return predictionReturn
		}
		preds = append(preds, prediction) // Append each match's prediction
	}
	rows.Close()
	tx.Commit()

	predictionReturn.Predictions = preds
	predictionReturn.UserID = userID

	return predictionReturn // Return all the predictions
}

/*
 * Save a user's predictions
 */
func SavePredictions(preds constants.Predictions) {
	tx, err := DB.Begin() // Begin transaction
	if err != nil {
		log.Printf("ERROR, Failed to begin transaction, error: %v", err.Error())
		return
	}
	for i := range preds.Predictions { // Save each match under the same transaction
		savePrediction(tx, preds.UserID, preds.Predictions[i])
	}
	err = tx.Commit() // Commit transaction and close connection.
	if err != nil {
		log.Printf("ERROR, unable to commit savePredictions transactions, error: %v", err.Error())
	}
}

/*
 *	Saves a prediction to the database
 */
func savePrediction(tx *sql.Tx, userID int, prediction constants.Prediction) {
	// Replace into (update or insert) into DB
	_, err := tx.Exec(SQL_ADD_PREDICTION, userID, prediction.MatchID, prediction.HomeScore, prediction.AwayScore)
	if err != nil {
		log.Println("ERROR in savePrediction:", err.Error())
	}
}

// Clear ALL match predictions
func clearMatchPredictions(userID int) error {
	// Clear all the match predictions from DB for user
	err := DB.QueryRow(SQL_CLEAR_PREDICTIONS, userID).Err()
	if err != nil {
		return err
	}
	return nil
}

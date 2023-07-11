package dbaxs

import (
	"Integrasjonsprosjekt/constants"
	"database/sql"
	"log"
)

// Calculate points for a single match
func CalculatePoints(match constants.Match) {
	if pointsAreAdded(match) { // return if points already added in DB
		println("Points already added, but tried to be added again for: " + match.HomeTeam.Name + " - " + match.AwayTeam.Name)
		return
	}
	tx, err := DB.Begin() // Create a single transaction per match
	if err != nil {
		log.Printf("failed to begin transaction, error: %v", err.Error())
	}

	println("Adding points for match: " + match.HomeTeam.Name + " - " + match.AwayTeam.Name)
	// Add points for match result
	addPointsGoalDifference(tx, match)
	addPointsHomeAndAwayGoals(tx, match)
	addPoints1X2(tx, match)
	addPointsNumGoals(tx, match)

	// Add points for home goals
	for _, goal := range match.HomeTeam.Goals {
		addGoal(tx, goal.Player.PlayerID)
		addPointsTopscorer(tx, constants.TOPSCORERGOALPOINTS, goal.Player.PlayerID)
	}

	// Add points for away goals
	for _, goal := range match.AwayTeam.Goals {
		addGoal(tx, goal.Player.PlayerID)
		addPointsTopscorer(tx, constants.TOPSCORERGOALPOINTS, goal.Player.PlayerID)
	}
	tx.Commit()
}

const SQL_ADD_POINTS_FOR_GOAL_DIFFERENCE = `
	UPDATE
		bruker 
		INNER JOIN kamp_valg
		ON bruker.bruker_id = kamp_valg.bruker_id
	SET
		bruker.poeng = bruker.poeng + ?,
		kamp_valg.poeng = kamp_valg.poeng + ?
	WHERE
		kamp_valg.kamp_id = ?
		AND
		kamp_valg.hjemme_score - kamp_valg.borte_score = ?;
`

// Add points where user guessed correct goal difference
func addPointsGoalDifference(tx *sql.Tx, match constants.Match) {
	// Add points to match in DB IF (home score - away score) predicted is same as result
	_, err := tx.Exec(SQL_ADD_POINTS_FOR_GOAL_DIFFERENCE, constants.GOALDIFFPOINTS, constants.GOALDIFFPOINTS, match.MatchID, (match.HomeTeam.Score - match.AwayTeam.Score))

	if err != nil {
		println("Error in addPointsForGoalDifference:", err.Error())
	}
}

var SQL_ADD_POINTS_FOR_HOME_GOALS = `
	UPDATE 
		bruker	
	INNER JOIN
		kamp_valg 
		ON
			bruker.bruker_id = kamp_valg.bruker_id 
	SET	
		bruker.poeng = bruker.poeng + ?, 
		kamp_valg.poeng = kamp_valg.poeng + ? 
	WHERE 
		kamp_valg.kamp_id = ?
		AND 
		kamp_valg.hjemme_score = ?;
`
var SQL_ADD_POINTS_FOR_AWAY_GOALS = `
	UPDATE
		bruker
		INNER JOIN kamp_valg
		ON bruker.bruker_id = kamp_valg.bruker_id
	SET
		bruker.poeng = bruker.poeng + ?,
		kamp_valg.poeng = kamp_valg.poeng + ?
	WHERE
		kamp_valg.kamp_id = ?
		AND
		kamp_valg.borte_score = ?;
`

// Add points for correct home and away goals predicted
func addPointsHomeAndAwayGoals(tx *sql.Tx, match constants.Match) {
	// Add points to DB where user guessed correct home goals
	_, err := tx.Exec(SQL_ADD_POINTS_FOR_HOME_GOALS, constants.HOMEANDAWAYGOALPOINTS, constants.HOMEANDAWAYGOALPOINTS, match.MatchID, match.HomeTeam.Score)

	if err != nil {
		println("Error in addPointsForHome:", err.Error())
	}
	// Add points to DB where user guessed correct away goals
	_, err = tx.Exec(SQL_ADD_POINTS_FOR_AWAY_GOALS, constants.HOMEANDAWAYGOALPOINTS, constants.HOMEANDAWAYGOALPOINTS, match.MatchID, match.AwayTeam.Score)

	if err != nil {
		println("Error in addPointsForAway:", err.Error())
	}
}

var SQL_ADD_POINTS_FOR_1X2 = `
	UPDATE
		bruker
		INNER JOIN kamp_valg
		ON bruker.bruker_id = kamp_valg.bruker_id
	SET
		bruker.poeng = bruker.poeng + ?,
		kamp_valg.poeng = kamp_valg.poeng + ?
	WHERE
		kamp_valg.kamp_id = ?
		AND
		kamp_valg.hjemme_score 
`

// Add points where user guessed the correct winnner (or draw)
func addPoints1X2(tx *sql.Tx, match constants.Match) {
	var symbol string
	// Symbol is set to whether home team or away team won or draw.
	if match.HomeTeam.Score > match.AwayTeam.Score {
		symbol = ">"
	} else if match.HomeTeam.Score < match.AwayTeam.Score {
		symbol = "<"
	} else if match.HomeTeam.Score == match.AwayTeam.Score {
		symbol = "="
	} else {
		println("ERROR in addpoints1X2, symbol ambiguous")
	}

	// Create query which checks if predicted result <>= to actual result
	var query = SQL_ADD_POINTS_FOR_1X2 + symbol + " kamp_valg.borte_score;"

	// Run query adding points for all the users with correct 1x2 prediction
	_, err := tx.Exec(query, constants.POINTS1X2, constants.POINTS1X2, match.MatchID)

	if err != nil {
		println("ERROR in addPoints1X2:", err.Error())
	}
}

const SQL_ADD_POINTS_FOR_NUM_GOALS = `
	UPDATE
		bruker
		INNER JOIN kamp_valg
		ON bruker.bruker_id = kamp_valg.bruker_id
	SET
		bruker.poeng = bruker.poeng + ?,
		kamp_valg.poeng = kamp_valg.poeng + ?
	WHERE
		kamp_valg.kamp_id = ?
		AND
		kamp_valg.hjemme_score + kamp_valg.borte_score > ?;
`

// Add points for user guessing correctly that match would have over a certain amount of goals
func addPointsNumGoals(tx *sql.Tx, match constants.Match) {
	var numGoals int = match.HomeTeam.Score + match.AwayTeam.Score // Add number of goals in the match togheter
	var goalLimit, numpoints int

	if numGoals > 7 { // If more than 7 goals in the game
		goalLimit = 7
		numpoints = constants.MORETHAN7GOALS - constants.MORETHAN5GOALS - constants.MORETHAN3GOALS

		// Add points for every user guessing over 7 goals
		_, err := tx.Exec(SQL_ADD_POINTS_FOR_NUM_GOALS, numpoints, numpoints, match.MatchID, goalLimit)

		if err != nil {
			println("ERROR in AddpointsNumGoals:", err.Error())
		}

	}
	if numGoals > 5 { // If more than 5 goals in the game
		goalLimit = 5
		numpoints = constants.MORETHAN5GOALS - constants.MORETHAN3GOALS

		// Add points for every user guessing over 5 goals
		_, err := tx.Exec(SQL_ADD_POINTS_FOR_NUM_GOALS, numpoints, numpoints, match.MatchID, goalLimit)

		if err != nil {
			println("ERROR in AddpointsNumGoals:", err.Error())
		}
	}
	if numGoals > 3 { // If more than 5 goals in the game
		goalLimit = 3
		numpoints = constants.MORETHAN3GOALS

		// Add points for every user guessing over 3 goals
		_, err := tx.Exec(SQL_ADD_POINTS_FOR_NUM_GOALS, numpoints, numpoints, match.MatchID, goalLimit)

		if err != nil {
			println("ERROR in AddpointsNumGoals:", err.Error())
		}
	}
}

var SQL_ADD_POINTS_GOAL_SCORED = `
	UPDATE
		bruker
		INNER JOIN toppscorer
		ON bruker.bruker_id = toppscorer.bruker_id
	SET
		bruker.poeng = bruker.poeng + ?
	WHERE
		toppscorer.spiller_id = ?;
`

// Add points for a goalscorer in a match
func addPointsTopscorer(tx *sql.Tx, points int, playerID int) {
	// Add points to user with the given player as a chosen topscorer
	_, err := tx.Exec(SQL_ADD_POINTS_GOAL_SCORED, points, playerID)

	if err != nil {
		println("ERROR in addPointsTopscorer:", err.Error())
	}

	println("Added points for toppscorer:", playerID)
}

var SQL_ADD_POINTS_MEDAL_WINNERS = `
	UPDATE
		bruker
		INNER JOIN medalje_valg
		ON bruker.bruker_id = medalje_valg.bruker_id
	SET
		bruker.poeng = bruker.poeng + ?
	WHERE
		
`

var SQL_ADD_POINTS_TOPSCORER_OF_TOURNAMENT = `
	UPDATE
		bruker
		INNER JOIN toppscorer
		ON bruker.bruker_id = toppscorer.bruker_id
	SET
		bruker.poeng = bruker.poeng + 5
	WHERE
		toppscorer.spiller_id = ?
`

/*
*	!NB! - USERID IN THE GIVEN WINNERS OBJECT MUST BE THE PLAYERID OF THE TOURNAMENTS TOP SCORER
*
*	Adds points in the databse to the players who predicted the correct trophy winners and topscorer of the tournament
 */
func AddPointsMedalWinners(winners constants.MedalPredictions) {
	gold := "medalje_valg.solv = ? OR medalje_valg.bronse = ?;"
	silver := "medalje_valg.gull = ? OR medalje_valg.bronse = ?;"
	bronze := "medalje_valg.gull = ? OR medalje_valg.solv = ?;"

	tx, err := DB.Begin()
	if err != nil {
		log.Printf("failed to begin transaction, error: %v", err.Error())
		return
	}
	// ADDING POINTS FOR MEDAL WINNER BUT WRONG PLACING
	_, err = tx.Exec(SQL_ADD_POINTS_MEDAL_WINNERS+gold, constants.INCORRECTPOINTS, winners.Gold, winners.Gold)
	if err != nil {
		println("ERROR in adding Gold Winner points:", err.Error())
	}

	_, err = tx.Exec(SQL_ADD_POINTS_MEDAL_WINNERS+silver, constants.INCORRECTPOINTS, winners.Silver, winners.Silver)
	if err != nil {
		println("ERROR in adding Silver Winner points:", err.Error())
	}

	_, err = tx.Exec(SQL_ADD_POINTS_MEDAL_WINNERS+bronze, constants.INCORRECTPOINTS, winners.Bronze, winners.Bronze)
	if err != nil {
		println("ERROR in adding Bronze Winner points:", err.Error())
	}

	// ADDING POINTS FOR CORRECT MEDAL GUESSED
	gold = "medalje_valg.gull = ?;"
	silver = "medalje_valg.solv = ?;"
	bronze = "medalje_valg.bronse = ?;"

	// Add points for users with correct gold winner
	_, err = tx.Exec(SQL_ADD_POINTS_MEDAL_WINNERS+gold, constants.CORRECTTROPHYPOINTS, winners.Gold)
	if err != nil {
		println("ERROR in adding Gold Winner points:", err.Error())
	}
	// Add points for users with correct silver winner
	_, err = tx.Exec(SQL_ADD_POINTS_MEDAL_WINNERS+silver, constants.CORRECTTROPHYPOINTS, winners.Silver)
	if err != nil {
		println("ERROR in adding Silver Winner points:", err.Error())
	}
	// Add points for users with correct bronze winner
	_, err = tx.Exec(SQL_ADD_POINTS_MEDAL_WINNERS+bronze, constants.CORRECTTROPHYPOINTS, winners.Bronze)
	if err != nil {
		println("ERROR in adding Bronze Winner points:", err.Error())
	}

	// ADD POINTS FOR TOPSCORER OF TOURNAMENT
	addPointsTopscorer(tx, constants.TOPSCOREROFTOURNAMENTPOINTS, winners.UserID)
	tx.Commit()
}

// Checking if any points are added for a given match
const SQL_CHECK_POINTS = `
	SELECT 
		SUM(poeng)
	FROM 
		kamp_valg
	WHERE
		kamp_id = ?;

`

// Check whether points have been added for a given match
func pointsAreAdded(match constants.Match) bool {
	tx, err := DB.Begin()
	if err != nil {
		log.Printf("Failed to begin transaction, error: %v", err.Error())
	}
	row := tx.QueryRow(SQL_CHECK_POINTS, match.MatchID) // Fetch total points for given match
	SQLerror := row.Err()

	if SQLerror != nil {
		println("Error in pointsAreAdded:", SQLerror)

	}
	var poeng int
	row.Scan(&poeng) // Scan row
	tx.Commit()
	return poeng > 0 // If any points have been added return true
}

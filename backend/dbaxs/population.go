package dbaxs

/*
*	Helper functions to populate the database with teams and players
*/

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"fmt"
	"log"
	"strings"
)

const SQL_GET_EXISTING_TEAM = "SELECT (lag_navn) FROM lagnavn WHERE lag_id = ?;"
const SQL_ADD_NEW_TEAM = "INSERT INTO lagnavn (lag_id, lag_navn) VALUES (?, ?);"
const SQL_GET_EXISTING_PLAYER = "SELECT navn FROM spiller WHERE spiller_id = ?;"
const SQL_ADD_NEW_PLAYER = "INSERT INTO spiller (spiller_id, fornavn, etternavn, lag_id, maal) VALUES (?, ?, ?, 0);"
const SQL_ADD_MANY_PLAYERS = "INSERT INTO spiller (spiller_id, fornavn, etternavn, lag_id, maal) VALUES "

// Checks if teamID exists in DB
func TeamExists(teamID int) bool {
	var name string
	err := DB.QueryRow(SQL_GET_EXISTING_TEAM, teamID).Scan(&name) // Fetch team by ID and scan into name

	if err != nil {
		println(err.Error())
	}
	if name != "" { // If team successfully fetched, return true
		return true
	} else {
		return false
	}
}

// Save a team to DB
func SaveTeam(teamID int, teamName string) {
	if TeamExists(teamID) { // Return if team already exists
		return
	}

	// Insert team into DB
	err := DB.QueryRow(SQL_ADD_NEW_TEAM, teamID, teamName).Err()
	if err != nil {
		log.Println("ERROR(dbaxs.go:404):", err.Error())
	}
}

// Save a single player to DB
func SavePlayer(name string, playerID int, teamID int) constants.Errorcode {
	var existingPlayerName string
	// Try to fetch player name from DB
	err := DB.QueryRow(SQL_GET_EXISTING_PLAYER, playerID).Scan(&existingPlayerName)

	if err != nil {
		println(err.Error())
	}
	if existingPlayerName != "" { // Return if player already exists
		return constants.PrimaryKeyExists
	}

	// Insert player into DB
	err = DB.QueryRow(SQL_ADD_NEW_PLAYER, playerID, name, teamID).Err()
	if err != nil {
		log.Println("ERROR(dbaxs.go:422):", err.Error())
	}
	return constants.No_error
}

// Save multiple players. Uses valueQuery instead of player list
func SavePlayers(valueQuery string) {
	FullQuery := SQL_ADD_MANY_PLAYERS + valueQuery

	// Insert players into DB
	err := DB.QueryRow(FullQuery).Err()
	if err != nil {
		log.Println("ERROR(dbaxs.go:432):", err.Error())
	}
}

func SaveTeams() {
	for i := range internal.Fixtures.Matches {
		SaveTeam(internal.Fixtures.Matches[i].HomeTeam.TeamID, internal.Fixtures.Matches[i].HomeTeam.Name)
	}
}

//
// Functions below was used prior to World cup starting to populate DB with players
//
/*
	func FindPlayersFromMatchArray(matches []int) {
		for i := range matches {
			findAndSaveTeamPlayers(matches[i])
		}
	}

	func findAndSaveTeamPlayers(matchID int) {
		match := internal.GetMatchInfo(matchID)

		if TeamExists(match.HomeTeam.TeamID) {
			if len(match.HomeTeam.Lineup) > 0 {
				addTeamSQL := createTeamInsert(match.HomeTeam.Lineup, match.HomeTeam.TeamID)
				println("Adding " + match.HomeTeam.Name)
				SavePlayers(addTeamSQL)
			} else {
				println("Cold not find players from ", match.HomeTeam.Name)
			}

		}

		if TeamExists(match.AwayTeam.TeamID) {
			if len(match.AwayTeam.Lineup) > 0 {
				println("Adding " + match.AwayTeam.Name)
				addTeamSQL := createTeamInsert(match.AwayTeam.Lineup, match.AwayTeam.TeamID)
				SavePlayers(addTeamSQL)
			} else {
				println("Cold not find players from ", match.AwayTeam.Name)

			}
		}
	}
*/

// Create valueQuery for MYSQL with all the players from a team
func createTeamInsert(lineup []constants.PlayerInfo, teamID int) string {
	query := ""
	teamid := fmt.Sprintf("%d", teamID)

	// Append each player to the query
	for i := range lineup {
		firstName := strings.Replace(lineup[i].Player.FirstName, "'", "_", -1)
		lastName := strings.Replace(lineup[i].Player.LastName, "'", "_", -1)

		playerid := fmt.Sprintf("%d", lineup[i].Player.ID)
		query += "(" + playerid + ", '" + firstName + "', '" + lastName + "', " + teamid + ", 0), "
	}

	// Remove the last ","
	if len(query) > 2 {
		query = query[0 : len(query)-2]
	}

	return query
}

// Save a single player
func savePlayer(player constants.Player, teamID int) {
	// Valid checks
	if player.FirstName == "" || player.LastName == "" || player.ID == 0 {
		println("ERROR: Player info invalid")
		return
	}

	fullname := player.FirstName + " " + player.LastName
	// Save player to DB
	err := SavePlayer(fullname, player.ID, teamID)
	if err != constants.No_error {
		println("ERROR: ", internal.PrintError(err))
	}

}

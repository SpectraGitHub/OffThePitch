package dbaxs

/*
*		Database functions connected to creating, joining and getting leagues
 */

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"log"
	"math/rand"
	"time"
)

const SQL_CREATE_NEW_LEAGUE = "INSERT INTO liganavn (liga_navn, liga_kode) VALUES (?, ?);"
const SQL_GET_LEAGUE_ID = "SELECT (liga_id) FROM liganavn WHERE liga_kode = ?;"
const SQL_GET_LEAGUE_NAME_AND_CODE = "SELECT liga_navn, liga_kode FROM liganavn WHERE liga_id = ?;"
const SQL_JOIN_LEAGUE = "INSERT INTO liga (liga_id, bruker_id) VALUES (?, ?);"
const SQL_GET_LEAGUE_PARTICIPANTS = "select liga.bruker_id, bruker.brukernavn, poeng from liga inner join bruker on liga.bruker_id = bruker.bruker_id WHERE liga.liga_id = ?;"
const SQL_GET_USERS_LEAGUES = "SELECT liga_id FROM liga WHERE bruker_id = ?;"
const SQL_LEAVE_LEAGUE = "DELETE FROM liga WHERE liga_id = ? AND bruker_id = ?;"
const SQL_DELETE_LEAGUE = "DELETE FROM liganavn WHERE liga_id = ?;"

// Creates league with the given league name
func CreateLeague(league constants.LeagueInfo) (constants.Errorcode, constants.LeagueInfo) {
	league.LeagueCode = CreateLeagueCode()
	tx, err := DB.Begin()
	if err != nil {
		log.Printf("Error, failed to start transaction, error: %v", err.Error())
	}
	err = tx.QueryRow(SQL_CREATE_NEW_LEAGUE, league.LeagueName, league.LeagueCode).Err() // Add league to database

	if err != nil {
		log.Println("ERROR(dbaxs.go:250):", err.Error())
		return constants.InternalFail, league
	}

	tx.QueryRow(SQL_GET_LEAGUE_ID, league.LeagueCode).Scan(&league.LeagueID) // Fetch league id

	err = tx.QueryRow(SQL_JOIN_LEAGUE, league.LeagueID, league.Players[0].UserId).Err() // Add user to specified league
	if err != nil {
		log.Println("ERROR in createleague:", err.Error())
		return constants.InternalFail, league
	}
	err = tx.Commit() // Commit and close connection to DB
	if err != nil {
		log.Printf("Error, failed to commit transaction, error: %v", err.Error())
	}
	return constants.No_error, league
}

// Create random 6 letter/number league code
func CreateLeagueCode() string {
	rand.Seed(time.Now().Unix())
	var i int
	isTaken := true
	var leagueCode string

	for isTaken { // If code already taken, create new league code
		i = 0
		leagueCode = ""
		for i < 6 {
			leagueCode += randomChar(rand.Intn(62)) // Add random letter or number
			i++
		}
		isTaken = LeagueCodeIsTaken(leagueCode)
	}

	return leagueCode
}

// Check database if league code already created
func LeagueCodeIsTaken(code string) bool {
	var ID int = 0

	row := DB.QueryRow(SQL_GET_LEAGUE_ID, code) // If league code exists, its written to ID, else it remains 0
	row.Scan(&ID)

	return ID != 0
}

// Generate random character or number
func randomChar(n int) string {
	var possiblecharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	return string(possiblecharacters[n])
}

// Join a league with given leagueInfo.leagueCode
func JoinLeague(league constants.LeagueInfo) (constants.Errorcode, constants.LeagueInfo) {
	row := DB.QueryRow(SQL_GET_LEAGUE_ID, league.LeagueCode) // Fetch id of given league
	row.Scan(&league.LeagueID)

	if row.Err() != nil { // League doesn't exist or DB failure
		log.Println("ERROR:", row.Err().Error())
		return constants.InternalFail, league
	}
	// Add user to the league
	err := DB.QueryRow(SQL_JOIN_LEAGUE, league.LeagueID, league.Players[0].UserId).Err()

	if err != nil {
		log.Println("ERROR:", err.Error())
		return constants.InternalFail, league
	}

	league, errorcode := GetLeague(league.LeagueID) // Fetch all information from the league, and return it

	return errorcode, league
}

// Remove given user from the specified league
func LeaveLeague(userID int, leagueID int) constants.Errorcode {
	err := DB.QueryRow(SQL_LEAVE_LEAGUE, leagueID, userID).Err() // Remove user from league

	if err != nil {
		log.Println("ERROR:", err.Error())
		return constants.InternalFail
	}

	league, _ := GetLeague(leagueID)
	if len(league.Players) == 0 { // If user was last user, remove league from DB
		deleteLeague(leagueID)
	}

	return constants.No_error
}

// Delete league
func deleteLeague(leagueID int) constants.Errorcode {
	err := DB.QueryRow(SQL_DELETE_LEAGUE, leagueID).Err() // Deletes league from DB

	if err != nil { // If league not empty or other DB error
		println("ERROR: ", err.Error())
		return constants.InternalFail
	}

	return constants.No_error
}

// Fetch all the leagues for the given user
func GetAllLeagues(userID int) ([]constants.LeagueInfo, constants.Errorcode) {
	var allLeages []constants.LeagueInfo
	tx, err := DB.Begin()
	if err != nil {
		log.Printf("ERROR, Failed to begin transaction, error: %v", err.Error())
	}
	// Fetch all the leagues the user is a part of
	rows, err := tx.Query(SQL_GET_USERS_LEAGUES, userID)
	if err != nil { // Rollback if error occurs
		log.Println(err.Error())
		tx.Rollback()
		return allLeages, constants.NoDataSet
	}

	var leagueID int
	var leagueBuf constants.LeagueInfo
	var errorcode constants.Errorcode = constants.No_error

	// Iterate through each row and fetch the leagueID
	for rows.Next() {
		if err := rows.Scan(&leagueID); err != nil {
			log.Fatal(err.Error())
		}
		leagueBuf, errorcode = GetLeague(leagueID) // Use leagueID to fetch leagues' info
		if errorcode != constants.No_error {
			println("ERROR in getAllLeagues: " + internal.PrintError(errorcode))
			return allLeages, errorcode
		}
		allLeages = append(allLeages, leagueBuf) // Add each league to a slice
	}
	rows.Close()
	err = tx.Commit()
	if err != nil {
		log.Printf("ERROR, unable to commit GetAllLeagues transactions, error: %v", err.Error())
	}

	return allLeages, errorcode // Return slice of leagues
}

// Get a given leagues information based on leagueID
func GetLeague(id int) (constants.LeagueInfo, constants.Errorcode) {
	var league constants.LeagueInfo
	league.LeagueID = id

	row := DB.QueryRow(SQL_GET_LEAGUE_NAME_AND_CODE, id) // Fetch League information
	if row.Err() != nil {
		log.Printf("ERROR(dbaxs.go:329): %v", row.Err().Error())
		return league, constants.NoDataSet
	}

	row.Scan(&league.LeagueName, &league.LeagueCode) // Scan rows into league struct

	rows, err := DB.Query(SQL_GET_LEAGUE_PARTICIPANTS, league.LeagueID) // Fetch all the participants from the given league

	if err != nil {
		log.Printf("ERROR(dbaxs.go:338): %v", row.Err().Error())
		rows.Close()
		return league, constants.NoDataSet
	}
	var playerbuf constants.LeaguePlayer
	// Iterate through users from league and add to list
	for rows.Next() {
		if err := rows.Scan(&playerbuf.UserId, &playerbuf.UserName, &playerbuf.Score); err != nil {
			return league, constants.InternalFail
		}
		league.Players = append(league.Players, playerbuf)
	}
	rows.Close() // Close DB connection

	league.Players = QuickSort(league.Players) // Sort players based on score

	return league, constants.No_error
}

/*
 *	Sorting a league's players based on their overall score using the quicksort algorithm,
 *	highest first, lowest last
 */
func QuickSort(list []constants.LeaguePlayer) []constants.LeaguePlayer {
	length := len(list)

	if length < 1 {
		return list
	}

	pivot := list[length-1]
	list = list[:length-1]

	var higher []constants.LeaguePlayer
	var lower []constants.LeaguePlayer

	for _, p := range list {
		if p.Score > pivot.Score {
			higher = append(higher, p)
		} else {
			lower = append(lower, p)
		}
	}

	// Runs recursive, appends the sorted list of lower items to the pivot in the middle and then the sorted higher items
	return append(append(QuickSort(higher), pivot), QuickSort(lower)...)
}

// Join overall league (league every user is part of)
func JoinOverall(user constants.UserInfo) {
	err := DB.QueryRow(SQL_JOIN_LEAGUE, 1, user.UserID).Err() // Insert user into overall league

	if err != nil {
		log.Println("ERROR:", err.Error())
		return
	}
}

const SQL_CLEAR_LEAGUES = "DELETE FROM liga WHERE bruker_id = ?;"

// Clear a given user from ALL leagues
func clearLeagues(userID int) (errorCode constants.Errorcode) {
	err := DB.QueryRow(SQL_CLEAR_LEAGUES, userID).Err() // Removes userID from leagues table
	if err != nil {
		return constants.InternalFail
	}
	return constants.No_error
}

package dbaxs

import (
	"Integrasjonsprosjekt/constants"
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const SQL_GET_USERNAME_BY_ID = `SELECT bruker_id FROM bruker WHERE brukernavn = ?;`
const SQL_GET_USER_BY_ID = `SELECT brukernavn FROM bruker WHERE bruker_id = ?;`
const SQL_ADD_USER = "INSERT INTO bruker (brukernavn, passord, epost, poeng) VALUES (?, ?, ?, 0);"
const SQL_REMOVE_USER = `DELETE FROM bruker WHERE brukernavn = ?`
const SQL_GET_HASHED_PASSWORD = `SELECT passord, bruker_id FROM bruker WHERE brukernavn = ?;`
const SQL_GET_USER_ID = "SELECT bruker_id FROM bruker WHERE brukernavn = ?;"

/*
 * getUserID fetches user ID from username
 * Returns ID and errorcode
 * Returns 0 and errorCode (UsernameDoesntExist OR InternalFail)
 */
func getUserID(username string) (ID int, errorCode constants.Errorcode) {
	err := DB.QueryRow(SQL_GET_USERNAME_BY_ID, username).Scan(&ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ID, constants.UsernameDoesntExist
		}
		return ID, constants.InternalFail
	}
	return ID, constants.No_error
}

/*
 * Returns true if user id exists
 */
func checkUserID(ID int) bool {
	var str string
	row := DB.QueryRow(SQL_GET_USER_BY_ID, ID)

	row.Scan(&str)
	return (len(str) > 0) // Found existing user or email
}

/*
 * Add a user, returns an enum: constants.No_error (0) if successful
 */
func addUserSQL(user constants.UserInfo) (constants.UserInfo, constants.Errorcode) {
	tx, err := DB.Begin()
	if err != nil {
		log.Printf("ERROR, Failed to begin transaction, error: %v", err.Error())
		return user, constants.InternalFail
	}
	// Insert user into DB
	_, err = tx.Exec(SQL_ADD_USER, user.Username, user.Password, user.Email)
	if err != nil { // Rollback if error occurs
		if strings.Contains(err.Error(), "unik_epost") {
			tx.Rollback()
			return user, constants.EmailInUse
		} else if strings.Contains(err.Error(), "unik_brukernavn") {
			tx.Rollback()
			return user, constants.UsernameInUse
		}
		log.Println("ERROR:", err.Error())
		tx.Rollback()
		return user, constants.InternalFail
	}
	// Fetch user id using username
	err = tx.QueryRow(SQL_GET_USER_ID, user.Username).Scan(&user.UserID)
	if err != nil {
		log.Println("Failed to get user id, error:", err.Error())
		return user, constants.InternalFail
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("ERROR, unable to commit addUserSQL transaction, error: %v", err.Error())
		return user, constants.InternalFail
	}

	return user, constants.No_error // Return user struct when successfull
}

/*
 * SQL query to delete given username
 * Removes user predictions for matches, medal and topscorer
 */
func removeUserSQL(username string) error {
	// Remove user predictions before removing user
	userID, errorCode := getUserID(username)
	if errorCode != constants.No_error {
		return fmt.Errorf("ERROR: Failed to fetch user id, error code: %v", errorCode)
	}
	err := clearMatchPredictions(userID) // Remove users match predictions
	if err != nil {
		return fmt.Errorf("ERROR: Failed to clear users match predictions, error: %v", err.Error())
	}
	err = clearMedalPredictions(userID) // Remove users medal predictions
	if err != nil {
		return fmt.Errorf("ERROR: Failed to clear user medal predictions, error: %v", err.Error())
	}
	err = clearTopScorerPredictions(userID) // Remove users topscorer predictions
	if err != nil {
		return fmt.Errorf("ERROR: Failed to clear users topscorer predictions, error: %v", err.Error())
	}

	errorCode = clearLeagues(userID) // Remove a user from all leagues
	if errorCode != constants.No_error {
		return fmt.Errorf("ERROR: Failed to clear leagues, error: %v", err.Error())
	}

	err = DB.QueryRow(SQL_REMOVE_USER, username).Err() // Finally remove user after all DB dependencies are removed
	if err != nil {
		return fmt.Errorf("ERROR: Failed to remove user from database: %v", err.Error())
	}
	return nil
}

/*
 * Fetches a hashed password from database, returns empty string if it fails
 */
func fetchHashedPassword(username string) (hashedPassword string, userID int) {
	// Fetch hashed password from DB
	row := DB.QueryRow(SQL_GET_HASHED_PASSWORD, username).Scan(&hashedPassword, &userID)
	if row != nil {
		log.Println(row.Error())
	}
	return hashedPassword, userID
}

/*
 * Adds a user to database if username and password is valid
 */
func AddUser(user constants.UserInfo) constants.Errorcode {
	var errorCode constants.Errorcode
	var err error
	user = userToLowercase(user)
	// Check that username and email is valid
	if errorCode = usernameAndEmailValid(user.Username, user.Email); errorCode != constants.No_error {
		return errorCode
	}
	// Check that password is valid
	if errorCode = passwordValid(user.Password); errorCode != constants.No_error {
		return constants.InvalidPassword
	}
	user.Password, err = hashAndSalt(user.Password) // Hash and salt password
	if err != nil {
		log.Println("ERROR(dbaxs.go:62): Failed to hash and salt password")
		return constants.InternalFail
	}

	user, errorCode = addUserSQL(user) // Add user with hashed password to DB
	if errorCode != constants.No_error {
		return errorCode
	}
	// Every new user joins the overall league
	JoinOverall(user)

	return errorCode

}

// Remove user a user from DB
func RemoveUser(user constants.UserInfo) (errorCode constants.Errorcode) {
	// Remove user from DB with username
	err := removeUserSQL(user.Username)
	if err != nil {
		log.Println("Error removing user:", err.Error())
		return constants.InternalFail
	}
	return constants.No_error
}

// Validate login, checks that DB's hashed password matches user's unhashed password and returns userID
func ValidateLogin(user constants.UserInfo) int {
	hashedPassword, userID := fetchHashedPassword(user.Username) // Fetch hashed Password
	// Check that user password matches hashed DB password
	if checkPassword(hashedPassword, user.Password) { //
		return userID
	}
	return -1
}

// Usernname must be either letter, number or underscore (connot start with underscore)
var isLetterOrUnderscore = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z_0-9]+[^_]$`)

/*
 * Check if username is compliant
 * - Must be five characters
 * - Must be unique
 * - Only contain characters or underscore (abc_de)
 * - Must start with a letter
 */
func usernameAndEmailValid(username string, email string) constants.Errorcode {
	if len(username) < 5 {
		return constants.InvalidUsername
	}
	if !isLetterOrUnderscore.Match([]byte(username)) {
		return constants.InvalidUsername
	}
	return constants.No_error
}

/*
 * Checks if the hashed and unhashed passwords match
 */
func checkPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // If password is correct compare results in nil
}

/*
 * Convert email and username to lowercase to ensure uniqueness and storage
 */
func userToLowercase(user constants.UserInfo) constants.UserInfo {
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)
	return user
}

/*
 * Check if password is compliant
 * - Must be atleast five characters
 */
func passwordValid(password string) constants.Errorcode {
	if len(password) < 5 {
		return constants.PasswordTooShort
	}
	// Any more checks needed?
	return constants.No_error
}

/*
 * Creates hashed password
 */
func hashAndSalt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		log.Println("(hashAndSalt) produced error:", err.Error())
		return "", err
	}
	return string(hashedPassword), nil
}

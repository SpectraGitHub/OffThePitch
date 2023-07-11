package dbaxs

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"testing"
)

// Test addUserSQL and removeUserSQL
func TestAddRemoveUserSQL(t *testing.T) {
	setTestServerInfo()
	ConnDB(server.username, server.password, server.ip, server.port, server.db_name)

	testName := "testExample9876"
	user := constants.UserInfo{Username: testName, Email: testName, Password: testName}

	// Try to add example user
	user, errorCode := addUserSQL(user)
	if errorCode != constants.No_error {
		t.Errorf("Failed to add user, errorCode: %v", internal.PrintError(errorCode))
	}

	// Try to overwrite user again
	user, errorCode = addUserSQL(user)
	if errorCode == constants.No_error { // If successfully overwritten user, test fails
		t.Errorf("ERROR, added duplicate user, should have failed")
	}
	err := removeUserSQL(user.Username) // Remove user from DB
	if err != nil {
		t.Errorf("Failed to remove given username: %s", user.Username)
	}
	DB.Close()
}

// Test fetchHashedPassword and checkUserID function
func TestFetchPasswordAndCheckUserID(t *testing.T) {
	setTestServerInfo()
	ConnDB(server.username, server.password, server.ip, server.port, server.db_name)

	// Initiate a test user
	_, testUserInfo, err := initTestUser()
	if err != nil {
		t.Errorf("Failed to INIT test user, error: %v", err.Error())
	}
	// Remover test user from DB
	err = removeUserSQL(testUserInfo.Username)
	if err != nil {
		t.Errorf(err.Error())
	}
	// Add test user to DB
	testUserInfo, errorCode := addUserSQL(testUserInfo)
	if errorCode != constants.No_error {
		t.Errorf("Failed to add user, errorcode: %v", internal.PrintError(errorCode))
	}

	// Fetch test user's hashed password
	hashedPassword, userID := fetchHashedPassword(testUserInfo.Username)
	if !checkUserID(userID) { // Check if user exists
		t.Errorf("Failed to retrieve user id: %v from database", userID)
	}
	// Check if hashedPassword generated matches testUser's password
	if hashedPassword == "" || hashedPassword != testUserInfo.Password {
		t.Errorf("Failed to hash password")
	}

	// Test fetching a password that doesnt exist
	err = removeUserSQL(testUserInfo.Username)
	if err != nil {
		t.Errorf(err.Error())
	}
	// Fetch password from user which doesn't exist
	hashedPassword, userID = fetchHashedPassword(testUserInfo.Username)
	if hashedPassword != "" { // If succesfully fetch hashPassword, test failed
		t.Errorf("Error fetching non-existing user password. Password = %v, should be ''", hashedPassword)
	}
	if userID != 0 {
		t.Errorf("Error, fetching non-existing user id. userID = %v, should be 0", userID)
	}
}

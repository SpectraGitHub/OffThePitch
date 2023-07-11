package dbaxs

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"fmt"
	"log"
	"math/rand"
	"testing"
)

/*
 * Create test user and fetch userID
 */
func initTestUser() (userID int, testUser constants.UserInfo, err error) {
	testUser = constants.UserInfo{Username: "TESTUSER", Email: "mndasdasda@asddas.dsad", Password: "abcddef"}
	errorCode := AddUser(testUser)
	if errorCode != constants.No_error && errorCode != constants.UsernameInUse {
		log.Println("ERROR ADDING USER, error code:", errorCode)
	}

	userID, errorCode = getUserID(testUser.Username)
	if errorCode != constants.No_error {
		return 0, testUser, fmt.Errorf("Failed to fetch userID from username: %v", testUser.Username)
	}
	return userID, testUser, nil
}

/**
 * Test savePrediction and GetPredictions functions
 */
func TestSaveAndGetPredictionsFromDB(t *testing.T) {
	// SET TEST SERVER INFO AND INITIATE TEST USER
	setTestServerInfo()
	ConnDB(server.username, server.password, server.ip, server.port, server.db_name)
	testUserID, _, err := initTestUser()
	if err != nil {
		t.Errorf("Failed to init TestUser")
	}

	// Initiate ten random guesses
	var preds []constants.Prediction
	testPredictionsData := constants.Predictions{UserID: testUserID}
	for i := 0; i < 10; i++ {
		preds = append(preds, constants.Prediction{MatchID: i, HomeScore: rand.Intn(10), AwayScore: rand.Intn(10), Points: 0})
	}
	testPredictionsData.Predictions = preds

	tx, err := DB.Begin()
	if err != nil {
		t.Errorf("Failed to begin transaction, error: %v", err.Error())
	}
	for _, pred := range testPredictionsData.Predictions {
		savePrediction(tx, testPredictionsData.UserID, pred)
	}

	// Fetch guesses from the test user
	predictionResult := GetPredictions(testPredictionsData.UserID)
	if testPredictionsData.UserID != predictionResult.UserID {
		t.Errorf("ERROR, fetched userID %v doesn't match test user %v", predictionResult.UserID, testPredictionsData.UserID)
	}
	if len(predictionResult.Predictions) != len(testPredictionsData.Predictions) {
		t.Errorf("ERROR, fetched %v predictions from database, should have been %v", predictionResult.Predictions, testPredictionsData.Predictions)
	}
	for i := range predictionResult.Predictions {
		if testPredictionsData.Predictions[i] != predictionResult.Predictions[i] {
			t.Errorf("ERROR, prediction: %v, doesn't match test data: %v", predictionResult.Predictions[i], testPredictionsData.Predictions[i])
		}
	}

}

/*
 * Tests the functions: saveTopScorers and clearTopScorePredictions
 * saveTopScorers: 1. test data should  succeed
 * 				   2. test data should fail to overwrite
 * clearTopScorers: 1. clears data.
 *                  2. Now Test that saveTopScorers works again
 */
func TestSaveAndClearTopScorers(t *testing.T) {
	// SET TEST SERVER INFO AND INITIATE TEST USER
	setTestServerInfo()
	ConnDB(server.username, server.password, server.ip, server.port, server.db_name)
	userID, testUser, err := initTestUser()
	if err != nil {
		t.Errorf("Failed to init TestUser")
	}

	resultUserID, _ := getUserID(testUser.Username)
	if userID != resultUserID {
		t.Errorf("Recieved wrong user id: %v, should have been: %v", resultUserID, userID)
	}

	var topscorer constants.TopScorer
	var topScorerData constants.TopScorerPredictions
	topScorerData.UserID = userID

	// Initiate top scorer data
	players := [3]int{944234, 944199, 943998}
	for _, i := range players {
		topscorer.Player_id = i
		topScorerData.Topscorers = append(topScorerData.Topscorers, topscorer)
	}

	errorCode := SaveTopScorerPredictions(topScorerData) // Save test data
	if errorCode != constants.No_error {
		t.Errorf("Failed to save topScorers, errorcode: %v", internal.PrintError(errorCode))
	}

	topScorerResults, errorCode := GetTopscorers(userID)
	if errorCode != constants.No_error {
		t.Errorf("Failed to fetch topscorer results, ERROR: %v", internal.PrintError(errorCode))
	}
	match := checkSlicesMatch(topScorerData, topScorerResults)
	if !match {
		t.Errorf("Topscorers doesn't match. Test data: %v, Result data: %v", topScorerData, topScorerResults)
	}

	// TRY TO OVERWRITE OLD DATA
	players = [3]int{36756, 40350, 438567}
	topScorerData.Topscorers = []constants.TopScorer{}
	for _, i := range players {
		topscorer.Player_id = i
		topScorerData.Topscorers = append(topScorerData.Topscorers, topscorer)
	}
	errorCode = SaveTopScorerPredictions(topScorerData) // Try to save already existing row should succeed
	if errorCode != constants.No_error {
		t.Errorf("Test failed, failed to overwrite topscorer data on userID: %v", topScorerData.UserID)
	}
	// COMPARE NEW DATA
	topScorerResults, errorCode = GetTopscorers(topScorerData.UserID)
	if errorCode != constants.No_error {
		t.Errorf("Failed to fetch topscorer results, ERROR: %v", internal.PrintError(errorCode))
	}

	match = checkSlicesMatch(topScorerData, topScorerResults)
	if !match {
		t.Errorf("Topscorers doesn't match. Test data: %v, Result data: %v", topScorerData, topScorerResults)
	}

	// Finally remove test data
	errorCode = RemoveUser(testUser)
	if errorCode != constants.No_error {
		t.Errorf("Failed to remove username: %v, error: %v", testUser.Username, internal.PrintError(constants.No_error))
	}
}

/*
 * Compares content of two slices, if match return true
 */
func checkSlicesMatch(t1 constants.TopScorerPredictions, t2 constants.TopScorerPredictions) bool {

	if len(t1.Topscorers) != len(t2.Topscorers) {
		return false
	}
	count := 0
	for i := range t1.Topscorers {
		for j := range t2.Topscorers {
			if t1.Topscorers[i].Player_id == t2.Topscorers[j].Player_id {
				count++
			}
		}
	}
	return count == len(t1.Topscorers)
}

/**
 * Test saveMedalPredictions and getMedalPredictions
 */
func TestMedalPredictions(t *testing.T) {
	// SET TEST SERVER INFO AND INITIATE TEST USER
	setTestServerInfo()
	ConnDB(server.username, server.password, server.ip, server.port, server.db_name)
	testUserID, testUser, err := initTestUser() // Create test user
	if err != nil {
		t.Errorf("Failed to init TestUser")
	}

	// INSERT PREDICTIONS FOR TEST USER
	testMedalPrediction := constants.MedalPredictions{UserID: testUserID, Gold: 550, Silver: 927, Bronze: 511}
	errorCode := SaveMedalPredictions(testMedalPrediction)
	if errorCode != constants.No_error {
		t.Errorf("Failed to save medal predictions: %v", internal.PrintError(errorCode))
	}

	// FETCH PREDICTIONS FOR TEST USER
	resultMedalPrediction, errorCode := getMedalPredictions(testUserID)
	if errorCode != constants.No_error {
		t.Errorf("Failed to fetch medal predictions: %v", internal.PrintError(errorCode))
	}
	if testMedalPrediction != resultMedalPrediction { // Check that fetched and sent data matches
		t.Errorf("Test data: %v, doesn't match result: %v", testMedalPrediction, resultMedalPrediction)
	}

	// Try to overwrite data
	testMedalPrediction = constants.MedalPredictions{UserID: testMedalPrediction.UserID, Gold: 514, Silver: 559, Bronze: 556}
	errorCode = SaveMedalPredictions(testMedalPrediction)
	if errorCode != constants.No_error {
		t.Errorf("Failed to overwrite medal predictions: %v", internal.PrintError(errorCode))
	}

	// FETCH PREDICTIONS FOR TEST USER
	resultMedalPrediction, errorCode = getMedalPredictions(testUserID)
	if errorCode != constants.No_error {
		t.Errorf("Failed to fetch overwritten medal predictions: %v", internal.PrintError(errorCode))
	}
	if testMedalPrediction != resultMedalPrediction { // Check that fetched and sent data matches
		t.Errorf("Test data: %v, doesn't match result: %v", testMedalPrediction, resultMedalPrediction)
	}
	// Finally clear data
	clearMedalPredictions(testUserID)
	removeUserSQL(testUser.Username)
}

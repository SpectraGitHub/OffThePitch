package dbaxs

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"testing"
)

/*
 * Server configurations
 */
type serverInfo struct {
	username string
	password string
	ip       string
	port     string
	db_name  string
}

var server serverInfo

/*
 * Set test server information here
 */
func setTestServerInfo() {
	server = serverInfo{username: ADM_USRNM, password: ADM_PSWD, ip: SRV_IP, port: SRV_PORT, db_name: TEST_DB_NAME}
}

/*
 * Connect to test database which is identical to the main database
 */
func TestConnDB(t *testing.T) {
	setTestServerInfo()
	connected := ConnDB(server.username, server.password, server.ip, server.port, server.db_name)
	if !connected {
		t.Errorf("Failed to connect to database: %v", server.db_name)
	}
}

/*
 * Test ValidateLogin and RemoveUser function
 */
func TestAddValidateRemoveUser(t *testing.T) {
	setTestServerInfo()
	ConnDB(server.username, server.password, server.ip, server.port, server.db_name)
	userID, testUser, err := initTestUser()

	testUser.UserID = userID
	if err != nil {
		t.Errorf("Failed to init testuser, error: %v", err.Error())
	}

	userID = ValidateLogin(testUser)
	if userID == -1 {
		t.Errorf("Failed to validate user: %v", testUser.Username)
	}
	errorCode := RemoveUser(testUser)
	if errorCode != constants.No_error {
		t.Errorf("Failed to remove user, error: %v", internal.PrintError(errorCode))
	}
}

/*
func TestGetUserTopscorers(t *testing.T) {
	setTestServerInfo()
	ConnDB(server.username, server.password, server.ip, server.port, server.db_name)

	// TEST INSERT TOPSCORERS

	topscorers, errorCode := GetTopscorers(176)
	if errorCode != constants.No_error {
		t.Errorf("Failed to get user test scorers")
	}
	log.Println(topscorers)
}
*/

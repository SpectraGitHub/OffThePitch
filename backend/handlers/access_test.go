package handlers

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"net/http"
	"testing"
)

// Unit test for the GenerateJWT and VerifyJWT functions
func TestGenerateAndVerifyJWT(t *testing.T) {
	// The user ID to test
	userIdTest := 144
	// Generate JWT using the set ID
	stringToken, err := generateJWT(userIdTest)
	if err != nil {
		t.Errorf("Failed to generate JWT, error: %v", err.Error())
	}

	// Create request and add token
	req, err := http.NewRequest("GET", "http://test.com", nil)
	if err != nil {
		t.Errorf("ERROR: Failed to create request")
	}
	req.Header.Add("Token", stringToken)

	// Testing the verifyJWT function
	userIdResult, errorCode := verifyJWT(req)
	if errorCode != constants.No_error {
		t.Errorf("Failed to verify JWT, error: %v", internal.PrintError(errorCode))
	}

	// Checking if the result is as expected
	if userIdTest != userIdResult {
		t.Errorf("Verification resulted in wrong user id, got: %v, expected: %v", userIdResult, userIdTest)
	}

}

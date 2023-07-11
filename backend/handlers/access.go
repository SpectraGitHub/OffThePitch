package handlers

/*
*	General API handler functions related to access and security
 */

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

/*
 * Generates a token with userid and an expiration time specified at: constants.JWT_EXPIRATION_MINUTES
 * Returns the token and potential error
 */
func generateJWT(userid int) (stringToken string, err error) {
	expiresAt := time.Now().Add(time.Duration(constants.JWT_EXPIRATION_MINUTES) * time.Minute)
	claims := &constants.JWTClaim{
		Userid: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err = token.SignedString([]byte(constants.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return stringToken, nil
}

// Function used to get userID from a request containing a JWT token in the header
func GetUserId(w http.ResponseWriter, r *http.Request) int {
	id, errorcode := verifyJWT(r)
	if errorcode != constants.No_error {
		w.WriteHeader(401)
		writeResponse(w, (internal.PrintError(errorcode)))
		return 0
	}

	return id
}

// Function to retrieve the User ID from the request containing a JWT token
func verifyJWT(r *http.Request) (userid int, errorCode constants.Errorcode) {
	tokenString := r.Header.Get("Token")
	if tokenString == "" {
		log.Println("Failed to find token in header")
		return 0, constants.Unauthorized
	}
	// Parsing the token from the request header
	token, err := jwt.ParseWithClaims(
		r.Header.Get("Token"),
		&constants.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.JWT_SECRET_KEY), nil
		})
	if err != nil { // Failed to fetch token
		log.Println("ERROR in verifyJWT:", err.Error())
		return 0, constants.Unauthorized
	}
	// Checking if the token is valid
	if !token.Valid {
		log.Println("failed, token is invalid")
		return 0, constants.Unauthorized
	}
	claims, ok := token.Claims.(*constants.JWTClaim)
	if !ok {
		log.Println("failed to fetch claims from token")
		return 0, constants.Unauthorized
	}
	return claims.Userid, constants.No_error
}

// Function to enable CORS.
// Needed in order for axios to be able to send requests to the API
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST")
	(*w).Header().Set("Access-Control-Allow-Headers", "token")
	(*w).Header().Set("Content-Type", "text/html; charset=utf-8")
}

// Function to write a response with a given body of any type
func writeResponse(w http.ResponseWriter, res interface{}) {
	// Marshaling class/response to JSON
	output, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Failed converting response to JSON format: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Printing response
	_, err = fmt.Fprint(w, string(output))
	if err != nil {
		http.Error(w, "Failed writing response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

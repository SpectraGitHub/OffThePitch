package handlers

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/dbaxs"
	"Integrasjonsprosjekt/internal"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Function to handle requests to create a new user
func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != "POST" {
		println("Invalid request. Expected POST, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use POST with this URL.", http.StatusMethodNotAllowed)
		return
	}
	r.Header.Add("content-type", "application/json")

	// Decode the given user info from the body of the request to an empty user object
	var newUser constants.UserInfo
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Fatal(err)
	}

	// Add new user to the database
	errorcode := dbaxs.AddUser(newUser)

	if errorcode == constants.No_error {
		// Write status created and a message saying the player is created
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "New user created")
	} else {
		fmt.Println("Error: " + internal.PrintError(errorcode))
		http.Error(w, "Error: "+internal.PrintError(errorcode), http.StatusNotAcceptable)
	}
}

// Function to handle requests to login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" && r.Method != "OPTIONS" {
		println("Invalid request. Expected POST, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use POST with this URL.", http.StatusMethodNotAllowed)
		return
	}
	enableCors(&w)
	r.Header.Add("content-type", "application/json")

	// Decode the login info from the request body to a new user object
	var user constants.UserInfo
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	// Validate that the login info is correct
	userId := dbaxs.ValidateLogin(user)

	// If the info is correct
	if userId != -1 {
		// Generate a new JWT token
		token, err := generateJWT(userId)
		if err != nil {
			println("Error in generateJWT: ", err.Error())
		}

		// Set status created and write the token back as a response
		w.WriteHeader(202)
		writeResponse(w, token)

		return
	} else {
		w.WriteHeader(401)
	}
}

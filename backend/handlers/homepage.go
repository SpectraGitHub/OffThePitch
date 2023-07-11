package handlers

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/dbaxs"
	"Integrasjonsprosjekt/internal"
	"net/http"
)

// Function to return all info needed to display the homepage correctly
func HomepageInfoHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "GET" && r.Method != "OPTIONS" {
		println("Invalid request. Expected POST, recieved: ", r.Method)
		http.Error(w, "Method not supported. Please use GET with this URL.", http.StatusMethodNotAllowed)
		return
	} else if r.Method == "OPTIONS" {
		return
	}

	var homepageInfo constants.HomePageInfo
	var errorcode constants.Errorcode

	// Get user ID from token in header
	homepageInfo.UserId = GetUserId(w, r)
	if homepageInfo.UserId == 0 {
		return
	}

	// Use userID to get the needed information from the database
	homepageInfo, errorcode = dbaxs.GetHomepageInfo(homepageInfo)

	if errorcode != constants.No_error {
		println(internal.PrintError(errorcode))
	}

	writeResponse(w, homepageInfo)
}

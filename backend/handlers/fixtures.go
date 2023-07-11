package handlers

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
	"net/http"
)

// Handler to return all fixtures with corresponding info
func FixtureHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	fixtures := GetFixtures()
	writeResponse(w, fixtures)
}

// Function to return fixtures if they are stored internally, if not they are updated from the API
func GetFixtures() []constants.Match {
	if len(internal.Fixtures.Matches) == 0 {
		internal.UpdateFixtures(constants.WORLDCUPID)
	}
	return internal.Fixtures.Matches
}

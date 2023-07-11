package main

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/dbaxs"
	"Integrasjonsprosjekt/handlers"
	"Integrasjonsprosjekt/internal"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT not set. Default: 8080")
		port = "8080"
	}

	// Connecting to database
	dbaxs.ConnDB(dbaxs.ADM_USRNM, dbaxs.ADM_PSWD, dbaxs.SRV_IP, dbaxs.SRV_PORT, dbaxs.DB_NAME)

	var tournament = "worldCup"
	internal.LoadSavedFixtures(tournament)
	internal.UpdateFixtures(constants.WORLDCUPID)

	internal.SaveFixtures(tournament)
	dbaxs.GetTeams()

	for i, match := range internal.Fixtures.Matches {
		if match.Status.Updated && !match.Status.PointsCalculated {
			dbaxs.CalculatePoints(match)
			internal.Fixtures.Matches[i].Status.PointsCalculated = true
		}
	}

	http.HandleFunc(constants.DEFAULT+constants.FIXTUREENDPOINT, handlers.FixtureHandler)
	http.HandleFunc(constants.DEFAULT+constants.NEWUSER, handlers.NewUserHandler)
	http.HandleFunc(constants.DEFAULT+constants.LOGIN, handlers.LoginHandler)
	http.HandleFunc(constants.DEFAULT+constants.SAVEPREDICTIONS, handlers.SavePredictionsHandler)
	http.HandleFunc(constants.DEFAULT+constants.GETPREDICTIONS, handlers.GetPredictionsHandler)
	http.HandleFunc(constants.DEFAULT+constants.SAVETOPSCOREPREDICTIONS, handlers.SaveTopScorerPredictionHandler)
	http.HandleFunc(constants.DEFAULT+constants.GETTOPSCOREPREDICTIONS, handlers.GetTopScorerPredictionHandler)
	http.HandleFunc(constants.DEFAULT+constants.GETMEDALPREDICTIONS, handlers.GetMedalPredictionHandler)
	http.HandleFunc(constants.DEFAULT+constants.SAVEMEDALPREDICTIONS, handlers.SaveMedalPredictionHandler)
	http.HandleFunc(constants.DEFAULT+constants.GETTEAMNAMES, handlers.GetTeamNamesHandler)
	http.HandleFunc(constants.DEFAULT+constants.GETALLPLAYERS, handlers.GetPlayersHandler)
	http.HandleFunc(constants.DEFAULT+constants.CREATELEAGUE, handlers.CreateLeagueHandler)
	http.HandleFunc(constants.DEFAULT+constants.JOINLEAGUE, handlers.JoinLeagueHandler)
	http.HandleFunc(constants.DEFAULT+constants.GETHPI, handlers.HomepageInfoHandler)
	http.HandleFunc(constants.DEFAULT+constants.GETLEAGUES, handlers.GetLeaguesHandler)
	http.HandleFunc(constants.DEFAULT+constants.LEAVELEAGUE, handlers.LeaveLeagueHandler)

	go checkForUpdates(time.Minute * time.Duration(constants.UPDATEFREQUENCY))

	go checkConnections(time.Minute * 5)

	log.Println("Running on port: ", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}

// Function to run by a separate goroutine to update the fixtures with scores, as well as all points every x minutes
func checkForUpdates(t time.Duration) {
	for range time.Tick(t) {
		println("Update at", internal.GetTime())
		internal.LoadSavedFixtures("worldCup")

		internal.UpdateFixtures(constants.WORLDCUPID)

		for i, match := range internal.Fixtures.Matches {
			if match.Status.Updated && !match.Status.PointsCalculated {
				dbaxs.CalculatePoints(match)
				internal.Fixtures.Matches[i].Status.PointsCalculated = true
			}
		}

		if len(internal.Fixtures.Matches) > 0 {
			internal.SaveFixtures("worldCup")

		} else {
			println("PROBLEM SAVING FIXTURES, NO MATCHES EXIST IN internal.Fixtures")
		}
	}
}

/**
 * Used to check for DB connections to monitor that DB connections properly close
 */
func checkConnections(t time.Duration) {
	for range time.Tick(t) {
		println("DB Connections:", dbaxs.DB.Stats().OpenConnections)

		println("CLOSED LIFETIME connections:", dbaxs.DB.Stats().MaxLifetimeClosed)

	}
}

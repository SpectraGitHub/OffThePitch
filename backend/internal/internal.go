package internal

import (
	"Integrasjonsprosjekt/constants"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

// Globally saved fixtureinfo
var Fixtures constants.FixturesResults
var RequestsLastMin int
var TeamMap = make(map[int]string)
var FixtureTimeMap = make(map[int]string)

// Function to update fixtures from API and set the correct status for all existing fixtures
func UpdateFixtures(id int) constants.FixturesResults {
	if time.Since(Fixtures.LastUpdated) >= time.Minute*time.Duration(constants.UPDATEFREQUENCY-1) {
		Fixtures = GetAllFixtures(id)
	}
	SetFixtureTimeMap()

	Fixtures = SetMatchdays(Fixtures)
	SetRoundsStarted()
	GetGoalScorers()

	return Fixtures
}

// Function to get all the matches from a given league
func GetAllFixtures(leagueID int) constants.FixturesResults {
	url := constants.API + constants.FIXTURES + fmt.Sprintf("%d", leagueID)
	var allFix constants.FixturesResults

	// Create a new request
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Printf("Error, failed to fetch all fixtures from API, error: %v", err.Error())
	}

	// Set the API key header
	request.Header.Set(constants.KEYNAME, constants.KEYVALUE)

	client := &http.Client{}
	res, err := client.Do(request)
	requestSent()

	if err != nil {
		log.Printf("Error, failed to fetch all fixtures from API, error: %v", err.Error())
	}

	decodedFix := decodeFixtures(res)

	if len(Fixtures.Matches) > 0 {
		allFix = updateFromDecoded(Fixtures, decodedFix)
	} else {
		allFix = decodedFix
	}

	allFix.LastUpdated = time.Now()
	return allFix
}

// Function to decode the fixtures recieved from the external API
func decodeFixtures(res *http.Response) constants.FixturesResults {
	var allFixtures constants.FixturesInfo

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&allFixtures); err != nil {
		log.Printf("Error, failed to decode fixtures from API, error: %v", err.Error())
	}

	return allFixtures.FixturesResults
}

// Function used to update existing fixtures with the new info recieved from the external API
// The other status than "short" is NOT updated
func updateFromDecoded(prev constants.FixturesResults, new constants.FixturesResults) constants.FixturesResults {
	for i := range prev.Matches {
		prev.Matches[i].AwayTeam.Name = new.Matches[i].AwayTeam.Name
		prev.Matches[i].HomeTeam.Name = new.Matches[i].HomeTeam.Name
		prev.Matches[i].AwayTeam.Score = new.Matches[i].AwayTeam.Score
		prev.Matches[i].HomeTeam.Score = new.Matches[i].HomeTeam.Score
		prev.Matches[i].Date = new.Matches[i].Date
		prev.Matches[i].Time = new.Matches[i].Time
		prev.Matches[i].MatchID = new.Matches[i].MatchID
		prev.Matches[i].Round = new.Matches[i].Round
		prev.Matches[i].Status.Short = new.Matches[i].Status.Short
	}
	return prev
}

// Setting a map with ID and gametime for easier sorting
func SetFixtureTimeMap() {
	for _, match := range Fixtures.Matches {
		FixtureTimeMap[match.MatchID] = match.Date + match.Time
	}
}

// Setting what matchday each match belongs to depending on their gametime
func SetMatchdays(allFix constants.FixturesResults) constants.FixturesResults {
	var startTime string
	for x, match := range allFix.Matches {
		startTime = match.Date + match.Time
		for i := range constants.MATCHDAYDIVIDERS {
			if startTime >= constants.MATCHDAYDIVIDERS[i] && startTime < constants.MATCHDAYDIVIDERS[i+1] {
				allFix.Matches[x].Round.Matchday = i
				continue
			}
		}
	}
	return allFix
}

// Sets if the round is started for all matches, based on the current time
func SetRoundsStarted() {
	timenow := GetTime()
	for i, match := range Fixtures.Matches {
		if timenow > constants.MATCHDAYDIVIDERS[match.Round.Matchday] {
			Fixtures.Matches[i].Status.RoundStarted = true
		} else {
			Fixtures.Matches[i].Status.RoundStarted = false
		}
	}
}

// Get full information about all finished matches
func GetGoalScorers() {
	for i := range Fixtures.Matches {
		// Whenever a match is finished and it is not yet updated, find score and goalscorers
		if Fixtures.Matches[i].Status.Short == "FT" && !Fixtures.Matches[i].Status.Updated && RequestsLastMin < constants.MAXREQUESTSPERMIN {
			println("GET SINGLE MATCH INFO FOR: ", Fixtures.Matches[i].HomeTeam.Name, Fixtures.Matches[i].AwayTeam.Name)
			newInfo := GetMatchInfo(Fixtures.Matches[i].MatchID)
			Fixtures.Matches[i] = mergeStructs(Fixtures.Matches[i], newInfo)
		}
	}
}

// Function to get specific info about a single match (Get goalscorers etc)
func GetMatchInfo(matchID int) constants.FinishedMatch {
	url := constants.API + constants.MATCH + fmt.Sprint(matchID)
	// Create a new request
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Printf("Error, failed to fetch matchInfo from API, error: %v", err.Error())
	}

	// Set the API key header
	request.Header.Set(constants.KEYNAME, constants.KEYVALUE)

	client := &http.Client{}
	res, err := client.Do(request)

	if err != nil {
		log.Printf("Error, failed to fetch matchInfo from API, error: %v", err.Error())
	}
	// Add a request to the number of requests sent in the last minute
	requestSent()

	finishedMatchInfo := decodeMatchInfo(res)
	return finishedMatchInfo
}

// Function to decode the specified match info recieved from the external API
func decodeMatchInfo(res *http.Response) constants.FinishedMatch {
	var buf constants.MatchResultBuffer
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&buf); err != nil {
		log.Printf("Error, failed to decode matchInfo from API, error: %v", err.Error())
	}

	return buf.FinishedMatch
}

// Function to load the saved fixtures from the "tournament"Fixtures.json file
func LoadSavedFixtures(tournament string) {
	file, err := ioutil.ReadFile(tournament + "Fixtures.json")
	if err != nil {
		log.Println(err)
	}

	_ = json.Unmarshal([]byte(file), &Fixtures)
}

// Function to save all fixtures to the "tournament"Fixtures.json file
func SaveFixtures(tournament string) {
	sortFixtures()

	file, err := json.MarshalIndent(Fixtures, "", "\t")
	if err != nil {
		log.Printf("Error, failed to saveFixtures, error: %v", err.Error())
	}

	err = ioutil.WriteFile(tournament+"Fixtures.json", file, 0644)
	if err != nil {
		log.Printf("Error, failed to saveFixtures, error: %v", err.Error())
	}
}

// Function to sort all players alphabetically
func SortAllPlayers(allplayers constants.AllPlayers) constants.AllPlayers {

	allplayers = sortTeams(allplayers)

	for _, team := range allplayers.Teams {
		team.Players = sortPlayers(team.Players)
	}

	return allplayers
}

// Function to sort all teams alphabetically
func sortTeams(allteams constants.AllPlayers) constants.AllPlayers {
	sort.SliceStable(allteams.Teams, func(i, j int) bool {
		return allteams.Teams[i].TeamName < allteams.Teams[j].TeamName
	})

	return allteams
}

// Function to sort a slice of players alphabetically
func sortPlayers(players []constants.Player) []constants.Player {
	sort.SliceStable(players, func(i, j int) bool {
		return players[i].FirstName < players[j].FirstName
	})
	return players
}

/*
*	Sorts the global fixtures based on the time they will be played
 */
func sortFixtures() {
	log.Printf("GOALSCORER BEFORE SORT: %v", Fixtures.Matches[11])
	sortedFix := quickSortFixtures(Fixtures.Matches)
	Fixtures.Matches = sortedFix
	log.Printf("GOALSCORER AFTER SORT: %v", Fixtures.Matches[11])

}

// Self made quickSort function to sort fixtures after the time they are played
func quickSortFixtures(list []constants.Match) []constants.Match {
	length := len(list)

	if length < 1 {
		return list
	}

	pivot := list[length-1]
	list = list[:length-1]

	var higher []constants.Match
	var lower []constants.Match

	for _, p := range list {
		if p.Date > pivot.Date {
			higher = append(higher, p)
		} else if p.Date == pivot.Date {
			if p.Time > pivot.Time {
				higher = append(higher, p)
			} else {
				lower = append(lower, p)
			}
		} else {
			lower = append(lower, p)
		}
	}

	// Runs recursive, appends the sorted list of lower items to the pivot in the middle and then the sorted higher items
	return append(append(quickSortFixtures(lower), pivot), quickSortFixtures(higher)...)
}

/*
*	Sorts a set of predictions based on the time they will be played
 */
func SortPredictions(preds constants.Predictions) constants.Predictions {
	sort.Slice(preds.Predictions, func(i, j int) bool {
		return FixtureTimeMap[preds.Predictions[i].MatchID] < FixtureTimeMap[preds.Predictions[j].MatchID]
	},
	)
	return preds
}

// Merges two structs, and sets the status to updated
func mergeStructs(prev constants.Match, new constants.FinishedMatch) constants.Match {
	prev.HomeTeam = new.HomeTeam
	prev.AwayTeam = new.AwayTeam
	prev.Status.Short = new.Status.Short

	prev.Status.Updated = true

	return prev
}

// Adds one to the requests sent last minute
func requestSent() {
	if RequestsLastMin == 0 {
		go resetRequestsIn(time.Minute)
	}
	RequestsLastMin++

	println("Requests sent in the last minute: ", RequestsLastMin)

}

// Resets requests after the given amount of time
func resetRequestsIn(t time.Duration) {
	time.Sleep(t)
	RequestsLastMin = 0
}

// Get the current time.
// Used for getting the current time and easily being able to change it to test the application
func GetTime() string {
	time := time.Now().String()
	//time := "2022-11-25 17:00"
	return time
}

// Function returning an error message based on a costom errorcode
func PrintError(err constants.Errorcode) string {
	switch err {
	case constants.EmailInUse:
		return "Email already in use"
	case constants.UsernameInUse:
		return "Username already in use"
	case constants.InternalFail:
		return "Internal fail, please try again later"
	case constants.InvalidEmail:
		return "Invalid email"
	case constants.InvalidPassword:
		return "Invalid password"
	case constants.PasswordTooShort:
		return "Password too short"
	case constants.InvalidUsername:
		return "Invalid username"
	case constants.No_error:
		return "Success"
	case constants.WrongInput:
		return "Wrong input, please double check JSON body"
	case constants.NoDataSet:
		return "No data of the requested type set for this user"
	case constants.Unauthorized:
		return "Failed to verify JWT token, user unauthorized"
	}
	return "Unknown error"
}

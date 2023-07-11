package dbaxs

import (
	"Integrasjonsprosjekt/constants"
	"Integrasjonsprosjekt/internal"
)

/*
 * Fetches the information required in the homepage of our application:
 * Leagues the user is in
 * The trophy winners selected
 * Topscorers the user has selected
 * The next four upcoming matches
 */
func GetHomepageInfo(HPI constants.HomePageInfo) (constants.HomePageInfo, constants.Errorcode) {
	var errorcode constants.Errorcode
	HPI.Leagues, errorcode = GetAllLeagues(HPI.UserId)
	if errorcode != constants.No_error {
		return HPI, errorcode
	}

	HPI.MedalWinners, errorcode = getMedalPredictions(HPI.UserId)
	if errorcode != constants.No_error {
		return HPI, errorcode
	}
	HPI.Predictions = GetPredictions(HPI.UserId)
	HPI.Topscorers, errorcode = GetTopscorers(HPI.UserId)
	HPI.Matches = internal.Fixtures
	HPI.Matches = getNextFourMatches(HPI.Matches)

	return HPI, errorcode
}

/*
 * Fetches the four next matches
 */
func getNextFourMatches(fixtures constants.FixturesResults) constants.FixturesResults {
	now := internal.GetTime()
	for i, f := range fixtures.Matches {
		fixtureTime := f.Date + " " + f.Time
		if now < fixtureTime {
			if len(fixtures.Matches) >= i+4 {
				fixtures.Matches = fixtures.Matches[i : i+4]
				return fixtures
			} else { // Fetch old matches if less than 4 matches left
				fixtures.Matches = fixtures.Matches[:1+4]
				return fixtures
			}

		}
	}
	return fixtures
}

package constants

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type FixturesInfo struct {
	FixturesResults FixturesResults `json:"fixtures-results"`
}

type FixturesResults struct {
	Matches     []Match `json:"matches"`
	LastUpdated time.Time
}

type Match struct {
	MatchID  int      `json:"id"`
	Date     string   `json:"date"`
	Time     string   `json:"time"`
	HomeTeam TeamInfo `json:"home-team"`
	AwayTeam TeamInfo `json:"away-team"`

	Round struct {
		Matchday int    `json:"matchday"`
		Group    string `json:"name"`
	} `json:"round"`

	Status struct {
		Short            string `json:"short"`
		Updated          bool   `json:"updated"`           // Updated after fulltime
		PointsCalculated bool   `json:"points_calculated"` // Points calculated in game
		RoundStarted     bool   `json:"round_started"`     // Has any match in the round started?
	} `json:"status"`
}

type TeamInfo struct {
	Name   string       `json:"name"`
	Score  int          `json:"score"`
	TeamID int          `json:"id"`
	Goals  []GoalInfo   `json:"goals"`
	//Lineup []PlayerInfo `json:"line-up"`		Used for adding players to database
}

type PlayerInfo struct {
	Player Player `json:"player"`
}

type AllPlayers struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	TeamID   int      `json:"team_id"`
	TeamName string   `json:"team_name"`
	Players  []Player `json:"players"`
}

type Player struct {
	ID        int    `json:"id"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	TeamID    int    `json:"team_id"`
}

type GoalInfo struct {
	Minute int `json:"minute"`
	Player struct {
		PlayerID  int    `json:"id"`
		FirstName string `json:"first-name"`
		LastName  string `json:"last-name"`
	} `json:"player"`
}

type MatchResultBuffer struct {
	FinishedMatch FinishedMatch `json:"match"`
}

type FinishedMatch struct {
	MatchID  int      `json:"id"`
	HomeTeam TeamInfo `json:"home-team"`
	AwayTeam TeamInfo `json:"away-team"`
	Status   struct {
		Short string `json:"short"`
	} `json:"status"`
}

type UserInfo struct {
	Username string `json:"username"`
	UserID   int    `json:"userid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Predictions struct {
	UserID      int          `json:"userid"`
	Predictions []Prediction `json:"predictions"`
}

type Prediction struct {
	MatchID   int `json:"matchid"`
	HomeScore int `json:"homescore"`
	AwayScore int `json:"awayscore"`
	Points    int `json:"points"`
}

type TopScorerPredictions struct {
	UserID     int         `json:"userid"`
	Topscorers []TopScorer `json:"topscorers"`
}

type TopScorer struct {
	Player_id   int    `json:"player_id"`
	Player_name string `json:"player_name"`
	Team_id     int    `json:"team_id"`
	Team_name   string `json:"team_name"`
	Goals       int    `json:"goals"`
}

type MedalPredictions struct {
	UserID int `json:"userid"`
	// The team IDs of the top 3 predictions
	Gold   int `json:"gold"`
	Silver int `json:"silver"`
	Bronze int `json:"bronze"`
}

type LeagueInfo struct {
	LeagueID   int            `json:"league_id"`
	LeagueCode string         `json:"league_code"`
	LeagueName string         `json:"league_name"`
	Players    []LeaguePlayer `json:"players"`
}

type LeaguePlayer struct {
	UserId   int    `json:"userid"`
	UserName string `json:"user_name"`
	Score    int    `json:"score"`
}

type HomePageInfo struct {
	UserId       int                  `json:"userid"`
	Matches      FixturesResults      `json:"matches"`
	Predictions  Predictions          `json:"predictions"`
	Topscorers   TopScorerPredictions `json:"topscorers"`
	MedalWinners MedalPredictions     `json:"medals"`
	Leagues      []LeagueInfo         `json:"leagues"`
}

type JWTClaim struct {
	Userid int `json:"userid"`
	jwt.StandardClaims
}

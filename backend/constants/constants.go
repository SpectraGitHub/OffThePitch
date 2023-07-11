package constants

const API string = "https://football-web-pages1.p.rapidapi.com/"
const FIXTURES string = "fixtures-results.json?comp="
const MATCH string = "match.json?match="

const WORLDCUPID int = 88

// API CONSTANTS
const KEYNAME string = "X-RapidAPI-Key"
const KEYVALUE string = "f2742af98emshbeadfc3aaca76ccp149593jsn8765c0a3cb3e"
const UPDATEFREQUENCY int = 5    // How often the matches should be updated from the API (In minutes)
const MAXREQUESTSPERMIN int = 24 // Max number of allowed requests per minute

// POINT CONSTNATS
const INCORRECTPOINTS int = 5
const CORRECTTROPHYPOINTS int = 12

const POINTS1X2 int = 2
const GOALDIFFPOINTS int = 2
const HOMEANDAWAYGOALPOINTS int = 1

const TOPSCOREROFTOURNAMENTPOINTS int = 10
const TOPSCORERGOALPOINTS int = 2

const MORETHAN7GOALS int = 4
const MORETHAN5GOALS int = 2
const MORETHAN3GOALS int = 1

// JWT TOKENS
const JWT_SECRET_KEY string = "something"
const JWT_EXPIRATION_MINUTES int = 60

// Local endpoints
const DEFAULT = "/otp/"
const FIXTUREENDPOINT = "fixtures"
const NEWUSER = "newuser"
const LOGIN = "login"
const GETPREDICTIONS = "getpredictions"
const SAVEPREDICTIONS = "savepredictions"
const GETTOPSCOREPREDICTIONS = "gettopscorers"
const SAVETOPSCOREPREDICTIONS = "savetopscorers"
const GETMEDALPREDICTIONS = "getmedals"
const SAVEMEDALPREDICTIONS = "savemedals"
const GETTEAMNAMES = "getteams"
const GETALLPLAYERS = "getallplayers"
const CREATELEAGUE = "createleague"
const JOINLEAGUE = "joinleague"
const GETLEAGUES = "getleagues"
const LEAVELEAGUE = "leaveleague"
const GETHPI = "gethomepageinfo"

var MATCHDAYDIVIDERS = []string{"2022-01-01 00:00", "2022-11-20 16:00", "2022-11-25 10:00", "2022-11-29 15:00", "2022-12-03 15:00", "2022-12-09 15:00", "2022-12-13 19:00", "2022-12-17 15:00", "2023-01-01 00:00"}

var MATCHFROMEVERYTEAM = []int{407300, 407301, 407429, 407342, 407343, 407298, 407294, 407295, 407296, 407297, 422632, 425526, 425715, 422347, 421815, 421991, 409441, 424167, 423300, 424760, 425628, 411140, 423301}
var MATCHESWITHWCLINEUPS = []int{426237}

// Error code Enums
type Errorcode = int32

const (
	No_error Errorcode = iota
	EmailInUse
	UsernameInUse
	UsernameDoesntExist
	InternalFail // General something went wrong in the system error. Should log the error before returning this
	InvalidEmail
	InvalidPassword
	InvalidUsername
	Unauthorized
	PasswordTooShort
	WrongInput
	NoDataSet
	PrimaryKeyExists
)

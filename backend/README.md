# Off The Pitch - Tournament predictor

Directory containing all files connected to the backend API for OffThePitch

Made by Benjamin Loxley Wood, Sander Thorsen, Martin Rui Andorsen and Stian Tusvik

The backend handles incoming API requests to give information about the game and users, and sends requests to an external API to update the information about the real life tournament to use in the game.

The external API used for this is: https://www.https://www.footballwebpages.co.uk/

In this README file we will explain each endpoint and their possible requests, as well as some other aspects of the backend.

## File tree
The files are separated into folders with the same names as the go packages they contain. The main package and main file is in the top folder.

## How to use the application
The application is ran through SkyHigh, and requires you to be on an NTNU connection to work. This can be achieved either by being physically on the NTNU campus, or by having Cisco VPN and connecting to NTNU.

If you want to run the application on your own machine, you have to clone the project using `git clone https://git.gvk.idi.ntnu.no/course/prog2052/prog2052-2022-workspace/group-04-stian_martin_sander_benjamin/prog2053_04.git`, and then navigate to the project folder.
When in the project folder, run `go run .`to start the application. You can then use the endpoints specified in this file, but change the base URL to localhost:8080.

## Dockerizing the app
```
    docker build --tag otp-backend .
    -d means hidden
    docker run -d --restart unless-stopped -p 8080:8080 otp-backend

```

## Information about all endpoints
All endpoints start with the base URL `http://10.212.170.213:8080/` followed by `otp/` and then the desired endpoint.
An example of a full URL would be `http://10.212.170.213:8080/otp/fixtures`

To test each individual endpoint we recommend using Postman.

## Basic information endpoints
Endpoints used to retrieve basic information about teams, players and fixtures

### Getfixtures
Returns all previous and upcoming fixtures for the tournament.
A fixture contains a fixture id, date, time, status and home and away teams.

If a match is finished this will also include the score and goalscorers.

You do not need any authentication to access this endpoint.
##### - Request
```
Method: GET
URL: http://10.212.170.213:8080/otp/fixtures
```

##### - Response
```
{
   [
    {
        "id": 423193,
        "date": "2022-11-20",
        "time": "16:00",
        "home-team": {
            "name": "Qatar",
            "score": 0,
            "id": 965,
            "goals": null
        },
        "away-team": {
            "name": "Ecuador",
            "score": 2,
            "id": 515,
            "goals": [
                {
                    "minute": 16,
                    "player": {
                        "id": 922583,
                        "first-name": "Enner",
                        "last-name": "Valencia"
                    }
                },
                {
                    "minute": 31,
                    "player": {
                        "id": 922583,
                        "first-name": "Enner",
                        "last-name": "Valencia"
                    }
                }
            ]
        },
        "round": {
            "matchday": 1,
            "name": "Group A"
        },
        "status": {
            "short": "FT",
            "updated": true,
            "points_calculated": true,
            "round_started": true
        }
    },
    ...
    
    ]
}
```

### Get teams
Used to retrieve all teams in the tournament, with their names and team IDs.
All teams are sorted in aalphabetical order

##### - Request
```
Method: GET
URL: http://10.212.170.213:8080/otp/getteams
```

##### - Response
```
[
     {
        "name": "Argentina",
        "score": 0,
        "id": 556,
        "goals": null
    },
    {
        "name": "Australia",
        "score": 0,
        "id": 927,
        "goals": null
    },
    {
        "name": "Belgium",
        "score": 0,
        "id": 507,
        "goals": null
    }, ... // The rest of the teams here
]
```


### Get players
Used to retrieve a list of teams including each of their players playing in the tournament.
Teams are sorted alphabetically, as well as players within each team.

##### - Request
```
Method: GET
URL: http://10.212.170.213:8080/otp/getallplayers
```

##### - Response
```
"teams": [
        {
            "team_id": 556,
            "team_name": "Argentina",
            "players": [
                {
                    "id": 457832,
                    "first-name": "Alejandro",
                    "last-name": "Gomez",
                    "team_id": 556
                },
                {
                    "id": 917353,
                    "first-name": "Alexis",
                    "last-name": "Mac Allister",
                    "team_id": 556
                },
                {
                    "id": 39889,
                    "first-name": "Angel",
                    "last-name": "Correa",
                    "team_id": 556
                }, ...  // Rest of Argentina players
            ]
        }, 
        {
            "team_id": 927,
            "team_name": "Australia",
            "players": [
                {
                    "id": 936137,
                    "first-name": "Andrew",
                    "last-name": "Redmayne",
                    "team_id": 927
                },
                {
                    "id": 931428,
                    "first-name": "Cameron",
                    "last-name": "Devlin",
                    "team_id": 927
                },
                {
                    "id": 936140,
                    "first-name": "Conor",
                    "last-name": "Metcalfe",
                    "team_id": 927
                }, ... // Rest of Australia players
            ]
        }, ... // Rest of teams
```

### Get homepage info
Used to retrieve a short summary of info regarding a single user.

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/gethomepageinfo
Headers:
{
    Key: token
    Value: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE4NCwiZXhwIjoxNjcwMTYyNzU4fQ.YG7yaofh_9Q96YTiidou-E2uhkIlWPi6pkgEMVw_Vns
}
(The token must be a valid token recieved from logging in, expires after 30 mins)
```

##### - Response
```
     "userid": 184,
    "matches": {
        "matches": [
            {
                "id": 427911,
                "date": "2022-12-04",
                "time": "15:00",
                "home-team": {
                    "name": "France",
                    "score": 0,
                    "id": 516,
                    "goals": null
                },
                "away-team": {
                    "name": "Poland",
                    "score": 0,
                    "id": 885,
                    "goals": null
                },
                "round": {
                    "matchday": 4,
                    "name": "Round of 16"
                },
                "status": {
                    "short": "3pm",
                    "updated": false,
                    "points_calculated": false,
                    "round_started": true
                }
            }, ... // Rest of the next 4 matches
        ],
        "LastUpdated": "2022-12-04T14:26:32.73487462Z"
    },
    "predictions": {
        "userid": 184,
        "predictions": [
            {
                "matchid": 423193,
                "homescore": 0,
                "awayscore": 2,
                "points": 0
            },
            {
                "matchid": 423194,
                "homescore": 1,
                "awayscore": 1,
                "points": 0
            }, ... // Rest of predictions
        ]
    },
    "topscorers": {
        "userid": 184,
        "topscorers": null
    },
    "medals": {
        "userid": 184,
        "gold": 0,
        "silver": 0,
        "bronze": 0
    },
    "leagues": [
        {
            "league_id": 1,
            "league_code": "overall",
            "league_name": "overall",
            "players": [
                {
                    "userid": 179,
                    "user_name": "kyrre",
                    "score": 132
                },
                {
                    "userid": 145,
                    "user_name": "sanderth",
                    "score": 114
                },
                {
                    "userid": 172,
                    "user_name": "dennisw",
                    "score": 112
                }, ... // Rest of players in league
            ]
        }, ... // Rest of leagues
        
    ]
}
```

## User endpoints
Endpoints related to users, including creating users and logging in.


### Newuser
Used to create a new user

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/newuser
Body:
{
    "username": "Testuser10",
    "email": "testuser10@mail.com",
    "password": "MyStrongPassword123"
}
```

##### - Response
```
    Status 201: Created

    New user created
```



### Login
Used to log into an existing account

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/login
Body:
{
    "username": "Testuser10",
    "password": "MyStrongPassword123"
}
```

##### - Response
```
    Status 202: Accepted

    (Body contains a string used as an authorization token in later requests)
    Eg:
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0NCwiZXhwIjoxNjcwMTYyMTIxfQ.G86sNwpPEKtByFFYii3i1BO_OIyVsapA0NJC4sgvCnk"
```

## Prediction related endpoints

### Get match predictions
Used to retrieve information about what predictions a specified user has set for every match

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/getpredictions
Body:
{
    "userid": 184
}
```

##### - Response
```
    Status 200: OK

    Body contains all game predictions for the specified user.
    Eg:
    {
    "userid": 184,
    "predictions": [
        {
            "matchid": 423193,
            "homescore": 0,
            "awayscore": 2,
            "points": 6
        },
        {
            "matchid": 407105,
            "homescore": 2,
            "awayscore": 0,
            "points": 2
        },
        ...
        ]
    }
```

### Save match predictions
Used to save new match predictions for a user.

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/savepredictions
Headers:
{
    Key: token
    Value: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE4NCwiZXhwIjoxNjcwMTYyNzU4fQ.YG7yaofh_9Q96YTiidou-E2uhkIlWPi6pkgEMVw_Vns
}
(The token must be a valid token recieved from logging in, expires after 30 mins)

Body:
{
    "predictions": [
        {
            "matchid": 423193,
            "homescore": 0,
            "awayscore": 2
        },
        {
            "matchid": 423194,
            "homescore": 1,
            "awayscore": 1
        }
    ]
}
```

##### - Response
```
    Status 200: OK
```


## Top scorer related endpoints

### Get top scorers
Used to retrieve information about what players a specified user has set for as their predicted top scorers

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/gettopscorers
Body:
{
    "userid": 143
}
```

##### - Response
```
    Status 200: OK

    Body contains the three topscorers for the specified user, 
    including information about the chosen players.
    
    Eg:
    {
    "userid": 143,
    "topscorers": [
        {
            "player_id": 34290,
            "player_name": "Harry Kane",
            "team_id": 516,
            "team_name": "England",
            "goals": 0
        },
        {
            "player_id": 36756,
            "player_name": "Lionel Messi",
            "team_id": 556,
            "team_name": "Argentina",
            "goals": 2
        },
        {
            "player_id": 438567,
            "player_name": "Kylian Mbappe",
            "team_id": 521,
            "team_name": "France",
            "goals": 3
        }
    ]
}
```

### Save top scorers
Used to save a users desired top scorers.

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/savetopscorers
Body:
{
    "topscorers": [
    {
        "player_id": 34290,
        "player_name": "Harry Kane",
        "team_id": 516,
        "team_name": "England"
    },
    {
        "player_id": 36756,
        "player_name": "Lionel Messi",
        "team_id": 556,
        "team_name": "Argentina"
    },
    {
        "player_id": 438567,
        "player_name": "Kylian Mbappe",
        "team_id": 521,
        "team_name": "France"
    }
}
```

##### - Response
```
    Status 200: OK

    Body contains the three topscorers for the specified user, 
    including information about the chosen players.
    
    Eg:
    {
    "topscorers": [
        {
            "player_id": 34290,
            "player_name": "Harry Kane",
            "team_id": 516,
            "team_name": "England"
        },
        {
            "player_id": 36756,
            "player_name": "Lionel Messi",
            "team_id": 556,
            "team_name": "Argentina"
        },
        {
            "player_id": 438567,
            "player_name": "Kylian Mbappe",
            "team_id": 521,
            "team_name": "France"
        }
    ]
}
```

## Medal related endpoints

### Get medal predictions
Used to retrieve information about what teams the specified player has predicted to win the gold, silver and bronze medal in the tournament.

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/getmedals
Body:
{
    "userid": 143
}
```

##### - Response
```
    Status 200: OK

    Body contains the medal predictions for the specified user.
    
    Eg:
    {
        "userid": 143,
        "gold": 521,
        "silver": 523,
        "bronze": 559
    }
```

### Save medal predictions
Used to save a users predictions for the gold, silver and bronze medal winners.

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/savemedals
Headers:
{
    Key: token
    Value: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE4NCwiZXhwIjoxNjcwMTYyNzU4fQ.YG7yaofh_9Q96YTiidou-E2uhkIlWPi6pkgEMVw_Vns
}
(The token must be a valid token recieved from logging in, expires after 30 mins)
Body:
{
    "gold": 521,
    "silver": 523,
    "bronze": 559
}
```

##### - Response
```
    Status 200: OK

    Body contains the medal predictions for the specified user.
    Eg.
    {
        "userid": 143,
        "gold": 521,
        "silver": 523,
        "bronze": 559
    }
```


## League related endpoints

### Get leagues
Used to retrieve information about all the leagues a specified player is a part of.

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/getleagues
Headers:
{
    Key: token
    Value: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE4NCwiZXhwIjoxNjcwMTYyNzU4fQ.YG7yaofh_9Q96YTiidou-E2uhkIlWPi6pkgEMVw_Vns
}
(The token must be a valid token recieved from logging in, expires after 30 mins)
```

##### - Response
```
    Status 200: OK

    Body contains a list of leagues where the player is involved.
    The players in each league are sorted based on who has the most points.
    The player is specified through the authentication token.
    
    Eg:
    [
        {
        "league_id": 1,
        "league_code": "overall",
        "league_name": "overall",
        "players": [
            {
                "userid": 179,
                "user_name": "kyrre",
                "score": 132
            },
            {
                "userid": 145,
                "user_name": "sanderth",
                "score": 114
            },
            {
                "userid": 172,
                "user_name": "dennisw",
                "score": 112
            },
            {
                "userid": 143,
                "user_name": "martiran",
                "score": 111
            },... // More players from the league
        ]
        }, ... // More leagues
    ]
```

### Create league
Used to create a new league with a specified name. 
The player creating the league will automatically be added to the new league.

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/createleague
Headers:
{
    Key: token
    Value: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE4NCwiZXhwIjoxNjcwMTYyNzU4fQ.YG7yaofh_9Q96YTiidou-E2uhkIlWPi6pkgEMVw_Vns
}
(The token must be a valid token recieved from logging in, expires after 30 mins)
Body:
{
    "league_name": "MyNewLeague"
}
```

##### - Response
```
    Status 201: Created

    Body contains the information about the newly created league
    
    Eg:
    {
    "league_id": 340,
    "league_code": "0B37xi",
    "league_name": "MyNewLeague",
    "players": [
        {
            "userid": 184,
            "user_name": "",
            "score": 0
        }
    ]
    }
```

### Join league
Used to join an existing league with a specified league code. 

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/joinleague
Headers:
{
    Key: token
    Value: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE4NCwiZXhwIjoxNjcwMTYyNzU4fQ.YG7yaofh_9Q96YTiidou-E2uhkIlWPi6pkgEMVw_Vns
}
(The token must be a valid token recieved from logging in, expires after 30 mins)
Body:
{
    "league_code": "3xHT9H"
}
```

##### - Response
```
    Status 201: Created

    Body contains the information about the newly joined league
    
    Eg:
    {
    "league_id": 341,
    "league_code": "3xHT9H",
    "league_name": "testliga",
    "players": [
        {
            "userid": 145,
            "user_name": "sanderth",
            "score": 114
        },
        {
            "userid": 184,
            "user_name": "testuser10",
            "score": 0
        }
    ]
    }
```

### Leave league
Used to leave a league with a specified league id. 

##### - Request
```
Method: POST
URL: http://10.212.170.213:8080/otp/leaveleague
Headers:
{
    Key: token
    Value: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE4NCwiZXhwIjoxNjcwMTYyNzU4fQ.YG7yaofh_9Q96YTiidou-E2uhkIlWPi6pkgEMVw_Vns
}
(The token must be a valid token recieved from logging in, expires after 30 mins)
Body:
{
    "league_id": 341
}
```

##### - Response
```
    Status 200: OK

    League left.
```

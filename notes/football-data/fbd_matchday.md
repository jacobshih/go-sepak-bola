# matches

```
export FBDATA_TOKEN=
export FBDATA_COMPETITION=2021
export FBDATA_MATCHDAY=16
export FBDATA_MATCH=264497
```

```
curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/competitions/$FBDATA_COMPETITION/matches?matchday=$FBDATA_MATCHDAY
```

```
{
  "count": 10,
  "filters": {
    "matchday": "16"
  },
  "competition": {
    "id": 2021,
    "area": {
      "id": 2072,
      "name": "England"
    },
    "name": "Premier League",
    "code": "PL",
    "plan": "TIER_ONE",
    "lastUpdated": "2019-12-14T00:00:08Z"
  },
  "matches": [
    {
      "id": 264495,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-07T12:30:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-08T15:30:02Z",
      "score": {
        "winner": "HOME_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 3,
          "awayTeam": 1
        },
        "halfTime": {
          "homeTeam": 1,
          "awayTeam": 0
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 62,
        "name": "Everton FC"
      },
      "awayTeam": {
        "id": 61,
        "name": "Chelsea FC"
      },
      "referees": [
        {
          "id": 11585,
          "name": "Craig Pawson",
          "nationality": null
        },
        {
          "id": 11530,
          "name": "Lee Betts",
          "nationality": null
        },
        {
          "id": 11586,
          "name": "Richard West",
          "nationality": null
        },
        {
          "id": 11469,
          "name": "Darren England",
          "nationality": null
        },
        {
          "id": 11494,
          "name": "Stuart Attwell",
          "nationality": null
        }
      ]
    },
    {
      "id": 264494,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-07T15:00:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-08T15:35:02Z",
      "score": {
        "winner": "HOME_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 5,
          "awayTeam": 0
        },
        "halfTime": {
          "homeTeam": 3,
          "awayTeam": 0
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 73,
        "name": "Tottenham Hotspur FC"
      },
      "awayTeam": {
        "id": 328,
        "name": "Burnley FC"
      },
      "referees": [
        {
          "id": 11487,
          "name": "Kevin Friend",
          "nationality": null
        },
        {
          "id": 11595,
          "name": "Adrian Holmes",
          "nationality": null
        },
        {
          "id": 11521,
          "name": "Mark Scholes",
          "nationality": null
        },
        {
          "id": 11446,
          "name": "Robert Jones",
          "nationality": null
        },
        {
          "id": 11479,
          "name": "Lee Mason",
          "nationality": null
        }
      ]
    },
    {
      "id": 264499,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-07T15:00:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-08T15:35:02Z",
      "score": {
        "winner": "DRAW",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 0,
          "awayTeam": 0
        },
        "halfTime": {
          "homeTeam": 0,
          "awayTeam": 0
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 346,
        "name": "Watford FC"
      },
      "awayTeam": {
        "id": 354,
        "name": "Crystal Palace FC"
      },
      "referees": [
        {
          "id": 11551,
          "name": "Martin Atkinson",
          "nationality": null
        },
        {
          "id": 11606,
          "name": "Constantine Hatzidakis",
          "nationality": null
        },
        {
          "id": 11438,
          "name": "Dan Cook",
          "nationality": null
        },
        {
          "id": 11309,
          "name": "Peter Bankes",
          "nationality": null
        },
        {
          "id": 11610,
          "name": "Andre Marriner",
          "nationality": null
        }
      ]
    },
    {
      "id": 264500,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-07T15:00:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-08T15:35:02Z",
      "score": {
        "winner": "AWAY_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 0,
          "awayTeam": 3
        },
        "halfTime": {
          "homeTeam": 0,
          "awayTeam": 2
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 1044,
        "name": "AFC Bournemouth"
      },
      "awayTeam": {
        "id": 64,
        "name": "Liverpool FC"
      },
      "referees": [
        {
          "id": 11443,
          "name": "Chris Kavanagh",
          "nationality": null
        },
        {
          "id": 98555,
          "name": null,
          "nationality": null
        },
        {
          "id": 11504,
          "name": "Simon Long",
          "nationality": null
        },
        {
          "id": 11324,
          "name": "James Linington",
          "nationality": null
        },
        {
          "id": 11567,
          "name": "Jonathan Moss",
          "nationality": null
        }
      ]
    },
    {
      "id": 264492,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-07T17:30:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-08T15:40:03Z",
      "score": {
        "winner": "AWAY_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 1,
          "awayTeam": 2
        },
        "halfTime": {
          "homeTeam": 0,
          "awayTeam": 2
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 65,
        "name": "Manchester City FC"
      },
      "awayTeam": {
        "id": 66,
        "name": "Manchester United FC"
      },
      "referees": [
        {
          "id": 11580,
          "name": "Anthony Taylor",
          "nationality": null
        },
        {
          "id": 11581,
          "name": "Gary Beswick",
          "nationality": null
        },
        {
          "id": 11615,
          "name": "Adam Nunn",
          "nationality": null
        },
        {
          "id": 11575,
          "name": "Mike Dean",
          "nationality": null
        },
        {
          "id": 11605,
          "name": "Michael Oliver",
          "nationality": null
        }
      ]
    },
    {
      "id": 264491,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-08T14:00:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-08T23:59:26Z",
      "score": {
        "winner": "AWAY_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 1,
          "awayTeam": 4
        },
        "halfTime": {
          "homeTeam": 1,
          "awayTeam": 2
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 58,
        "name": "Aston Villa FC"
      },
      "awayTeam": {
        "id": 338,
        "name": "Leicester City FC"
      },
      "referees": [
        {
          "id": 11605,
          "name": "Michael Oliver",
          "nationality": null
        },
        {
          "id": 11564,
          "name": "Stuart Burt",
          "nationality": null
        },
        {
          "id": 11488,
          "name": "Simon Bennett",
          "nationality": null
        },
        {
          "id": 11520,
          "name": "Paul Tierney",
          "nationality": null
        },
        {
          "id": 11443,
          "name": "Chris Kavanagh",
          "nationality": null
        }
      ]
    },
    {
      "id": 264493,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-08T14:00:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-08T23:59:26Z",
      "score": {
        "winner": "HOME_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 2,
          "awayTeam": 1
        },
        "halfTime": {
          "homeTeam": 0,
          "awayTeam": 0
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 67,
        "name": "Newcastle United FC"
      },
      "awayTeam": {
        "id": 340,
        "name": "Southampton FC"
      },
      "referees": [
        {
          "id": 11556,
          "name": "David Coote",
          "nationality": null
        },
        {
          "id": 11570,
          "name": "Harry Lennard",
          "nationality": null
        },
        {
          "id": 11552,
          "name": "Peter Kirkup",
          "nationality": null
        },
        {
          "id": 23568,
          "name": "Jarred Gillett",
          "nationality": null
        },
        {
          "id": 11551,
          "name": "Martin Atkinson",
          "nationality": null
        }
      ]
    },
    {
      "id": 264496,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-08T14:00:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-08T23:59:26Z",
      "score": {
        "winner": "AWAY_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 1,
          "awayTeam": 2
        },
        "halfTime": {
          "homeTeam": 1,
          "awayTeam": 0
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 68,
        "name": "Norwich City FC"
      },
      "awayTeam": {
        "id": 356,
        "name": "Sheffield United FC"
      },
      "referees": [
        {
          "id": 11430,
          "name": "Simon Hooper",
          "nationality": null
        },
        {
          "id": 11424,
          "name": "Neil Davies",
          "nationality": null
        },
        {
          "id": 11480,
          "name": "Eddie Smart",
          "nationality": null
        },
        {
          "id": 14713,
          "name": "Oliver Yates",
          "nationality": null
        },
        {
          "id": 11309,
          "name": "Peter Bankes",
          "nationality": null
        }
      ]
    },
    {
      "id": 264498,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-08T16:30:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-08T23:59:26Z",
      "score": {
        "winner": "DRAW",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 2,
          "awayTeam": 2
        },
        "halfTime": {
          "homeTeam": 2,
          "awayTeam": 2
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 397,
        "name": "Brighton & Hove Albion FC"
      },
      "awayTeam": {
        "id": 76,
        "name": "Wolverhampton Wanderers FC"
      },
      "referees": [
        {
          "id": 11567,
          "name": "Jonathan Moss",
          "nationality": null
        },
        {
          "id": 11531,
          "name": "Marc Perry",
          "nationality": null
        },
        {
          "id": 11611,
          "name": "Scott Ledger",
          "nationality": null
        },
        {
          "id": 11396,
          "name": "Tim Robinson",
          "nationality": null
        },
        {
          "id": 11503,
          "name": "Graham Scott",
          "nationality": null
        }
      ]
    },
    {
      "id": 264497,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 17
      },
      "utcDate": "2019-12-09T20:00:00Z",
      "status": "FINISHED",
      "matchday": 16,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-10T01:40:01Z",
      "score": {
        "winner": "AWAY_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 1,
          "awayTeam": 3
        },
        "halfTime": {
          "homeTeam": 1,
          "awayTeam": 0
        },
        "extraTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "penalties": {
          "homeTeam": null,
          "awayTeam": null
        }
      },
      "homeTeam": {
        "id": 563,
        "name": "West Ham United FC"
      },
      "awayTeam": {
        "id": 57,
        "name": "Arsenal FC"
      },
      "referees": [
        {
          "id": 11575,
          "name": "Mike Dean",
          "nationality": null
        },
        {
          "id": 11495,
          "name": "Ian Hussin",
          "nationality": null
        },
        {
          "id": 11431,
          "name": "Daniel Robathan",
          "nationality": null
        },
        {
          "id": 11479,
          "name": "Lee Mason",
          "nationality": null
        },
        {
          "id": 11487,
          "name": "Kevin Friend",
          "nationality": null
        }
      ]
    }
  ]
}
```

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
    "matchday": "21"
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
    "lastUpdated": "2020-01-01T14:05:44Z"
  },
  "matches": [
    {
      "id": 264546,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-01T12:30:00Z",
      "status": "IN_PLAY",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2020-01-01T14:05:43Z",
      "score": {
        "winner": "AWAY_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 0,
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
        "id": 328,
        "name": "Burnley FC"
      },
      "awayTeam": {
        "id": 58,
        "name": "Aston Villa FC"
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
          "id": 11378,
          "name": "Tony Harrington",
          "nationality": null
        },
        {
          "id": 23568,
          "name": "Jarred Gillett",
          "nationality": null
        }
      ]
    },
    {
      "id": 264549,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-01T12:30:00Z",
      "status": "IN_PLAY",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2020-01-01T14:05:43Z",
      "score": {
        "winner": "AWAY_TEAM",
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": 0,
          "awayTeam": 1
        },
        "halfTime": {
          "homeTeam": 0,
          "awayTeam": 1
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
        "id": 61,
        "name": "Chelsea FC"
      },
      "referees": [
        {
          "id": 11494,
          "name": "Stuart Attwell",
          "nationality": null
        },
        {
          "id": 11530,
          "name": "Lee Betts",
          "nationality": null
        },
        {
          "id": 11615,
          "name": "Adam Nunn",
          "nationality": null
        },
        {
          "id": 9382,
          "name": "Gavin Ward",
          "nationality": null
        },
        {
          "id": 11556,
          "name": "David Coote",
          "nationality": null
        }
      ]
    },
    {
      "id": 264543,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-01T15:00:00Z",
      "status": "SCHEDULED",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2020-01-01T14:05:43Z",
      "score": {
        "winner": null,
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "halfTime": {
          "homeTeam": null,
          "awayTeam": null
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
        "id": 338,
        "name": "Leicester City FC"
      },
      "referees": [
        {
          "id": 11551,
          "name": "Martin Atkinson",
          "nationality": null
        },
        {
          "id": 11581,
          "name": "Gary Beswick",
          "nationality": null
        },
        {
          "id": 11425,
          "name": "Nicholas Hopton",
          "nationality": null
        },
        {
          "id": 58084,
          "name": "Leigh Doughty",
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
      "id": 264548,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-01T15:00:00Z",
      "status": "SCHEDULED",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2020-01-01T14:05:44Z",
      "score": {
        "winner": null,
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "halfTime": {
          "homeTeam": null,
          "awayTeam": null
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
        "id": 340,
        "name": "Southampton FC"
      },
      "awayTeam": {
        "id": 73,
        "name": "Tottenham Hotspur FC"
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
          "id": 11519,
          "name": "Keith Stroud",
          "nationality": null
        },
        {
          "id": 11580,
          "name": "Anthony Taylor",
          "nationality": null
        }
      ]
    },
    {
      "id": 264550,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-01T15:00:00Z",
      "status": "SCHEDULED",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2020-01-01T14:05:44Z",
      "score": {
        "winner": null,
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "halfTime": {
          "homeTeam": null,
          "awayTeam": null
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
        "id": 76,
        "name": "Wolverhampton Wanderers FC"
      },
      "referees": [
        {
          "id": 11423,
          "name": "Andy Madley",
          "nationality": null
        },
        {
          "id": 11595,
          "name": "Adrian Holmes",
          "nationality": null
        },
        {
          "id": 11544,
          "name": "Simon Beck",
          "nationality": null
        },
        {
          "id": 11497,
          "name": "Charles Breakspear",
          "nationality": null
        },
        {
          "id": 11430,
          "name": "Simon Hooper",
          "nationality": null
        }
      ]
    },
    {
      "id": 264542,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-01T17:30:00Z",
      "status": "SCHEDULED",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-31T00:00:01Z",
      "score": {
        "winner": null,
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "halfTime": {
          "homeTeam": null,
          "awayTeam": null
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
        "id": 62,
        "name": "Everton FC"
      },
      "referees": []
    },
    {
      "id": 264545,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-01T17:30:00Z",
      "status": "SCHEDULED",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-31T00:00:01Z",
      "score": {
        "winner": null,
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "halfTime": {
          "homeTeam": null,
          "awayTeam": null
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
        "id": 354,
        "name": "Crystal Palace FC"
      },
      "referees": []
    },
    {
      "id": 264547,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-01T17:30:00Z",
      "status": "SCHEDULED",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-31T00:00:02Z",
      "score": {
        "winner": null,
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "halfTime": {
          "homeTeam": null,
          "awayTeam": null
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
        "id": 1044,
        "name": "AFC Bournemouth"
      },
      "referees": []
    },
    {
      "id": 264544,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-01T20:00:00Z",
      "status": "SCHEDULED",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2019-12-31T00:00:02Z",
      "score": {
        "winner": null,
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "halfTime": {
          "homeTeam": null,
          "awayTeam": null
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
        "id": 57,
        "name": "Arsenal FC"
      },
      "awayTeam": {
        "id": 66,
        "name": "Manchester United FC"
      },
      "referees": []
    },
    {
      "id": 264541,
      "season": {
        "id": 468,
        "startDate": "2019-08-09",
        "endDate": "2020-05-17",
        "currentMatchday": 21
      },
      "utcDate": "2020-01-02T20:00:00Z",
      "status": "SCHEDULED",
      "matchday": 21,
      "stage": "REGULAR_SEASON",
      "group": "Regular Season",
      "lastUpdated": "2020-01-01T00:00:00Z",
      "score": {
        "winner": null,
        "duration": "REGULAR",
        "fullTime": {
          "homeTeam": null,
          "awayTeam": null
        },
        "halfTime": {
          "homeTeam": null,
          "awayTeam": null
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
        "id": 64,
        "name": "Liverpool FC"
      },
      "awayTeam": {
        "id": 356,
        "name": "Sheffield United FC"
      },
      "referees": []
    }
  ]
}
```

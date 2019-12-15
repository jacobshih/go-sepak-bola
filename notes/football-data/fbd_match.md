# matches

```
export FBDATA_TOKEN=
export FBDATA_COMPETITION=2021
export FBDATA_MATCHDAY=16
export FBDATA_MATCH=264497
```

```
curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/matches/$FBDATA_MATCH
```

```
{
  "head2head": {
    "numberOfMatches": 10,
    "totalGoals": 32,
    "homeTeam": {
      "wins": 2,
      "draws": 2,
      "losses": 6
    },
    "awayTeam": {
      "wins": 6,
      "draws": 2,
      "losses": 2
    }
  },
  "match": {
    "id": 264497,
    "competition": {
      "id": 2021,
      "name": "Premier League"
    },
    "season": {
      "id": 468,
      "startDate": "2019-08-09",
      "endDate": "2020-05-17",
      "currentMatchday": 16,
      "winner": null
    },
    "utcDate": "2019-12-09T20:00:00Z",
    "status": "FINISHED",
    "venue": "London Stadium",
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
}```

# standings

```
export FBDATA_TOKEN=
export FBDATA_COMPETITION=2021
export FBDATA_MATCHDAY=16
export FBDATA_MATCH=264497
```

```
curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/competitions/$FBDATA_COMPETITION/standings
```

```
{
  "filters": {},
  "competition": {
    "id": 2021,
    "area": {
      "id": 2072,
      "name": "England"
    },
    "name": "Premier League",
    "code": "PL",
    "plan": "TIER_ONE",
    "lastUpdated": "2019-12-10T01:40:01Z"
  },
  "season": {
    "id": 468,
    "startDate": "2019-08-09",
    "endDate": "2020-05-17",
    "currentMatchday": 16,
    "winner": null
  },
  "standings": [
    {
      "stage": "REGULAR_SEASON",
      "type": "TOTAL",
      "group": null,
      "table": [
        {
          "position": 1,
          "team": {
            "id": 64,
            "name": "Liverpool FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/0/0a/FC_Liverpool.svg"
          },
          "playedGames": 16,
          "won": 15,
          "draw": 1,
          "lost": 0,
          "points": 46,
          "goalsFor": 40,
          "goalsAgainst": 14,
          "goalDifference": 26
        },
        {
          "position": 2,
          "team": {
            "id": 338,
            "name": "Leicester City FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/en/6/63/Leicester02.png"
          },
          "playedGames": 16,
          "won": 12,
          "draw": 2,
          "lost": 2,
          "points": 38,
          "goalsFor": 39,
          "goalsAgainst": 10,
          "goalDifference": 29
        },
        {
          "position": 3,
          "team": {
            "id": 65,
            "name": "Manchester City FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/e/eb/Manchester_City_FC_badge.svg"
          },
          "playedGames": 16,
          "won": 10,
          "draw": 2,
          "lost": 4,
          "points": 32,
          "goalsFor": 44,
          "goalsAgainst": 19,
          "goalDifference": 25
        },
        {
          "position": 4,
          "team": {
            "id": 61,
            "name": "Chelsea FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/5/5c/Chelsea_crest.svg"
          },
          "playedGames": 16,
          "won": 9,
          "draw": 2,
          "lost": 5,
          "points": 29,
          "goalsFor": 31,
          "goalsAgainst": 24,
          "goalDifference": 7
        },
        {
          "position": 5,
          "team": {
            "id": 66,
            "name": "Manchester United FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/d/da/Manchester_United_FC.svg"
          },
          "playedGames": 16,
          "won": 6,
          "draw": 6,
          "lost": 4,
          "points": 24,
          "goalsFor": 25,
          "goalsAgainst": 19,
          "goalDifference": 6
        },
        {
          "position": 6,
          "team": {
            "id": 76,
            "name": "Wolverhampton Wanderers FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/f/fc/Wolverhampton_Wanderers.svg"
          },
          "playedGames": 16,
          "won": 5,
          "draw": 9,
          "lost": 2,
          "points": 24,
          "goalsFor": 23,
          "goalsAgainst": 19,
          "goalDifference": 4
        },
        {
          "position": 7,
          "team": {
            "id": 73,
            "name": "Tottenham Hotspur FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/b/b4/Tottenham_Hotspur.svg"
          },
          "playedGames": 16,
          "won": 6,
          "draw": 5,
          "lost": 5,
          "points": 23,
          "goalsFor": 30,
          "goalsAgainst": 23,
          "goalDifference": 7
        },
        {
          "position": 8,
          "team": {
            "id": 356,
            "name": "Sheffield United FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/9/9c/Sheffield_United_FC_logo.svg"
          },
          "playedGames": 16,
          "won": 5,
          "draw": 7,
          "lost": 4,
          "points": 22,
          "goalsFor": 19,
          "goalsAgainst": 16,
          "goalDifference": 3
        },
        {
          "position": 9,
          "team": {
            "id": 57,
            "name": "Arsenal FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/en/5/53/Arsenal_FC.svg"
          },
          "playedGames": 16,
          "won": 5,
          "draw": 7,
          "lost": 4,
          "points": 22,
          "goalsFor": 24,
          "goalsAgainst": 24,
          "goalDifference": 0
        },
        {
          "position": 10,
          "team": {
            "id": 354,
            "name": "Crystal Palace FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/b/bf/Crystal_Palace_F.C._logo_%282013%29.png"
          },
          "playedGames": 16,
          "won": 6,
          "draw": 4,
          "lost": 6,
          "points": 22,
          "goalsFor": 14,
          "goalsAgainst": 18,
          "goalDifference": -4
        },
        {
          "position": 11,
          "team": {
            "id": 67,
            "name": "Newcastle United FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/5/56/Newcastle_United_Logo.svg"
          },
          "playedGames": 16,
          "won": 6,
          "draw": 4,
          "lost": 6,
          "points": 22,
          "goalsFor": 17,
          "goalsAgainst": 23,
          "goalDifference": -6
        },
        {
          "position": 12,
          "team": {
            "id": 397,
            "name": "Brighton & Hove Albion FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/f/fd/Brighton_%26_Hove_Albion_logo.svg"
          },
          "playedGames": 16,
          "won": 5,
          "draw": 4,
          "lost": 7,
          "points": 19,
          "goalsFor": 20,
          "goalsAgainst": 24,
          "goalDifference": -4
        },
        {
          "position": 13,
          "team": {
            "id": 328,
            "name": "Burnley FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/0/02/Burnley_FC_badge.png"
          },
          "playedGames": 16,
          "won": 5,
          "draw": 3,
          "lost": 8,
          "points": 18,
          "goalsFor": 21,
          "goalsAgainst": 29,
          "goalDifference": -8
        },
        {
          "position": 14,
          "team": {
            "id": 62,
            "name": "Everton FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/f/f9/Everton_FC.svg"
          },
          "playedGames": 16,
          "won": 5,
          "draw": 2,
          "lost": 9,
          "points": 17,
          "goalsFor": 19,
          "goalsAgainst": 28,
          "goalDifference": -9
        },
        {
          "position": 15,
          "team": {
            "id": 1044,
            "name": "AFC Bournemouth",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/de/4/41/Afc_bournemouth.svg"
          },
          "playedGames": 16,
          "won": 4,
          "draw": 4,
          "lost": 8,
          "points": 16,
          "goalsFor": 18,
          "goalsAgainst": 24,
          "goalDifference": -6
        },
        {
          "position": 16,
          "team": {
            "id": 563,
            "name": "West Ham United FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/e/e0/West_Ham_United_FC.svg"
          },
          "playedGames": 16,
          "won": 4,
          "draw": 4,
          "lost": 8,
          "points": 16,
          "goalsFor": 18,
          "goalsAgainst": 28,
          "goalDifference": -10
        },
        {
          "position": 17,
          "team": {
            "id": 58,
            "name": "Aston Villa FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/9/9f/Aston_Villa_logo.svg"
          },
          "playedGames": 16,
          "won": 4,
          "draw": 3,
          "lost": 9,
          "points": 15,
          "goalsFor": 23,
          "goalsAgainst": 28,
          "goalDifference": -5
        },
        {
          "position": 18,
          "team": {
            "id": 340,
            "name": "Southampton FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/c/c9/FC_Southampton.svg"
          },
          "playedGames": 16,
          "won": 4,
          "draw": 3,
          "lost": 9,
          "points": 15,
          "goalsFor": 18,
          "goalsAgainst": 35,
          "goalDifference": -17
        },
        {
          "position": 19,
          "team": {
            "id": 68,
            "name": "Norwich City FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/8/8c/Norwich_City.svg"
          },
          "playedGames": 16,
          "won": 3,
          "draw": 2,
          "lost": 11,
          "points": 11,
          "goalsFor": 17,
          "goalsAgainst": 34,
          "goalDifference": -17
        },
        {
          "position": 20,
          "team": {
            "id": 346,
            "name": "Watford FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/e/e2/Watford.svg"
          },
          "playedGames": 16,
          "won": 1,
          "draw": 6,
          "lost": 9,
          "points": 9,
          "goalsFor": 9,
          "goalsAgainst": 30,
          "goalDifference": -21
        }
      ]
    },
    {
      "stage": "REGULAR_SEASON",
      "type": "HOME",
      "group": null,
      "table": [
        {
          "position": 1,
          "team": {
            "id": 64,
            "name": "Liverpool FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/0/0a/FC_Liverpool.svg"
          },
          "playedGames": 8,
          "won": 8,
          "draw": 0,
          "lost": 0,
          "points": 24,
          "goalsFor": 24,
          "goalsAgainst": 9,
          "goalDifference": 15
        },
        {
          "position": 2,
          "team": {
            "id": 338,
            "name": "Leicester City FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/en/6/63/Leicester02.png"
          },
          "playedGames": 8,
          "won": 7,
          "draw": 1,
          "lost": 0,
          "points": 22,
          "goalsFor": 18,
          "goalsAgainst": 4,
          "goalDifference": 14
        },
        {
          "position": 3,
          "team": {
            "id": 73,
            "name": "Tottenham Hotspur FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/b/b4/Tottenham_Hotspur.svg"
          },
          "playedGames": 8,
          "won": 5,
          "draw": 2,
          "lost": 1,
          "points": 17,
          "goalsFor": 19,
          "goalsAgainst": 7,
          "goalDifference": 12
        },
        {
          "position": 4,
          "team": {
            "id": 65,
            "name": "Manchester City FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/e/eb/Manchester_City_FC_badge.svg"
          },
          "playedGames": 8,
          "won": 5,
          "draw": 1,
          "lost": 2,
          "points": 16,
          "goalsFor": 22,
          "goalsAgainst": 8,
          "goalDifference": 14
        },
        {
          "position": 5,
          "team": {
            "id": 66,
            "name": "Manchester United FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/d/da/Manchester_United_FC.svg"
          },
          "playedGames": 8,
          "won": 4,
          "draw": 3,
          "lost": 1,
          "points": 15,
          "goalsFor": 15,
          "goalsAgainst": 8,
          "goalDifference": 7
        },
        {
          "position": 6,
          "team": {
            "id": 61,
            "name": "Chelsea FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/5/5c/Chelsea_crest.svg"
          },
          "playedGames": 8,
          "won": 4,
          "draw": 2,
          "lost": 2,
          "points": 14,
          "goalsFor": 11,
          "goalsAgainst": 7,
          "goalDifference": 4
        },
        {
          "position": 7,
          "team": {
            "id": 57,
            "name": "Arsenal FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/en/5/53/Arsenal_FC.svg"
          },
          "playedGames": 8,
          "won": 3,
          "draw": 4,
          "lost": 1,
          "points": 13,
          "goalsFor": 14,
          "goalsAgainst": 12,
          "goalDifference": 2
        },
        {
          "position": 8,
          "team": {
            "id": 76,
            "name": "Wolverhampton Wanderers FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/f/fc/Wolverhampton_Wanderers.svg"
          },
          "playedGames": 8,
          "won": 3,
          "draw": 4,
          "lost": 1,
          "points": 13,
          "goalsFor": 12,
          "goalsAgainst": 10,
          "goalDifference": 2
        },
        {
          "position": 9,
          "team": {
            "id": 67,
            "name": "Newcastle United FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/5/56/Newcastle_United_Logo.svg"
          },
          "playedGames": 8,
          "won": 3,
          "draw": 4,
          "lost": 1,
          "points": 13,
          "goalsFor": 9,
          "goalsAgainst": 7,
          "goalDifference": 2
        },
        {
          "position": 10,
          "team": {
            "id": 62,
            "name": "Everton FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/f/f9/Everton_FC.svg"
          },
          "playedGames": 8,
          "won": 4,
          "draw": 1,
          "lost": 3,
          "points": 13,
          "goalsFor": 11,
          "goalsAgainst": 11,
          "goalDifference": 0
        },
        {
          "position": 11,
          "team": {
            "id": 397,
            "name": "Brighton & Hove Albion FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/f/fd/Brighton_%26_Hove_Albion_logo.svg"
          },
          "playedGames": 8,
          "won": 3,
          "draw": 3,
          "lost": 2,
          "points": 12,
          "goalsFor": 12,
          "goalsAgainst": 10,
          "goalDifference": 2
        },
        {
          "position": 12,
          "team": {
            "id": 328,
            "name": "Burnley FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/0/02/Burnley_FC_badge.png"
          },
          "playedGames": 8,
          "won": 4,
          "draw": 0,
          "lost": 4,
          "points": 12,
          "goalsFor": 12,
          "goalsAgainst": 13,
          "goalDifference": -1
        },
        {
          "position": 13,
          "team": {
            "id": 58,
            "name": "Aston Villa FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/9/9f/Aston_Villa_logo.svg"
          },
          "playedGames": 8,
          "won": 3,
          "draw": 2,
          "lost": 3,
          "points": 11,
          "goalsFor": 11,
          "goalsAgainst": 11,
          "goalDifference": 0
        },
        {
          "position": 14,
          "team": {
            "id": 354,
            "name": "Crystal Palace FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/b/bf/Crystal_Palace_F.C._logo_%282013%29.png"
          },
          "playedGames": 8,
          "won": 3,
          "draw": 2,
          "lost": 3,
          "points": 11,
          "goalsFor": 6,
          "goalsAgainst": 7,
          "goalDifference": -1
        },
        {
          "position": 15,
          "team": {
            "id": 356,
            "name": "Sheffield United FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/9/9c/Sheffield_United_FC_logo.svg"
          },
          "playedGames": 8,
          "won": 3,
          "draw": 1,
          "lost": 4,
          "points": 10,
          "goalsFor": 9,
          "goalsAgainst": 9,
          "goalDifference": 0
        },
        {
          "position": 16,
          "team": {
            "id": 1044,
            "name": "AFC Bournemouth",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/de/4/41/Afc_bournemouth.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 3,
          "lost": 3,
          "points": 9,
          "goalsFor": 9,
          "goalsAgainst": 12,
          "goalDifference": -3
        },
        {
          "position": 17,
          "team": {
            "id": 563,
            "name": "West Ham United FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/e/e0/West_Ham_United_FC.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 1,
          "lost": 5,
          "points": 7,
          "goalsFor": 11,
          "goalsAgainst": 17,
          "goalDifference": -6
        },
        {
          "position": 18,
          "team": {
            "id": 68,
            "name": "Norwich City FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/8/8c/Norwich_City.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 1,
          "lost": 5,
          "points": 7,
          "goalsFor": 13,
          "goalsAgainst": 20,
          "goalDifference": -7
        },
        {
          "position": 19,
          "team": {
            "id": 340,
            "name": "Southampton FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/c/c9/FC_Southampton.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 1,
          "lost": 5,
          "points": 7,
          "goalsFor": 9,
          "goalsAgainst": 23,
          "goalDifference": -14
        },
        {
          "position": 20,
          "team": {
            "id": 346,
            "name": "Watford FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/e/e2/Watford.svg"
          },
          "playedGames": 8,
          "won": 0,
          "draw": 4,
          "lost": 4,
          "points": 4,
          "goalsFor": 4,
          "goalsAgainst": 13,
          "goalDifference": -9
        }
      ]
    },
    {
      "stage": "REGULAR_SEASON",
      "type": "AWAY",
      "group": null,
      "table": [
        {
          "position": 1,
          "team": {
            "id": 64,
            "name": "Liverpool FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/0/0a/FC_Liverpool.svg"
          },
          "playedGames": 8,
          "won": 7,
          "draw": 1,
          "lost": 0,
          "points": 22,
          "goalsFor": 16,
          "goalsAgainst": 5,
          "goalDifference": 11
        },
        {
          "position": 2,
          "team": {
            "id": 338,
            "name": "Leicester City FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/en/6/63/Leicester02.png"
          },
          "playedGames": 8,
          "won": 5,
          "draw": 1,
          "lost": 2,
          "points": 16,
          "goalsFor": 21,
          "goalsAgainst": 6,
          "goalDifference": 15
        },
        {
          "position": 3,
          "team": {
            "id": 65,
            "name": "Manchester City FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/e/eb/Manchester_City_FC_badge.svg"
          },
          "playedGames": 8,
          "won": 5,
          "draw": 1,
          "lost": 2,
          "points": 16,
          "goalsFor": 22,
          "goalsAgainst": 11,
          "goalDifference": 11
        },
        {
          "position": 4,
          "team": {
            "id": 61,
            "name": "Chelsea FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/5/5c/Chelsea_crest.svg"
          },
          "playedGames": 8,
          "won": 5,
          "draw": 0,
          "lost": 3,
          "points": 15,
          "goalsFor": 20,
          "goalsAgainst": 17,
          "goalDifference": 3
        },
        {
          "position": 5,
          "team": {
            "id": 356,
            "name": "Sheffield United FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/9/9c/Sheffield_United_FC_logo.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 6,
          "lost": 0,
          "points": 12,
          "goalsFor": 10,
          "goalsAgainst": 7,
          "goalDifference": 3
        },
        {
          "position": 6,
          "team": {
            "id": 76,
            "name": "Wolverhampton Wanderers FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/f/fc/Wolverhampton_Wanderers.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 5,
          "lost": 1,
          "points": 11,
          "goalsFor": 11,
          "goalsAgainst": 9,
          "goalDifference": 2
        },
        {
          "position": 7,
          "team": {
            "id": 354,
            "name": "Crystal Palace FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/b/bf/Crystal_Palace_F.C._logo_%282013%29.png"
          },
          "playedGames": 8,
          "won": 3,
          "draw": 2,
          "lost": 3,
          "points": 11,
          "goalsFor": 8,
          "goalsAgainst": 11,
          "goalDifference": -3
        },
        {
          "position": 8,
          "team": {
            "id": 66,
            "name": "Manchester United FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/d/da/Manchester_United_FC.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 3,
          "lost": 3,
          "points": 9,
          "goalsFor": 10,
          "goalsAgainst": 11,
          "goalDifference": -1
        },
        {
          "position": 9,
          "team": {
            "id": 57,
            "name": "Arsenal FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/en/5/53/Arsenal_FC.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 3,
          "lost": 3,
          "points": 9,
          "goalsFor": 10,
          "goalsAgainst": 12,
          "goalDifference": -2
        },
        {
          "position": 10,
          "team": {
            "id": 563,
            "name": "West Ham United FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/e/e0/West_Ham_United_FC.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 3,
          "lost": 3,
          "points": 9,
          "goalsFor": 7,
          "goalsAgainst": 11,
          "goalDifference": -4
        },
        {
          "position": 11,
          "team": {
            "id": 67,
            "name": "Newcastle United FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/5/56/Newcastle_United_Logo.svg"
          },
          "playedGames": 8,
          "won": 3,
          "draw": 0,
          "lost": 5,
          "points": 9,
          "goalsFor": 8,
          "goalsAgainst": 16,
          "goalDifference": -8
        },
        {
          "position": 12,
          "team": {
            "id": 340,
            "name": "Southampton FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/c/c9/FC_Southampton.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 2,
          "lost": 4,
          "points": 8,
          "goalsFor": 9,
          "goalsAgainst": 12,
          "goalDifference": -3
        },
        {
          "position": 13,
          "team": {
            "id": 1044,
            "name": "AFC Bournemouth",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/de/4/41/Afc_bournemouth.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 1,
          "lost": 5,
          "points": 7,
          "goalsFor": 9,
          "goalsAgainst": 12,
          "goalDifference": -3
        },
        {
          "position": 14,
          "team": {
            "id": 397,
            "name": "Brighton & Hove Albion FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/f/fd/Brighton_%26_Hove_Albion_logo.svg"
          },
          "playedGames": 8,
          "won": 2,
          "draw": 1,
          "lost": 5,
          "points": 7,
          "goalsFor": 8,
          "goalsAgainst": 14,
          "goalDifference": -6
        },
        {
          "position": 15,
          "team": {
            "id": 73,
            "name": "Tottenham Hotspur FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/b/b4/Tottenham_Hotspur.svg"
          },
          "playedGames": 8,
          "won": 1,
          "draw": 3,
          "lost": 4,
          "points": 6,
          "goalsFor": 11,
          "goalsAgainst": 16,
          "goalDifference": -5
        },
        {
          "position": 16,
          "team": {
            "id": 328,
            "name": "Burnley FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/0/02/Burnley_FC_badge.png"
          },
          "playedGames": 8,
          "won": 1,
          "draw": 3,
          "lost": 4,
          "points": 6,
          "goalsFor": 9,
          "goalsAgainst": 16,
          "goalDifference": -7
        },
        {
          "position": 17,
          "team": {
            "id": 346,
            "name": "Watford FC",
            "crestUrl": "https://upload.wikimedia.org/wikipedia/en/e/e2/Watford.svg"
          },
          "playedGames": 8,
          "won": 1,
          "draw": 2,
          "lost": 5,
          "points": 5,
          "goalsFor": 5,
          "goalsAgainst": 17,
          "goalDifference": -12
        },
        {
          "position": 18,
          "team": {
            "id": 58,
            "name": "Aston Villa FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/9/9f/Aston_Villa_logo.svg"
          },
          "playedGames": 8,
          "won": 1,
          "draw": 1,
          "lost": 6,
          "points": 4,
          "goalsFor": 12,
          "goalsAgainst": 17,
          "goalDifference": -5
        },
        {
          "position": 19,
          "team": {
            "id": 62,
            "name": "Everton FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/f/f9/Everton_FC.svg"
          },
          "playedGames": 8,
          "won": 1,
          "draw": 1,
          "lost": 6,
          "points": 4,
          "goalsFor": 8,
          "goalsAgainst": 17,
          "goalDifference": -9
        },
        {
          "position": 20,
          "team": {
            "id": 68,
            "name": "Norwich City FC",
            "crestUrl": "http://upload.wikimedia.org/wikipedia/de/8/8c/Norwich_City.svg"
          },
          "playedGames": 8,
          "won": 1,
          "draw": 1,
          "lost": 6,
          "points": 4,
          "goalsFor": 4,
          "goalsAgainst": 14,
          "goalDifference": -10
        }
      ]
    }
  ]
}
```

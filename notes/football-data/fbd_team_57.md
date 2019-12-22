# teams

```
export FBDATA_TOKEN=
export FBDATA_COMPETITION=2021
export FBDATA_MATCHDAY=16
export FBDATA_MATCH=264497
export FBDATA_TEAM=57
```

```
curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/teams/$FBDATA_TEAM
```

```
{
  "id": 57,
  "area": {
    "id": 2072,
    "name": "England"
  },
  "activeCompetitions": [
    {
      "id": 2021,
      "area": {
        "id": 2072,
        "name": "England"
      },
      "name": "Premier League",
      "code": "PL",
      "plan": "TIER_ONE",
      "lastUpdated": "2019-12-21T23:59:25Z"
    },
    {
      "id": 2139,
      "area": {
        "id": 2072,
        "name": "England"
      },
      "name": "Football League Cup",
      "code": "FLC",
      "plan": "TIER_THREE",
      "lastUpdated": "2019-12-19T14:30:00Z"
    },
    {
      "id": 2146,
      "area": {
        "id": 2077,
        "name": "Europe"
      },
      "name": "UEFA Europa League",
      "code": "EL",
      "plan": "TIER_TWO",
      "lastUpdated": "2019-12-13T01:00:01Z"
    }
  ],
  "name": "Arsenal FC",
  "shortName": "Arsenal",
  "tla": "ARS",
  "crestUrl": "http://upload.wikimedia.org/wikipedia/en/5/53/Arsenal_FC.svg",
  "address": "75 Drayton Park London N5 1BU",
  "phone": "+44 (020) 76195003",
  "website": "http://www.arsenal.com",
  "email": "info@arsenal.co.uk",
  "founded": 1886,
  "clubColors": "Red / White",
  "venue": "Emirates Stadium",
  "squad": [
    {
      "id": 3141,
      "name": "Emiliano Martínez",
      "position": "Goalkeeper",
      "dateOfBirth": "1992-09-02T00:00:00Z",
      "countryOfBirth": "Argentina",
      "nationality": "Argentina",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 3174,
      "name": "Bernd Leno",
      "position": "Goalkeeper",
      "dateOfBirth": "1992-03-04T00:00:00Z",
      "countryOfBirth": "Germany",
      "nationality": "Germany",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 7779,
      "name": "Matt Macey",
      "position": "Goalkeeper",
      "dateOfBirth": "1994-09-09T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 137,
      "name": "Sokratis Papastathopoulos",
      "position": "Defender",
      "dateOfBirth": "1988-06-09T00:00:00Z",
      "countryOfBirth": "Greece",
      "nationality": "Greece",
      "shirtNumber": 5,
      "role": "PLAYER"
    },
    {
      "id": 7783,
      "name": "Héctor Bellerín",
      "position": "Defender",
      "dateOfBirth": "1995-03-19T00:00:00Z",
      "countryOfBirth": "Spain",
      "nationality": "Spain",
      "shirtNumber": 2,
      "role": "PLAYER"
    },
    {
      "id": 7784,
      "name": "Rob Holding",
      "position": "Defender",
      "dateOfBirth": "1995-09-20T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": 16,
      "role": "PLAYER"
    },
    {
      "id": 7785,
      "name": "Shkodran Mustafi",
      "position": "Defender",
      "dateOfBirth": "1992-04-17T00:00:00Z",
      "countryOfBirth": "Germany",
      "nationality": "Germany",
      "shirtNumber": 20,
      "role": "PLAYER"
    },
    {
      "id": 7786,
      "name": "Sead Kolašinac",
      "position": "Defender",
      "dateOfBirth": "1993-06-20T00:00:00Z",
      "countryOfBirth": "Germany",
      "nationality": "Bosnia and Herzegovina",
      "shirtNumber": 31,
      "role": "PLAYER"
    },
    {
      "id": 7787,
      "name": "Calum Chambers",
      "position": "Defender",
      "dateOfBirth": "1995-01-20T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": 21,
      "role": "PLAYER"
    },
    {
      "id": 7789,
      "name": "Konstantinos Mavropanos",
      "position": "Defender",
      "dateOfBirth": "1997-12-11T00:00:00Z",
      "countryOfBirth": "Greece",
      "nationality": "Greece",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 7805,
      "name": "David Luiz",
      "position": "Defender",
      "dateOfBirth": "1987-04-22T00:00:00Z",
      "countryOfBirth": "Brazil",
      "nationality": "Brazil",
      "shirtNumber": 23,
      "role": "PLAYER"
    },
    {
      "id": 15615,
      "name": "Kieran Tierney",
      "position": "Defender",
      "dateOfBirth": "1997-06-05T00:00:00Z",
      "countryOfBirth": "Isle of Man",
      "nationality": "Scotland",
      "shirtNumber": 3,
      "role": "PLAYER"
    },
    {
      "id": 98411,
      "name": "Zechariah Medley",
      "position": "Defender",
      "dateOfBirth": "2000-07-07T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 131038,
      "name": "Tolaji Bola",
      "position": "Defender",
      "dateOfBirth": "1999-01-04T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 68,
      "name": "Dani Ceballos",
      "position": "Midfielder",
      "dateOfBirth": "1996-08-07T00:00:00Z",
      "countryOfBirth": "Spain",
      "nationality": "Spain",
      "shirtNumber": 8,
      "role": "PLAYER"
    },
    {
      "id": 600,
      "name": "Mattéo Guendouzi",
      "position": "Midfielder",
      "dateOfBirth": "1999-04-14T00:00:00Z",
      "countryOfBirth": "France",
      "nationality": "France",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 2260,
      "name": "Lucas Torreira",
      "position": "Midfielder",
      "dateOfBirth": "1996-02-11T00:00:00Z",
      "countryOfBirth": "Uruguay",
      "nationality": "Uruguay",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 3180,
      "name": "Mesut Özil",
      "position": "Midfielder",
      "dateOfBirth": "1988-10-15T00:00:00Z",
      "countryOfBirth": "Germany",
      "nationality": "Germany",
      "shirtNumber": 10,
      "role": "PLAYER"
    },
    {
      "id": 3477,
      "name": "Granit Xhaka",
      "position": "Midfielder",
      "dateOfBirth": "1992-09-27T00:00:00Z",
      "countryOfBirth": "Switzerland",
      "nationality": "Switzerland",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 7792,
      "name": "Ainsley Maitland-Niles",
      "position": "Midfielder",
      "dateOfBirth": "1997-08-29T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": 15,
      "role": "PLAYER"
    },
    {
      "id": 61450,
      "name": "Martinelli",
      "position": "Midfielder",
      "dateOfBirth": "2001-06-18T00:00:00Z",
      "countryOfBirth": "Brazil",
      "nationality": "Brazil",
      "shirtNumber": 35,
      "role": "PLAYER"
    },
    {
      "id": 85570,
      "name": "Emile Smith-Rowe",
      "position": "Midfielder",
      "dateOfBirth": "2000-07-28T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": 32,
      "role": "PLAYER"
    },
    {
      "id": 99813,
      "name": "Bukayo Saka",
      "position": "Midfielder",
      "dateOfBirth": "2001-09-05T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": 77,
      "role": "PLAYER"
    },
    {
      "id": 131039,
      "name": "Robbie Burton",
      "position": "Midfielder",
      "dateOfBirth": "1999-12-26T00:00:00Z",
      "countryOfBirth": "Wales",
      "nationality": "Wales",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 137016,
      "name": "James Olayinka",
      "position": "Midfielder",
      "dateOfBirth": "2000-10-05T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 7797,
      "name": "Alexandre Lacazette",
      "position": "Attacker",
      "dateOfBirth": "1991-05-28T00:00:00Z",
      "countryOfBirth": "France",
      "nationality": "France",
      "shirtNumber": 9,
      "role": "PLAYER"
    },
    {
      "id": 7798,
      "name": "Joe Willock",
      "position": "Attacker",
      "dateOfBirth": "1999-08-20T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": 28,
      "role": "PLAYER"
    },
    {
      "id": 7799,
      "name": "Reiss Nelson",
      "position": "Attacker",
      "dateOfBirth": "1999-12-10T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": 24,
      "role": "PLAYER"
    },
    {
      "id": 7801,
      "name": "Pierre-Emerick Aubameyang",
      "position": "Attacker",
      "dateOfBirth": "1989-06-18T00:00:00Z",
      "countryOfBirth": "France",
      "nationality": "Gabon",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 8412,
      "name": "Nicolas Pépé",
      "position": "Attacker",
      "dateOfBirth": "1995-05-29T00:00:00Z",
      "countryOfBirth": "France",
      "nationality": "Côte d’Ivoire",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 99809,
      "name": "Tyreece John-Jules",
      "position": "Attacker",
      "dateOfBirth": "2001-02-14T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 131040,
      "name": "Folarin Balogun",
      "position": "Attacker",
      "dateOfBirth": "2001-07-03T00:00:00Z",
      "countryOfBirth": "England",
      "nationality": "England",
      "shirtNumber": null,
      "role": "PLAYER"
    },
    {
      "id": 136402,
      "name": "Fredrik Ljungberg",
      "position": null,
      "dateOfBirth": "1977-04-16T00:00:00Z",
      "countryOfBirth": "Sweden",
      "nationality": "Sweden",
      "shirtNumber": null,
      "role": "INTERIM_COACH"
    }
  ],
  "lastUpdated": "2019-12-19T02:24:23Z"
}
```

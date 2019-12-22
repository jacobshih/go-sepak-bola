# football-data notes

## references

[foorball-data.org API documentation](https://www.football-data.org/documentation/quickstart)

## source.me

```
export FBDATA_TOKEN=
export FBDATA_COMPETITION=2021
export FBDATA_MATCHDAY=16
export FBDATA_MATCH=264497
```

## api

### competitions

- List all available competitions.

    ```
    curl -X GET http://api.football-data.org/v2/competitions/
    ```

- List one particular competition.

    ```
    curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/competitions/$FBDATA_COMPETITION
    ```

### teams

- List all teams for a particular competition.

    ```
    curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/competitions/$FBDATA_COMPETITION/teams
    ```

- Show one particular team.

    ```
    curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/teams/$FBDATA_TEAM

    ```

### matches

- List all matches for a particular competition.

    ```
    curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/competitions/$FBDATA_COMPETITION/matches?matchday=$FBDATA_MATCHDAY
    ```

- show a particular match.

    ```
    curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/matches/$FBDATA_MATCH
    ```

### standings

- Show Standings for a particular competition.

    ```
    curl -X GET -H "X-Auth-Token: $FBDATA_TOKEN" http://api.football-data.org/v2/competitions/$FBDATA_COMPETITION/standings
    ```

## resources

### logos


- [premier league](https://www.thesportsdb.com/images/media/league/badge/i6o0kh1549879062.png)

    <img src="https://www.thesportsdb.com/images/media/league/badge/i6o0kh1549879062.png" width="192" />

- [bundesliga](https://www.thesportsdb.com/images/media/league/badge/0j55yv1534764799.png)

    <img src="https://www.thesportsdb.com/images/media/league/badge/0j55yv1534764799.png" width="192" />

- [primera division](https://www.thesportsdb.com/images/media/league/badge/7onmyv1534768460.png)

    <img src="https://www.thesportsdb.com/images/media/league/badge/7onmyv1534768460.png" width="192" />

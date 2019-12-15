package fbd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// POST http post
	POST string = "POST"
	// GET http get
	GET string = "GET"
)

const (
	uriCompetitions string = "http://api.football-data.org/v2/competitions/"
	uriCompetition  string = "http://api.football-data.org/v2/competitions/%d"
	uriTeams        string = "http://api.football-data.org/v2/competitions/%d/teams"
	uriMatches      string = "http://api.football-data.org/v2/competitions/%d/matches"
	uriMatch        string = "http://api.football-data.org/v2/matches/%d"
	uriStandings    string = "http://api.football-data.org/v2/competitions/%d/standings"
)

// FBDAPI type
type FBDAPI struct {
	token string
}

var fbdapi FBDAPI

// Instance returns pointer of FBDataAPI instance.
func Instance() *FBDAPI {
	return &fbdapi
}

// SetToken sets token for football-data.org to instance of FBDAPI.
func (fbd *FBDAPI) SetToken(token string) {
	fbd.token = token
}

// GetToken returns token for football-data.org
func (fbd *FBDAPI) GetToken() string {
	return fbd.token
}

// httpGet sends http request with method GET.
func (fbd *FBDAPI) httpGet(uri string) (content []byte) {
	req, err := http.NewRequest(GET, uri, nil)
	if err != nil {
		fmt.Println(err.Error())
		return content
	}
	req.Header.Set("X-Auth-Token", fbd.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return content
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		content, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return content
		}
	}
	return content
}

// GetCompetitions http://api.football-data.org/v2/competitions/
func (fbd *FBDAPI) GetCompetitions(id int) (content []byte) {
	uri := uriCompetitions
	if id > 0 {
		uri = fmt.Sprintf(uriCompetition, id)
	}
	content = fbd.httpGet(uri)
	return content
}

// GetMatches http://api.football-data.org/v2/competitions/$FBDATA_COMPETITION/standings
func (fbd *FBDAPI) GetMatches(id, matchday int) (content []byte) {
	if id > 0 {
		uri := fmt.Sprintf(uriMatches, id)
		if matchday > 0 {
			queryMatchday := fmt.Sprintf("?matchday=%d", matchday)
			uri += queryMatchday
		}
		content = fbd.httpGet(uri)
	}
	return content
}

// GetTeams http://api.football-data.org/v2/competitions/$FBDATA_COMPETITION/teams
func (fbd *FBDAPI) GetTeams(id int) (content []byte) {
	if id > 0 {
		uri := fmt.Sprintf(uriTeams, id)
		content = fbd.httpGet(uri)
	}
	return content
}

// GetStandings http://api.football-data.org/v2/competitions/$FBDATA_COMPETITION/standings
func (fbd *FBDAPI) GetStandings(id int) (content []byte) {
	if id > 0 {
		uri := fmt.Sprintf(uriStandings, id)
		content = fbd.httpGet(uri)
	}
	return content
}

// Area describes the area used for competitions, teams,...
type Area struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Team contains information of a particular team.
type Team struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	ShortName  string          `json:"shortName"`
	TLA        string          `json:"tla"`
	CrestURL   string          `json:"crestUrl"`
	Address    string          `json:"address"`
	Phone      string          `json:"phone"`
	Email      string          `json:"email"`
	Website    string          `json:"website"`
	Founded    int             `json:"founded"`
	ClubColors string          `json:"clubColors"`
	Venue      string          `json:"venue"`
	RawArea    json.RawMessage `json:"area"`
	Area       *Area
}

// Season contains information for a season.
type Season struct {
	ID              int             `json:"id"`
	StartDate       string          `json:"startDate"`
	EndDate         string          `json:"endDate"`
	CurrentMatchday int             `json:"currentMatchday"`
	RawWinner       json.RawMessage `json:"winner"`
	Winner          *Team
}

// Competition contains information for a particular competition.
type Competition struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Code       string          `json:"code"`
	EmblemURL  string          `json:"emblemUrl"`
	RawArea    json.RawMessage `json:"area"`
	RawSeason  json.RawMessage `json:"currentSeason"`
	RawSeasons json.RawMessage `json:"seasons"`
	Area       *Area
	Season     *Season
	Seasons    []*Season
	Teams      map[int]*Team
	Matches    []*Match
}

// GetTeams retrieves teams information for the particular competition.
func (comp *Competition) GetTeams() {
	var fbdteams TeamsData
	content := fbdteams.Get(comp.ID)
	if err := fbdteams.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "FBDTeams.Deserialize()", err)
		return
	}
	if comp.Teams == nil {
		comp.Teams = make(map[int]*Team)
	}
	for _, team := range fbdteams.Teams {
		comp.Teams[team.ID] = team
	}
}

// GetMatches retrieves matches information for the particular competition.
func (comp *Competition) GetMatches() {
	var matches MatchesData
	content := matches.Get(comp.ID, 0)
	if err := matches.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "FBDMatches.Deserialize()", err)
		return
	}
	comp.Matches = matches.Matches
}

// ScoreInfo contains information of the score of home/away team.
type ScoreInfo struct {
	HomeTeam int `json:"homeTeam"`
	AwayTeam int `json:"awayTeam"`
}

// Score contains detail score information.
type Score struct {
	Winner       string          `json:"winner"`
	Duration     string          `json:"duration"`
	RawFullTime  json.RawMessage `json:"fullTime"`
	RawHalfTime  json.RawMessage `json:"halfTime"`
	RawExtraTime json.RawMessage `json:"extraTime"`
	RawPenalties json.RawMessage `json:"penalties"`
	FullTime     *ScoreInfo
	HalfTime     *ScoreInfo
	ExtraTime    *ScoreInfo
	Penalties    *ScoreInfo
}

// Referee contains information for referee.
type Referee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
}

// Match contains detail information of a match.
type Match struct {
	ID          int             `json:"id"`
	RawSeason   json.RawMessage `json:"season"`
	UtcDate     string          `json:"utcDate"`
	Status      string          `json:"status"`
	Matchday    int             `json:"matchday"`
	Stage       string          `json:"stage"`
	Group       string          `json:"group"`
	RawScore    json.RawMessage `json:"score"`
	RawHomeTeam json.RawMessage `json:"homeTeam"`
	RawAwayTeam json.RawMessage `json:"awayTeam"`
	RawReferees json.RawMessage `json:"referees"`
	Season      *Season
	Score       *Score
	HomeTeam    *Team
	AwayTeam    *Team
	Referees    []*Referee
}

// MatchesData is a struct mapped to json string retrieved from api.football-data.org.
type MatchesData struct {
	Count          int             `json:"count"`
	RawCompetition json.RawMessage `json:"competition"`
	RawMatches     json.RawMessage `json:"matches"`
	Competition    *Competition
	Matches        []*Match
}

// Get retrieves matches information.
func (matches *MatchesData) Get(competitionID, matchday int) (content []byte) {
	content = fbdapi.GetMatches(competitionID, matchday)
	return content
}

// Deserialize decodes matches information from given json string.
func (matches *MatchesData) Deserialize(content []byte) (err error) {
	if err = json.Unmarshal(content, &matches); err != nil {
		return err
	}
	if err = json.Unmarshal(matches.RawCompetition, &matches.Competition); err != nil {
		return err
	}
	if err = json.Unmarshal(matches.RawMatches, &matches.Matches); err != nil {
		return err
	}
	for _, match := range matches.Matches {
		if err = json.Unmarshal(match.RawSeason, &match.Season); err != nil {
			return err
		}
		if err = json.Unmarshal(match.RawScore, &match.Score); err != nil {
			return err
		}
		if match.Score != nil {
			score := match.Score
			if err = json.Unmarshal(score.RawFullTime, &score.FullTime); err != nil {
				return err
			}
			if err = json.Unmarshal(score.RawHalfTime, &score.HalfTime); err != nil {
				return err
			}
			if err = json.Unmarshal(score.RawExtraTime, &score.ExtraTime); err != nil {
				return err
			}
			if err = json.Unmarshal(score.RawPenalties, &score.Penalties); err != nil {
				return err
			}
		}
		if err = json.Unmarshal(match.RawHomeTeam, &match.HomeTeam); err != nil {
			return err
		}
		if err = json.Unmarshal(match.RawAwayTeam, &match.AwayTeam); err != nil {
			return err
		}
		if err = json.Unmarshal(match.RawReferees, &match.Referees); err != nil {
			return err
		}
	}
	return err
}

// TableItem contains information for an item of standings table.
type TableItem struct {
	Position       int             `json:"position"`
	RawTeam        json.RawMessage `json:"team"`
	PlayedGames    int             `json:"playedGames"`
	Won            int             `json:"won"`
	Draw           int             `json:"draw"`
	Lost           int             `json:"lost"`
	Points         int             `json:"points"`
	GoalsFor       int             `json:"goalsFor"`
	GoalsAgainst   int             `json:"goalsAgainst"`
	GoalDifference int             `json:"goalDifference"`
	Team           *Team
}

// Standings contains information for standings table .
type Standings struct {
	Stage    string          `json:"stage"`
	Type     string          `json:"type"`
	Group    string          `json:"group"`
	RawTable json.RawMessage `json:"table"`
	Table    []*TableItem
}

// StandingsData is a struct mapped to json string retrieved from api.football-data.org.
type StandingsData struct {
	RawCompetition json.RawMessage `json:"competition"`
	RawSeason      json.RawMessage `json:"season"`
	RawStandings   json.RawMessage `json:"standings"`
	Competition    *Competition
	Season         *Season
	StandingsList  []*Standings
}

// Get retrieves standings of particular competition.
func (standings *StandingsData) Get(competitionID int) (content []byte) {
	content = fbdapi.GetStandings(competitionID)
	return content
}

// Deserialize decodes standings information from given json string.
func (standings *StandingsData) Deserialize(content []byte) (err error) {
	if err = json.Unmarshal(content, &standings); err != nil {
		return err
	}
	if err = json.Unmarshal(standings.RawCompetition, &standings.Competition); err != nil {
		return err
	}
	if err = json.Unmarshal(standings.RawSeason, &standings.Season); err != nil {
		return err
	}
	if err = json.Unmarshal(standings.RawStandings, &standings.StandingsList); err != nil {
		return err
	}
	for i, standingsItem := range standings.StandingsList {
		if err = json.Unmarshal(standingsItem.RawTable, &standingsItem.Table); err != nil {
			return err
		}
		for j, tableItem := range standingsItem.Table {
			if err = json.Unmarshal(tableItem.RawTeam, &tableItem.Team); err != nil {
				return err
			}
			standingsItem.Table[j] = tableItem
		}
		standings.StandingsList[i] = standingsItem
	}
	return err
}

// TeamsData is a struct mapped to json string retrieved from api.football-data.org.
type TeamsData struct {
	RawCompetition json.RawMessage `json:"competition"`
	RawSeason      json.RawMessage `json:"season"`
	RawTeams       json.RawMessage `json:"teams"`
	Competition    *Competition
	Season         *Season
	Teams          []*Team
}

// Get retrieves teams information.
func (teams *TeamsData) Get(competitionID int) (content []byte) {
	content = fbdapi.GetTeams(competitionID)
	return content
}

// Deserialize decodes teams information from given json string.
func (teams *TeamsData) Deserialize(content []byte) (err error) {
	if err = json.Unmarshal(content, &teams); err != nil {
		return err
	}
	if err = json.Unmarshal(teams.RawCompetition, &teams.Competition); err != nil {
		return err
	}
	if err = json.Unmarshal(teams.RawSeason, &teams.Season); err != nil {
		return err
	}
	if err = json.Unmarshal(teams.RawTeams, &teams.Teams); err != nil {
		return err
	}
	return err
}

// CompetitionsData is a struct mapped to json string retrieved from api.football-data.org.
type CompetitionsData struct {
	Count          int             `json:"count"`
	RawCompetition json.RawMessage `json:"competitions"`
	Competitions   []*Competition
}

// Get retrieves competitions information.
func (comps *CompetitionsData) Get() (content []byte) {
	content = fbdapi.GetCompetitions(0)
	return content
}

// Deserialize decodes competitions information from given json string.
func (comps *CompetitionsData) Deserialize(content []byte) (err error) {
	if err = json.Unmarshal(content, &comps); err != nil {
		return err
	}
	if err = json.Unmarshal(comps.RawCompetition, &comps.Competitions); err != nil {
		return err
	}
	for i, comp := range comps.Competitions {
		if err = json.Unmarshal(comp.RawArea, &comp.Area); err != nil {
			return err
		}
		if err = json.Unmarshal(comp.RawSeason, &comp.Season); err != nil {
			return err
		}
		comps.Competitions[i] = comp
	}
	return err
}

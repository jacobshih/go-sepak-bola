// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gsb

import (
	"fmt"
	"go-sepak-bola/internal/fbd"
	"net/url"
	"time"
)

// pkg constants
const (
	PkgName             string = "gsb"
	EmojiSoccer         string = "\xe2\x9a\xbd"
	EmojiTrophy         string = "\xf0\x9f\x8f\x86"
	EmojiSquaredVS      string = "\xf0\x9f\x86\x9a"
	ActionMatches       string = "matches"
	ActionStandings     string = "standings"
	ActionTeams         string = "teams"
	ActionTeam          string = "team"
	ActionMatchCalendar string = "match calendar"
	ActionMatchday      string = "matchday"
	ActionSquad         string = "squad"
	TextMatches         string = "Matches"
	TextStandings       string = "Standings"
	TextTeams           string = "Teams"
	TextMatchCalendar   string = "Match calendar"
	TextCurrentMatchday string = "Current matchday"
	TextMatchday        string = "Matchday"
	TextReferees        string = "Referees"
	TextSquad           string = "Squad"
	TextMore            string = "More"
	TextRound           string = "Round"
	TextPremierLeague   string = "Premier League"
	TextBundesliga      string = "Bundesliga"
	TextPrimeraDivision string = "Primera Division"
	CodePremierLeague   string = "PL"
	CodeBundesliga      string = "BL1"
	CodePrimeraDivision string = "PD"
	TextSpace           string = " "

	StatusScheduled string = "SCHEDULED"
	StatusInPlay    string = "IN_PLAY"
	StatusPaused    string = "PAUSED"
	StatusFinished  string = "FINISHED"
	StatusPostponed string = "POSTPONED"
	StatusCanceled  string = "CANCELED"
	StatusSuspended string = "SUSPENDED"
	StatusAwarded   string = "AWARDED"

	datetimeFormat string = "2006-01-02T15:04:05Z"
	dateFormat     string = "2006-01-02"
	timeFormat     string = "15:04:05"
)

// competition codes
const (
	PL  int = 2021
	BL1 int = 2002
	PD  int = 2014
)

// colors
// https://htmlcolorcodes.com/color-names/
const (
	ColorDodgeBlue   = "#1e90ff"
	ColorIndigo      = "#4b0082"
	ColorAero        = "#7cb9e8"
	ColorGreenYellow = "#adff2f"
	ColorAcidGreen   = "#b0bf1a"
	ColorOrange      = "#fb8c00"
	ColorAmber       = "#ff7e00"
	ColorLightGray   = "#c0c0c0"
	ColorWhite       = "#ffffff"
	ColorBlack       = "#000000"
	ColorGray        = "#7f7f7f"
	ColorBlue        = "#0000ff"
)

// SepakBola type SepakBola
type SepakBola struct {
	Competitions map[int]*fbd.Competition
}

// Init initialize SepakBola instance.
func (sepakbola *SepakBola) Init() {
	sepakbola.Competitions = make(map[int]*fbd.Competition)
	sepakbola.Competitions[PL] = new(fbd.Competition)  // 2021: Premier League
	sepakbola.Competitions[BL1] = new(fbd.Competition) // 2002: Bundesliga
	sepakbola.Competitions[PD] = new(fbd.Competition)  // 2014: Primera Division
}

// GetCompetitions gets intrested competitions.
func (sepakbola *SepakBola) GetCompetitions() {
	var comps fbd.CompetitionsData
	content := comps.Get()
	if err := comps.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "FBDCompetitions.Deserialize()", err)
		return
	}
	for i, comp := range comps.Competitions {
		if _, ok := sepakbola.Competitions[comp.ID]; ok {
			comps.Competitions[i].FixEmblem()
			sepakbola.Competitions[comp.ID] = comps.Competitions[i]
			sepakbola.Competitions[comp.ID].GetTeams()
			// sepakbola.Competitions[comp.Code].GetMatches()
			for i, team := range sepakbola.Competitions[comp.ID].Teams {
				u, err := url.Parse(team.CrestURL)
				if err == nil {
					if u.Scheme == "http" {
						u.Scheme = "https"
					}
					sepakbola.Competitions[comp.ID].Teams[i].CrestURL = u.String()
				}
			}
		}
	}
}

func toLocalTime(utcTime string) *time.Time {
	theTime, err := time.Parse(datetimeFormat, utcTime)
	if err == nil {
		localLoc, err := time.LoadLocation("Local")
		if err == nil {
			theTime = theTime.In(localLoc)
		}
	}
	return &theTime
}

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
	TextMatchday        string = "Matches"
	TextStandings       string = "Standings"
	TextTeams           string = "Teams"
	TextPremierLeague   string = "Premier League"
	TextBundesliga      string = "Bundesliga"
	TextPrimeraDivision string = "Primera Division"
	CodePremierLeague   string = "PL"
	CodeBundesliga      string = "BL1"
	CodePrimeraDivision string = "PD"
)

// competition codes
const (
	PL  int = 2021
	BL1 int = 2002
	PD  int = 2014
)

// colors
const (
	ColorDodgeBlue   = "#1e90ff"
	ColorGreenYellow = "#adff2f"
	ColorAmber       = "#ff7e00"
	ColorOrange      = "#fb8c00"
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
		}
	}
}

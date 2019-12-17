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

// constants
const (
	PL  int = 2021
	BL1 int = 2002
	PD  int = 2014
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

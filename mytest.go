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

package main

import (
	"fmt"
	"go-sepak-bola/gsb"
	"go-sepak-bola/internal/fbd"
	"go-sepak-bola/ui"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

func mytestFbdCompetitions() {
	// result
	fmt.Printf("%s\n", "=== competitions ===")
	for _, comp := range sepakbola.Competitions {
		fmt.Printf("(%4d) [%-3s] %s\n", comp.ID, comp.Code, comp.Name)
		fmt.Printf("%6s %s\n", "", comp.EmblemURL)
		if comp.Area != nil {
			fmt.Printf("%6s [%4d] %s\n", "", comp.Area.ID, comp.Area.Name)
		}
		if comp.Season != nil {
			fmt.Printf("%6s [%4d] %s - %s\n", "", comp.Season.ID, comp.Season.StartDate, comp.Season.EndDate)
		}
		if comp.Teams != nil {
			for _, team := range comp.Teams {
				fmt.Printf("%6d | %s | %4d | %-15s | %-26s | %s\n", team.ID, team.TLA, team.Founded, team.ShortName, team.Name, team.Website)
				fmt.Printf("%6s %-36s | %-20s | %s\n", "", team.Email, team.Phone, team.Address)
				fmt.Printf("%6s %s\n", "", team.CrestURL)
			}
		}
		if comp.Matches != nil {
			for _, match := range comp.Matches {
				if comp.Teams != nil {
					teams := comp.Teams
					fmt.Printf("(%6d) %2d %s %-15s %s %s\n", match.ID, match.Matchday, match.UtcDate, match.Status, teams[match.HomeTeam.ID].TLA, teams[match.AwayTeam.ID].TLA)
				} else {
					fmt.Printf("(%6d) %2d %s %-15s %s %s\n", match.ID, match.Matchday, match.UtcDate, match.Status, match.HomeTeam.Name, match.AwayTeam.Name)
				}
				if match.Score != nil {
					score := match.Score
					fmt.Printf("%8s %s:%s\n", "", "winner", score.Winner)
					fmt.Printf("%8s H(%d:%d) F(%d:%d) E(%d:%d) P(%d:%d)\n", "",
						score.HalfTime.HomeTeam, score.HalfTime.AwayTeam,
						score.FullTime.HomeTeam, score.FullTime.AwayTeam,
						score.ExtraTime.HomeTeam, score.ExtraTime.AwayTeam,
						score.Penalties.HomeTeam, score.Penalties.AwayTeam,
					)
				}
				for _, referee := range match.Referees {
					fmt.Printf("%8s %6d %s\n", "", referee.ID, referee.Name)
				}
			}
		}
	}
	return
}

func mytestFbdMatches(competitionID, matchday int) {
	comp := sepakbola.Competitions[competitionID]
	var matches fbd.MatchesData
	content := matches.Get(comp.ID, matchday)
	if err := matches.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "FBDMatches.Deserialize()", err)
		return
	}
	// result
	competition := matches.Competition
	teams := comp.Teams
	fmt.Printf("%s %s\n", "=== matches ===", comp.Name)
	fmt.Printf("(%d) [%-3s] %s\n", competition.ID, competition.Code, competition.Name)
	for _, match := range matches.Matches {
		fmt.Printf("(%6d) %2d %s %-15s %s %s\n", match.ID, match.Matchday, match.UtcDate, match.Status, teams[match.HomeTeam.ID].TLA, teams[match.AwayTeam.ID].TLA)
		if match.Score != nil {
			score := match.Score
			fmt.Printf("%8s %s:%s\n", "", "winner", score.Winner)
			fmt.Printf("%8s H(%d:%d) F(%d:%d) E(%d:%d) P(%d:%d)\n", "",
				score.HalfTime.HomeTeam, score.HalfTime.AwayTeam,
				score.FullTime.HomeTeam, score.FullTime.AwayTeam,
				score.ExtraTime.HomeTeam, score.ExtraTime.AwayTeam,
				score.Penalties.HomeTeam, score.Penalties.AwayTeam,
			)
		}
		for _, referee := range match.Referees {
			fmt.Printf("%8s %6d %s\n", "", referee.ID, referee.Name)
		}
	}
}

func mytestFbdTeams() {
	for _, comp := range sepakbola.Competitions {
		var fbdteams fbd.TeamsData
		content := fbdteams.Get(comp.ID)
		if err := fbdteams.Deserialize(content); err != nil {
			fmt.Printf("[ERROR] %s (%s)\n", "FBDTeams.Deserialize()", err)
			return
		}
		// result
		fmt.Printf("%s %s\n", "=== teams ===", comp.Name)
		for _, team := range fbdteams.Teams {
			fmt.Printf("%4d | %s | %4d | %-15s | %-26s | %s\n", team.ID, team.TLA, team.Founded, team.ShortName, team.Name, team.Website)
			fmt.Printf("%4s %-36s | %-20s | %s\n", "", team.Email, team.Phone, team.Address)
			fmt.Printf("%4s %s\n", "", team.CrestURL)
		}
	}
}

func mytestFbdTeam(teamID int) {
	var fbdteam fbd.Team
	content := fbdteam.Get(teamID)
	if err := fbdteam.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "FBDTeams.Deserialize()", err)
		return
	}
	// result
	fmt.Printf("%s %s | %s\n", "=== teams ===", fbdteam.Name, fbdteam.Area.Name)
	for _, squad := range fbdteam.Squads {
		fmt.Printf("%5d | %2d | %s\n", squad.ID, squad.ShirtNumber, squad.Name)
		fmt.Printf("%5s | %s\n", "", squad.DateOfBirth)
		fmt.Printf("%5s | %s\n", "", squad.CountryOfBirth)
		fmt.Printf("%5s | %s\n", "", squad.Nationality)
		fmt.Printf("%5s | %s\n", "", squad.Position)
		fmt.Printf("%5s | %s\n", "", squad.Role)
	}
	for _, competition := range fbdteam.Competitions {
		fmt.Printf("%5d | %4s | %s\n", competition.ID, competition.Code, competition.Name)
	}
}

func mytestFbdStandings(competitionID int) {
	comp := sepakbola.Competitions[competitionID]
	var standings fbd.StandingsData
	content := standings.Get(comp.ID)
	if err := standings.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "FBDStandings.Deserialize()", err)
		return
	}
	// result
	competition := standings.Competition
	season := standings.Season
	fmt.Printf("%s\n", "=== standings ===")
	fmt.Printf("(%d) [%-3s] %s", competition.ID, competition.Code, competition.Name)
	fmt.Printf("%6s (%s - %s) [%d]\n", "", season.StartDate, season.EndDate, season.CurrentMatchday)
	for _, standingsItem := range standings.StandingsList {
		if standingsItem.Type == "TOTAL" {
			for _, it := range standingsItem.Table {
				fmt.Printf("[%2d] %s %2d %2d %2d %2d %2d %2d %3d %3d\n",
					it.Position,
					comp.Teams[it.Team.ID].TLA,
					it.PlayedGames,
					it.Won,
					it.Draw,
					it.Lost,
					it.GoalsFor,
					it.GoalsAgainst,
					it.GoalDifference,
					it.Points)
			}
		}
	}
}

func mytestFbdUICompetitions() {
	contents := sepakbola.CompetitionsContents()
	contents.Dump()
}

func mytestFbdStandingsContents(competitionID int) {
	comp := sepakbola.Competitions[competitionID]
	contents := sepakbola.StandingsContents(comp)
	contents.Dump()
}

func mytestFbdTeamsContents(competitionID int) {
	comp := sepakbola.Competitions[competitionID]
	contents := sepakbola.TeamsContents(comp)
	contents.Dump()
}

func mytestFbdTeamContents(competitionID, teamID int) {
	comp := sepakbola.Competitions[competitionID]
	contents := sepakbola.TeamContents(comp, teamID)
	contents.Dump()
}

func mytestFbdSquadContents(competitionID, teamID, squadID int) {
	comp := sepakbola.Competitions[competitionID]
	for squadID != -1 {
		contents := sepakbola.SquadContents(comp, teamID, &squadID)
		contents.Dump()
	}
}

func mytestFbdMatchdayContents(competitionID int) {
	comp := sepakbola.Competitions[competitionID]
	contents := sepakbola.MatchdayContents(comp, 21)
	contents.Dump()
}

func mytestFbdMatchCalendarContents(competitionID int) {
	comp := sepakbola.Competitions[competitionID]
	contents := sepakbola.MatchCalendarContents(comp)
	contents.Dump()
}

func mytestFbdWarningContents(competitionID int) {
	uri := ui.URISomethingWrong800x329
	ratio := linebot.FlexImageAspectRatioType16to9
	somethingwrong := ui.WarningContents(ui.TextSomethingWrong, uri, ratio)
	somethingwrong.Dump()
}

func mytest() {
	competitionID := gsb.PL
	matchday := 16
	teamID := 57
	squadID := 0
	if os.Getenv("FBDATA_MATCHDAY") != "" {
		matchday, _ = strconv.Atoi(os.Getenv("FBDATA_MATCHDAY"))
	}
	if os.Getenv("MYTEST_FBD_COMPETITIONS") == "y" {
		mytestFbdCompetitions()
	}
	if os.Getenv("MYTEST_FBD_TEAMS") == "y" {
		mytestFbdTeams()
	}
	if os.Getenv("MYTEST_FBD_TEAM") == "y" {
		mytestFbdTeam(teamID)
	}
	if os.Getenv("MYTEST_FBD_MATCHES") == "y" {
		mytestFbdMatches(competitionID, 0)
	}
	if os.Getenv("MYTEST_FBD_MATCHDAY") == "y" {
		mytestFbdMatches(competitionID, matchday)
	}
	if os.Getenv("MYTEST_FBD_STANDINGS") == "y" {
		mytestFbdStandings(competitionID)
	}
	if os.Getenv("MYTEST_FBD_UI_COMPETITIONS") == "y" {
		mytestFbdUICompetitions()
	}
	if os.Getenv("MYTEST_FBD_TEAMS_CONTENTS") == "y" {
		mytestFbdTeamsContents(competitionID)
	}
	if os.Getenv("MYTEST_FBD_TEAM_CONTENTS") == "y" {
		mytestFbdTeamContents(competitionID, teamID)
	}
	if os.Getenv("MYTEST_FBD_SQUAD_CONTENTS") == "y" {
		mytestFbdSquadContents(competitionID, teamID, squadID)
	}
	if os.Getenv("MYTEST_FBD_STANDINGS_CONTENTS") == "y" {
		mytestFbdStandingsContents(competitionID)
	}
	if os.Getenv("MYTEST_FBD_MATCHDAY_CONTENTS") == "y" {
		mytestFbdMatchdayContents(competitionID)
	}
	if os.Getenv("MYTEST_FBD_MATCH_CALENDAR_CONTENTS") == "y" {
		mytestFbdMatchCalendarContents(competitionID)
	}
	os.Exit(0)
}

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
	"encoding/json"
	"fmt"
	"go-sepak-bola/internal/appdata"
	"go-sepak-bola/internal/fbd"
	"go-sepak-bola/ui"
	"regexp"
	"sort"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

// MatchdesTypeMessage function reply quick reply button message.
func (sepakbola *SepakBola) MatchdesTypeMessage(competition *fbd.Competition, matchday int) linebot.SendingMessage {
	allMatchesData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionMatchCalendar,
		Params:   nil,
	})
	matchdayData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionMatchday,
		Params: map[string]interface{}{
			"id":       competition.ID,
			"matchday": competition.Season.CurrentMatchday,
		},
	})

	return linebot.NewTextMessage(competition.Name).
		WithQuickReplies(linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton(
				"",
				linebot.NewPostbackAction(TextMatchCalendar, string(allMatchesData), "", TextMatchCalendar)),
			linebot.NewQuickReplyButton(
				"",
				linebot.NewPostbackAction(TextCurrentMatchday, string(matchdayData), "", TextCurrentMatchday)),
		))
}

// Matchdays defines map of matches for each matchday.
type Matchdays map[int][]*fbd.Match

// MatchCalendar defines map of MatchdayMatches for each month.
type MatchCalendar map[string]Matchdays

func (am MatchCalendar) getMatchesByMonth(matches []*fbd.Match) {
	for _, match := range matches {
		matchday := match.Matchday
		found := false
		for k, matchdays := range am {
			if _, ok := matchdays[matchday]; ok {
				am[k][matchday] = append(am[k][matchday], match)
				found = true
				break
			}
		}
		if !found {
			ex, _ := regexp.Compile("^\\d{4}-\\d{2}")
			yyyymm := ex.FindString(match.UtcDate)
			if _, ok := am[yyyymm]; !ok {
				am[yyyymm] = make(map[int][]*fbd.Match)
			}
			am[yyyymm][matchday] = append(am[yyyymm][matchday], match)
		}
	}
}

// BubbleMatchCalendar contains competition and matches information of month
// for creating a bubble container to list matchdays per month.
type BubbleMatchCalendar struct {
	Competition *fbd.Competition
	Matchdays   Matchdays
	Month       string
}

// Type is FlexContainerTypeBubble.
func (bm *BubbleMatchCalendar) Type() linebot.FlexContainerType {
	return linebot.FlexContainerTypeBubble
}

// Direction is text directionality and the order of components in horizontal
// boxes in the container. Specify one of the following values:
// 	ltr: Left to right
//	rtl: Right to left
func (bm *BubbleMatchCalendar) Direction() linebot.FlexBubbleDirectionType {
	return linebot.FlexBubbleDirectionTypeLTR
}

// Header block. Specify a box component.
func (bm *BubbleMatchCalendar) Header() *ui.ExtBoxComponent {
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: bm.Month,
				},
			},
		},
	}
}

// Hero block. Specify an image component.
func (bm *BubbleMatchCalendar) Hero() *linebot.ImageComponent {
	return nil
}

// Body block. Specify a box component.
func (bm *BubbleMatchCalendar) Body() *ui.ExtBoxComponent {
	rounds := make([]int, 0, len(bm.Matchdays))
	for round := range bm.Matchdays {
		rounds = append(rounds, round)
	}
	sort.Ints(rounds)
	buttons := make([]linebot.FlexComponent, 0, len(rounds))
	flexMatchday := 2
	for _, round := range rounds {
		matchdayData, _ := json.Marshal(&appdata.PostData{
			Category: PkgName,
			Action:   ActionMatchday,
			Params: map[string]interface{}{
				"id":       bm.Competition.ID,
				"matchday": strconv.Itoa(round),
			},
		})
		ex, _ := regexp.Compile("^\\d{4}-\\d{2}-\\d{2}")
		date := ex.FindString(bm.Matchdays[round][0].UtcDate)
		buttons = append(buttons, &linebot.ButtonComponent{
			Type:   linebot.FlexComponentTypeButton,
			Style:  linebot.FlexButtonStyleTypePrimary,
			Margin: linebot.FlexComponentMarginTypeXs,
			Color:  ColorOrange,
			Flex:   &flexMatchday,
			Action: &linebot.PostbackAction{
				Label:       strconv.Itoa(round) + " - " + date,
				Data:        string(matchdayData),
				Text:        "",
				DisplayText: "",
			},
		})
	}
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: buttons,
		},
	}
}

// Footer block. Specify a box component.
func (bm *BubbleMatchCalendar) Footer() *ui.ExtBoxComponent {
	return nil
}

// Styles of each block. Specify a bubble style object.
func (bm *BubbleMatchCalendar) Styles() *linebot.BubbleStyle {
	return nil
}

// MatchCalendarContents function generates CarouselContainer for matchday selection.
func (sepakbola *SepakBola) MatchCalendarContents(competition *fbd.Competition) *ui.ExtCarouselContainer {
	var bubbles []*ui.ExtBubbleContainer
	var matches fbd.MatchesData
	content := matches.Get(competition.ID, 0)
	if err := matches.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "MatchesData.Deserialize()", err)
		return ui.SomethingWrongContents()
	}
	calendar := MatchCalendar{}
	calendar.getMatchesByMonth(matches.Matches)
	months := make([]string, 0, len(calendar))
	for month := range calendar {
		months = append(months, month)
	}
	sort.Strings(months)
	for _, month := range months {
		bubble := BubbleMatchCalendar{
			Competition: competition,
			Matchdays:   calendar[month],
			Month:       month,
		}
		bubbles = append(bubbles, ui.Bubble(&bubble))
	}
	contents := ui.ExtCarouselContainer{
		CarouselContainer: linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
		},
		Contents: bubbles,
	}
	return &contents
}

// MatchCalendarMessage function generates FlexMessage for matchday selection.
func (sepakbola *SepakBola) MatchCalendarMessage(competition *fbd.Competition) *linebot.FlexMessage {
	altText := TextMatchCalendar
	fmt.Println("MatchCalendarMessage")
	contents := sepakbola.MatchCalendarContents(competition)
	fmt.Println("+++")
	contents.Dump()
	fmt.Println("---")
	return linebot.NewFlexMessage(altText, contents)
}

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
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

func (bs *BubbleMatchday) matchStatus() linebot.FlexComponent {
	statusComponent := &linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Text:  bs.Match.Status,
		Align: linebot.FlexComponentAlignTypeCenter,
		Size:  linebot.FlexTextSizeTypeSm,
	}
	return statusComponent
}

// BubbleMatchday contains competition and teams information for creating a bubble
// container of teams of the particular competition.
type BubbleMatchday struct {
	Competition *fbd.Competition
	Match       *fbd.Match
}

// Type is FlexContainerTypeBubble.
func (bs *BubbleMatchday) Type() linebot.FlexContainerType {
	return linebot.FlexContainerTypeBubble
}

// Direction is text directionality and the order of components in horizontal
// boxes in the container. Specify one of the following values:
// 	ltr: Left to right
//	rtl: Right to left
func (bs *BubbleMatchday) Direction() linebot.FlexBubbleDirectionType {
	return linebot.FlexBubbleDirectionTypeLTR
}

// Header block. Specify a box component.
func (bs *BubbleMatchday) Header() *ui.ExtBoxComponent {
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  bs.Competition.Name,
					Color: ColorAmber,
				},
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  TextRound + " " + strconv.Itoa(bs.Match.Season.CurrentMatchday),
					Align: linebot.FlexComponentAlignTypeEnd,
					Color: ColorAmber,
				},
			},
		},
	}
}

// Hero block. Specify an image component.
func (bs *BubbleMatchday) Hero() *linebot.ImageComponent {
	return nil
}

// Body block. Specify a box component.
func (bs *BubbleMatchday) Body() *ui.ExtBoxComponent {
	flexHomeTeam := 1
	flexAwayTeam := 1
	flexStatus := 2
	matchTime := toLocalTime(bs.Match.UtcDate)
	bodyContents := []linebot.FlexComponent{}
	bodyContents = append(bodyContents, &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeVertical,
					Flex:   &flexHomeTeam,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type: linebot.FlexComponentTypeText,
							Text: bs.Competition.Teams[bs.Match.HomeTeam.ID].TLA,
						},
					},
				},
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeVertical,
					Flex:   &flexStatus,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type:  linebot.FlexComponentTypeText,
							Text:  matchTime.Format(dateFormat),
							Align: linebot.FlexComponentAlignTypeCenter,
							Size:  linebot.FlexTextSizeTypeXs,
						},
						&linebot.TextComponent{
							Type:  linebot.FlexComponentTypeText,
							Text:  matchTime.Format(timeFormat),
							Align: linebot.FlexComponentAlignTypeCenter,
							Size:  linebot.FlexTextSizeTypeXs,
						},
						bs.matchStatus(),
					},
				},
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeVertical,
					Flex:   &flexAwayTeam,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type:  linebot.FlexComponentTypeText,
							Text:  bs.Competition.Teams[bs.Match.AwayTeam.ID].TLA,
							Align: linebot.FlexComponentAlignTypeEnd,
						},
					},
				},
			},
		},
	})
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: bodyContents,
		},
	}
}

// Footer block. Specify a box component.
func (bs *BubbleMatchday) Footer() *ui.ExtBoxComponent {
	return nil
}

// Styles of each block. Specify a bubble style object.
func (bs *BubbleMatchday) Styles() *linebot.BubbleStyle {
	return nil
}

// MatchdayContents function generates CarouselContainer for particular matchday.
func (sepakbola *SepakBola) MatchdayContents(competition *fbd.Competition, matchday int) *ui.ExtCarouselContainer {
	var bubbles []*ui.ExtBubbleContainer
	var matches fbd.MatchesData
	content := matches.Get(competition.ID, matchday)
	if err := matches.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "MatchesData.Deserialize()", err)
		return ui.SomethingWrongContents()
	}
	for _, match := range matches.Matches {
		bubble := BubbleMatchday{
			Competition: competition,
			Match:       match,
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

// MatchdayMessage function generates FlexMessage for particular matchday.
func (sepakbola *SepakBola) MatchdayMessage(competition *fbd.Competition, matchday int) *linebot.FlexMessage {
	altText := TextMatchday + " " + strconv.Itoa(matchday)
	contents := sepakbola.MatchdayContents(competition, matchday)
	return linebot.NewFlexMessage(altText, contents)
}

// MatchdesTypeMessage function reply quick reply button message.
func (sepakbola *SepakBola) MatchdesTypeMessage(competition *fbd.Competition, matchday int) linebot.SendingMessage {
	allMatchesData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionAllMatches,
		Params:   nil,
	})
	matchdayData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionMatchday,
		Params: map[string]interface{}{
			"id":              competition.ID,
			"currentMatchday": competition.Season.CurrentMatchday,
		},
	})

	return linebot.NewTextMessage(TextMatchesType).
		WithQuickReplies(linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton(
				"",
				linebot.NewPostbackAction(TextAllMatches, string(allMatchesData), "", TextAllMatches)),
			linebot.NewQuickReplyButton(
				"",
				linebot.NewPostbackAction(TextCurrentMatchday, string(matchdayData), "", TextCurrentMatchday)),
		))
}
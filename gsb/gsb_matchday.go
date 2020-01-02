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
	"go-sepak-bola/ui"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// BubbleMatchday contains competition and teams information for creating a bubble
// container of teams of the particular competition.
type BubbleMatchday struct {
	Competition *fbd.Competition
	Match       *fbd.Match
}

func (bm *BubbleMatchday) matchProgress() string {
	progress := "TBD"
	theTime, err := time.Parse(datetimeFormat, bm.Match.UtcDate)
	if err == nil {
		now := time.Now().UTC()
		diff := int(now.Sub(theTime).Minutes())
		if diff > 45 && diff < 60 {
			diff = 45
		} else if diff > 60 {
			diff -= 15
		}
		if diff > 90 {
			diff = 90
		}
		progress = fmt.Sprintf("%d", diff)
	}
	return progress
}

func (bm *BubbleMatchday) matchScore() *linebot.BoxComponent {
	scoreBox := linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Text:  strconv.Itoa(bm.Match.Score.FullTime.HomeTeam),
				Align: linebot.FlexComponentAlignTypeEnd,
				Size:  linebot.FlexTextSizeTypeMd,
			},
			&linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Text:  ":",
				Align: linebot.FlexComponentAlignTypeCenter,
				Size:  linebot.FlexTextSizeTypeMd,
			},
			&linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Text:  strconv.Itoa(bm.Match.Score.FullTime.AwayTeam),
				Align: linebot.FlexComponentAlignTypeStart,
				Size:  linebot.FlexTextSizeTypeMd,
			},
		},
	}
	var progress string
	if bm.Match.Status == StatusInPlay {
		progress = bm.matchProgress()
	} else if bm.Match.Status == StatusPaused {
		progress = "HT"
	} else if bm.Match.Status == StatusFinished {
		progress = "FT"
	}
	boxProgress := linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Text:  progress,
		Align: linebot.FlexComponentAlignTypeCenter,
		Size:  linebot.FlexTextSizeTypeMd,
	}
	return &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			&scoreBox,
			&boxProgress,
		},
	}
}

func (bm *BubbleMatchday) matchStatus() *linebot.FlexComponent {
	var status linebot.FlexComponent
	if bm.Match.Status == StatusInPlay ||
		bm.Match.Status == StatusPaused ||
		bm.Match.Status == StatusFinished {
		status = bm.matchScore()
	} else if bm.Match.Status == StatusScheduled {
		// no further information
	} else {
		status = &linebot.TextComponent{
			Type:  linebot.FlexComponentTypeText,
			Text:  bm.Match.Status,
			Align: linebot.FlexComponentAlignTypeCenter,
			Size:  linebot.FlexTextSizeTypeSm,
		}
	}
	return &status
}

func (bm *BubbleMatchday) matchReferees() *linebot.FlexComponent {
	var boxReferees linebot.FlexComponent
	if len(bm.Match.Referees) > 0 {
		var referees []linebot.FlexComponent
		referees = append(referees, &linebot.TextComponent{
			Type:  linebot.FlexComponentTypeText,
			Text:  "Referees",
			Color: ColorGray,
			Align: linebot.FlexComponentAlignTypeCenter,
		})
		for _, referee := range bm.Match.Referees {
			refereeName := referee.Name
			if len(refereeName) == 0 {
				refereeName = "-"
			}
			referees = append(referees, &linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Text:  refereeName,
				Color: ColorAmber,
				Align: linebot.FlexComponentAlignTypeCenter,
			})
		}
		boxReferees = &linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Margin:   linebot.FlexComponentMarginTypeMd,
			Contents: referees,
		}
	}
	return &boxReferees
}

func (bm *BubbleMatchday) matchInfo() []linebot.FlexComponent {
	matchTime := toLocalTime(bm.Match.UtcDate)
	contents := []linebot.FlexComponent{
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
	}
	matchStatus := bm.matchStatus()
	if *matchStatus != nil {
		contents = append(contents, *matchStatus)
	}
	return contents
}

// Type is FlexContainerTypeBubble.
func (bm *BubbleMatchday) Type() linebot.FlexContainerType {
	return linebot.FlexContainerTypeBubble
}

// Direction is text directionality and the order of components in horizontal
// boxes in the container. Specify one of the following values:
// 	ltr: Left to right
//	rtl: Right to left
func (bm *BubbleMatchday) Direction() linebot.FlexBubbleDirectionType {
	return linebot.FlexBubbleDirectionTypeLTR
}

// Header block. Specify a box component.
func (bm *BubbleMatchday) Header() *ui.ExtBoxComponent {
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  bm.Competition.Name,
					Color: ColorAmber,
				},
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  TextRound + " " + strconv.Itoa(bm.Match.Matchday),
					Align: linebot.FlexComponentAlignTypeEnd,
					Color: ColorAmber,
				},
			},
		},
	}
}

// Hero block. Specify an image component.
func (bm *BubbleMatchday) Hero() *linebot.ImageComponent {
	return nil
}

// Body block. Specify a box component.
func (bm *BubbleMatchday) Body() *ui.ExtBoxComponent {
	flexStatus := 2
	bodyContents := []linebot.FlexComponent{}
	bodyContents = append(bodyContents, &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Text:    bm.Competition.Teams[bm.Match.HomeTeam.ID].TLA,
					Align:   linebot.FlexComponentAlignTypeEnd,
					Gravity: linebot.FlexComponentGravityTypeCenter,
				},
				&linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Flex:     &flexStatus,
					Contents: bm.matchInfo(),
				},
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Text:    bm.Competition.Teams[bm.Match.AwayTeam.ID].TLA,
					Align:   linebot.FlexComponentAlignTypeStart,
					Gravity: linebot.FlexComponentGravityTypeCenter,
				},
			},
		},
	},
	)
	matchReferees := bm.matchReferees()
	if *matchReferees != nil {
		bodyContents = append(bodyContents, &linebot.SeparatorComponent{
			Type:   linebot.FlexComponentTypeSeparator,
			Margin: linebot.FlexComponentMarginTypeMd,
			Color:  ColorGreenYellow,
		})
		bodyContents = append(bodyContents, *matchReferees)
	}
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: bodyContents,
		},
	}
}

// Footer block. Specify a box component.
func (bm *BubbleMatchday) Footer() *ui.ExtBoxComponent {
	return nil
}

// Styles of each block. Specify a bubble style object.
func (bm *BubbleMatchday) Styles() *linebot.BubbleStyle {
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

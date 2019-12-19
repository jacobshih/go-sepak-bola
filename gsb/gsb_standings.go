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

	"github.com/line/line-bot-sdk-go/linebot"
)

// BubbleStandings contains competition information for creating a bubble
// container of particular competition.
type BubbleStandings struct {
	Competition *fbd.Competition
	Standings   *fbd.Standings
}

// Type is FlexContainerTypeBubble.
func (bs *BubbleStandings) Type() linebot.FlexContainerType {
	return linebot.FlexContainerTypeBubble
}

// Direction is text directionality and the order of components in horizontal
// boxes in the container. Specify one of the following values:
// 	ltr: Left to right
//	rtl: Right to left
func (bs *BubbleStandings) Direction() linebot.FlexBubbleDirectionType {
	return linebot.FlexBubbleDirectionTypeLTR
}

// Header block. Specify a box component.
func (bs *BubbleStandings) Header() *ui.ExtBoxComponent {
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   bs.Competition.Name,
					Margin: linebot.FlexComponentMarginTypeNone,
					Size:   linebot.FlexTextSizeTypeLg,
					Color:  ColorAero,
				},
			},
		},
	}
}

// Hero block. Specify an image component.
func (bs *BubbleStandings) Hero() *linebot.ImageComponent {
	return &linebot.ImageComponent{
		Type:        linebot.FlexComponentTypeImage,
		URL:         bs.Competition.EmblemURL,
		Size:        linebot.FlexImageSizeTypeFull,
		AspectRatio: linebot.FlexImageAspectRatioType1to1,
		AspectMode:  linebot.FlexImageAspectModeTypeCover,
	}
}

func standingsTextCell(text string, flex int, textColor, backgroundColor string) *ui.ExtBoxComponent {
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Text:    text,
					Flex:    &flex,
					Gravity: linebot.FlexComponentGravityTypeCenter,
					Margin:  linebot.FlexComponentMarginTypeNone,
					Size:    linebot.FlexTextSizeTypeXs,
					Color:   textColor,
				},
			},
		},
		BackgroundColor: backgroundColor,
	}
}

func standingsImageCell(imageURL string, flex int, imageSize linebot.FlexImageSizeType) *ui.ExtBoxComponent {
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ImageComponent{
					Type:        linebot.FlexComponentTypeImage,
					URL:         imageURL,
					Size:        imageSize,
					AspectRatio: linebot.FlexImageAspectRatioType1to1,
					AspectMode:  linebot.FlexImageAspectModeTypeCover,
				},
			},
		},
	}
}

// Body block. Specify a box component.
func (bs *BubbleStandings) Body() *ui.ExtBoxComponent {
	flexCrest := 14
	flexTLA := 14
	flexCell := 9
	thTextColor := ColorDodgeBlue
	thBackgroundColor := ColorLightGray
	comp := bs.Competition
	teams := comp.Teams
	bodyContents := []linebot.FlexComponent{}
	bodyContents = append(bodyContents, &linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Text:  bs.Standings.Type,
		Size:  linebot.FlexTextSizeTypeLg,
		Color: ColorGreenYellow,
	})
	bodyContents = append(bodyContents, &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				standingsTextCell(TextSpace, flexCrest, thTextColor, thBackgroundColor),
				standingsTextCell(TextSpace, flexTLA, thTextColor, thBackgroundColor),
				standingsTextCell("GP", flexCell, thTextColor, thBackgroundColor),
				standingsTextCell("W", flexCell, thTextColor, thBackgroundColor),
				standingsTextCell("D", flexCell, thTextColor, thBackgroundColor),
				standingsTextCell("L", flexCell, thTextColor, thBackgroundColor),
				standingsTextCell("PT", flexCell, thTextColor, thBackgroundColor),
				standingsTextCell("GF", flexCell, thTextColor, thBackgroundColor),
				standingsTextCell("GA", flexCell, thTextColor, thBackgroundColor),
				standingsTextCell("GD", flexCell, thTextColor, thBackgroundColor),
			},
		},
	})
	thTextColor = ColorBlack
	thBackgroundColor = ColorWhite
	for _, it := range bs.Standings.Table {
		bodyContents = append(bodyContents, &ui.ExtBoxComponent{
			BoxComponent: linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					standingsImageCell(teams[it.Team.ID].CrestURL, flexCrest, linebot.FlexImageSizeTypeXxs),
					standingsTextCell(teams[it.Team.ID].TLA, flexTLA, thTextColor, thBackgroundColor),
					standingsTextCell(strconv.Itoa(it.PlayedGames), flexCell, thTextColor, thBackgroundColor),
					standingsTextCell(strconv.Itoa(it.Won), flexCell, thTextColor, thBackgroundColor),
					standingsTextCell(strconv.Itoa(it.Draw), flexCell, thTextColor, thBackgroundColor),
					standingsTextCell(strconv.Itoa(it.Lost), flexCell, thTextColor, thBackgroundColor),
					standingsTextCell(strconv.Itoa(it.Points), flexCell, thTextColor, thBackgroundColor),
					standingsTextCell(strconv.Itoa(it.GoalsFor), flexCell, thTextColor, thBackgroundColor),
					standingsTextCell(strconv.Itoa(it.GoalsAgainst), flexCell, thTextColor, thBackgroundColor),
					standingsTextCell(strconv.Itoa(it.GoalDifference), flexCell, thTextColor, thBackgroundColor),
				},
			},
		})
		break
	}
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Spacing:  linebot.FlexComponentSpacingTypeMd,
			Contents: bodyContents,
		},
	}
}

// Footer block. Specify a box component.
func (bs *BubbleStandings) Footer() *ui.ExtBoxComponent {
	return nil
}

// Styles of each block. Specify a bubble style object.
func (bs *BubbleStandings) Styles() *linebot.BubbleStyle {
	return nil
}

// StandingsContents function generates CarouselContainer for standings.
func (sepakbola *SepakBola) StandingsContents(competition *fbd.Competition) *ui.ExtCarouselContainer {
	var bubbles []*ui.ExtBubbleContainer
	var standings fbd.StandingsData
	fmt.Println("[StandingsMessage] competitionID: ", competition.ID)
	content := standings.Get(competition.ID)
	if err := standings.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "FBDStandings.Deserialize()", err)
		// FIXME: return ui.SomethingWrongContents()
		return nil
	}
	for _, standingsItem := range standings.StandingsList {
		bubble := BubbleStandings{
			Competition: competition,
			Standings:   standingsItem,
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

// StandingsMessage function generates FlexMessage for standings.
func (sepakbola *SepakBola) StandingsMessage(competition *fbd.Competition) *linebot.FlexMessage {
	altText := "Standings"
	fmt.Println("[StandingsMessage] competitionID: ", competition.ID)
	contents := sepakbola.StandingsContents(competition)
	return linebot.NewFlexMessage(altText, contents)
}

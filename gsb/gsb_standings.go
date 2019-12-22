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

const (
	colorCLG = "#45ef00" // champions league group stage
	colorCLQ = "#00ef7f" // champions league qualification
	colorELG = "#00c7ef" // europa league group stage
	colorELQ = "#00efef" // europa league qualification
	colorPO  = "#ef8a00" // play offs
	colorREL = "#f96f5f" // relegation
)

// Legend type maps the position of standings table to a background color to
// indicate the status of promotion or relegation.
type Legend struct {
	table map[int]string
}

var legends = map[int]Legend{
	PL: {
		table: map[int]string{
			1:  colorCLG,
			2:  colorCLG,
			3:  colorCLG,
			4:  colorCLG,
			5:  colorELG,
			18: colorREL,
			19: colorREL,
			20: colorREL,
		},
	},
	BL1: {
		table: map[int]string{
			1:  colorCLG,
			2:  colorCLG,
			3:  colorCLG,
			4:  colorCLG,
			5:  colorELG,
			6:  colorELQ,
			16: colorPO,
			17: colorREL,
			18: colorREL,
		},
	},
	PD: {
		table: map[int]string{
			1:  colorCLG,
			2:  colorCLG,
			3:  colorCLG,
			4:  colorCLG,
			5:  colorELG,
			6:  colorELQ,
			18: colorREL,
			19: colorREL,
			20: colorREL,
		},
	},
}

// BubbleStandings contains competition and standings information for creating
// a bubble container of standings table of the particular competition.
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
		Type: linebot.FlexComponentTypeImage,
		URL:  bs.Competition.EmblemURL,
		Size: linebot.FlexImageSizeTypeFull,
	}
}

func standingsTextCell(text string, flex int, textColor, backgroundColor string, align linebot.FlexComponentAlignType) *ui.ExtBoxComponent {
	box := ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Flex:   &flex,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Text:    text,
					Gravity: linebot.FlexComponentGravityTypeCenter,
					Align:   linebot.FlexComponentAlignType(align),
					Size:    linebot.FlexTextSizeTypeXs,
					Color:   textColor,
				},
			},
		},
	}
	if len(backgroundColor) > 0 {
		box.BackgroundColor = backgroundColor
	}
	return &box
}

func standingsImageCell(imageURL string, flex int, imageSize linebot.FlexImageSizeType) *ui.ExtBoxComponent {
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Flex:   &flex,
			Contents: []linebot.FlexComponent{
				&linebot.ImageComponent{
					Type: linebot.FlexComponentTypeImage,
					URL:  imageURL,
					Size: imageSize,
				},
			},
		},
	}
}

// Body block. Specify a box component.
func (bs *BubbleStandings) Body() *ui.ExtBoxComponent {
	flexCrest := 14
	flexTLA := 14
	flexCell := 8
	alignL := linebot.FlexComponentAlignTypeStart
	alignR := linebot.FlexComponentAlignTypeEnd
	thColor := ColorIndigo
	noColor := ""
	comp := bs.Competition
	teams := comp.Teams
	legend := legends[comp.ID].table
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
			Margin: linebot.FlexComponentMarginTypeXs,
			Contents: []linebot.FlexComponent{
				standingsTextCell(TextSpace, flexCell, thColor, noColor, alignL),
				standingsTextCell(TextSpace, flexCrest, thColor, noColor, alignL),
				standingsTextCell(TextSpace, flexTLA, thColor, noColor, alignL),
				standingsTextCell("GP", flexCell, thColor, noColor, alignL),
				standingsTextCell("W", flexCell, thColor, noColor, alignL),
				standingsTextCell("D", flexCell, thColor, noColor, alignL),
				standingsTextCell("L", flexCell, thColor, noColor, alignL),
				standingsTextCell("PT", flexCell, thColor, noColor, alignL),
				standingsTextCell("GF", flexCell, thColor, noColor, alignL),
				standingsTextCell("GA", flexCell, thColor, noColor, alignL),
				standingsTextCell("GD", flexCell, thColor, noColor, alignL),
			},
		},
		BackgroundColor: ColorLightGray,
	})
	thColor = ColorBlack
	for _, it := range bs.Standings.Table {
		box := ui.ExtBoxComponent{
			BoxComponent: linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Margin: linebot.FlexComponentMarginTypeXs,
				Contents: []linebot.FlexComponent{
					standingsTextCell(strconv.Itoa(it.Position), flexCell, thColor, noColor, alignL),
					standingsImageCell(teams[it.Team.ID].CrestURL, flexCrest, linebot.FlexImageSizeTypeXxs),
					standingsTextCell(teams[it.Team.ID].TLA, flexTLA, thColor, noColor, alignL),
					standingsTextCell(strconv.Itoa(it.PlayedGames), flexCell, thColor, noColor, alignR),
					standingsTextCell(strconv.Itoa(it.Won), flexCell, thColor, noColor, alignR),
					standingsTextCell(strconv.Itoa(it.Draw), flexCell, thColor, noColor, alignR),
					standingsTextCell(strconv.Itoa(it.Lost), flexCell, thColor, noColor, alignR),
					standingsTextCell(strconv.Itoa(it.Points), flexCell, thColor, noColor, alignR),
					standingsTextCell(strconv.Itoa(it.GoalsFor), flexCell, thColor, noColor, alignR),
					standingsTextCell(strconv.Itoa(it.GoalsAgainst), flexCell, thColor, noColor, alignR),
					standingsTextCell(strconv.Itoa(it.GoalDifference), flexCell, thColor, noColor, alignR),
				},
			},
		}
		if color, ok := legend[it.Position]; ok {
			box.BackgroundColor = color
		}
		bodyContents = append(bodyContents, &box)
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
	content := standings.Get(competition.ID)
	if err := standings.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "FBDStandings.Deserialize()", err)
		return ui.SomethingWrongContents()
	}
	for _, standingsItem := range standings.StandingsList {
		bubble := BubbleStandings{
			Competition: competition,
			Standings:   standingsItem,
		}
		bubbles = append(bubbles, ui.Bubble(&bubble))
		// skip HOME and AWAY standings to reduce size of flex message.
		break
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
	contents := sepakbola.StandingsContents(competition)
	return linebot.NewFlexMessage(altText, contents)
}

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
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

// BubbleSquad contains competition and team information for creating a bubble
// container of team squads information.
type BubbleSquad struct {
	Competition *fbd.Competition
	Team        *fbd.Team
	Squad       *fbd.Squad
}

// Type is FlexContainerTypeBubble.
func (bs *BubbleSquad) Type() linebot.FlexContainerType {
	return linebot.FlexContainerTypeBubble
}

// Direction is text directionality and the order of components in horizontal
// boxes in the container. Specify one of the following values:
// 	ltr: Left to right
//	rtl: Right to left
func (bs *BubbleSquad) Direction() linebot.FlexBubbleDirectionType {
	return linebot.FlexBubbleDirectionTypeLTR
}

// Header block. Specify a box component.
func (bs *BubbleSquad) Header() *ui.ExtBoxComponent {
	flexName := 6
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Text:    bs.Squad.Name,
					Size:    linebot.FlexTextSizeTypeLg,
					Color:   ColorAmber,
					Gravity: linebot.FlexComponentGravityTypeCenter,
					Wrap:    true,
					Flex:    &flexName,
				},
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  strconv.Itoa(bs.Squad.ShirtNumber),
					Align: linebot.FlexComponentAlignTypeEnd,
					Size:  linebot.FlexTextSizeTypeXxl,
					Color: ColorAmber,
				},
			},
		},
	}
}

// Hero block. Specify an image component.
func (bs *BubbleSquad) Hero() *linebot.ImageComponent {
	return nil
}

// Body block. Specify a box component.
func (bs *BubbleSquad) Body() *ui.ExtBoxComponent {
	const (
		TextSquadBirthday    = "Birthday"
		TextSquadBirthplace  = "Birthplace"
		TextSquadNationality = "Nationality"
		TextSquadPosition    = "Position"
	)
	ex, _ := regexp.Compile("^\\d{4}-\\d{2}-\\d{2}")
	birthday := ex.FindString(bs.Squad.DateOfBirth)
	bodyContents := []linebot.FlexComponent{}
	bodyContents = append(bodyContents, bs.squadInfo(TextSquadBirthday, birthday))
	bodyContents = append(bodyContents, bs.squadInfo(TextSquadBirthplace, bs.Squad.CountryOfBirth))
	bodyContents = append(bodyContents, bs.squadInfo(TextSquadNationality, bs.Squad.Nationality))
	bodyContents = append(bodyContents, bs.squadInfo(TextSquadPosition, bs.Squad.Position))
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: bodyContents,
		},
	}
}

// Footer block. Specify a box component.
func (bs *BubbleSquad) Footer() *ui.ExtBoxComponent {
	return nil
}

// Styles of each block. Specify a bubble style object.
func (bs *BubbleSquad) Styles() *linebot.BubbleStyle {
	return &linebot.BubbleStyle{
		Body: &linebot.BlockStyle{
			Separator:      true,
			SeparatorColor: ColorAmber,
		},
	}
}

func (bs *BubbleSquad) squadInfo(label, value string) *ui.ExtBoxComponent {
	flexValue := 2
	labelComponent := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   label,
		Margin: linebot.FlexComponentMarginTypeMd,
		Size:   linebot.FlexTextSizeTypeSm,
	}
	valueComponent := linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   value,
		Margin: linebot.FlexComponentMarginTypeMd,
		Size:   linebot.FlexTextSizeTypeSm,
		Flex:   &flexValue,
	}
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&labelComponent,
				&linebot.SeparatorComponent{
					Type:  linebot.FlexComponentTypeSeparator,
					Color: ColorAmber,
				},
				&valueComponent,
			},
		},
	}
}

// SquadContents function generates CarouselContainer for team squads.
func (sepakbola *SepakBola) SquadContents(competition *fbd.Competition, teamID int, squadID *int) *ui.ExtCarouselContainer {
	if _, ok := sepakbola.Competitions[competition.ID].Teams[teamID]; !ok {
		return ui.SomethingWrongContents()
	}
	var nextSquadID = -1
	var team fbd.Team
	content := team.Get(teamID)
	if err := team.Deserialize(content); err != nil {
		fmt.Printf("[ERROR] %s (%s)\n", "Team.Deserialize()", err)
		return ui.SomethingWrongContents()
	}
	var bubbles []*ui.ExtBubbleContainer
	i := 0
	for i = 0; i < len(team.Squads); i++ {
		if *squadID == 0 || *squadID == team.Squads[i].ID {
			break
		}
	}
	for ; i < len(team.Squads); i++ {
		squad := team.Squads[i]
		if len(bubbles) >= 10 {
			nextSquadID = squad.ID
			break
		}
		bubble := BubbleSquad{
			Competition: competition,
			Team:        &team,
			Squad:       squad,
		}
		bubbles = append(bubbles, ui.Bubble(&bubble))
	}
	*squadID = nextSquadID
	contents := ui.ExtCarouselContainer{
		CarouselContainer: linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
		},
		Contents: bubbles,
	}
	return &contents
}

// SquadMessage function generates FlexMessage for team squads.
func (sepakbola *SepakBola) SquadMessage(competition *fbd.Competition, teamID, squadID int) linebot.SendingMessage {
	altText := competition.Teams[teamID].Name + " " + TextSquad
	contents := sepakbola.SquadContents(competition, teamID, &squadID)
	msg := linebot.NewFlexMessage(altText, contents)
	if squadID != -1 {
		squadData, _ := json.Marshal(&appdata.PostData{
			Category: PkgName,
			Action:   ActionSquad,
			Params: map[string]interface{}{
				"id":      competition.ID,
				"teamID":  teamID,
				"squadID": squadID,
			},
		})
		msg = msg.WithQuickReplies(linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton(
				"",
				linebot.NewPostbackAction(TextMore, string(squadData), "", "")),
		)).(*linebot.FlexMessage)
	}
	return msg
}

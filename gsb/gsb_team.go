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
	"go-sepak-bola/internal/appdata"
	"go-sepak-bola/internal/fbd"
	"go-sepak-bola/ui"
	"regexp"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

// BubbleTeam contains competition and team information for creating a bubble
// container of team information.
type BubbleTeam struct {
	Competition *fbd.Competition
	Team        *fbd.Team
}

// Type is FlexContainerTypeBubble.
func (bt *BubbleTeam) Type() linebot.FlexContainerType {
	return linebot.FlexContainerTypeBubble
}

// Direction is text directionality and the order of components in horizontal
// boxes in the container. Specify one of the following values:
// 	ltr: Left to right
//	rtl: Right to left
func (bt *BubbleTeam) Direction() linebot.FlexBubbleDirectionType {
	return linebot.FlexBubbleDirectionTypeLTR
}

// Header block. Specify a box component.
func (bt *BubbleTeam) Header() *ui.ExtBoxComponent {
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   bt.Team.Name,
					Margin: linebot.FlexComponentMarginTypeNone,
					Size:   linebot.FlexTextSizeTypeLg,
					Color:  ColorAmber,
				},
			},
		},
	}
}

// Hero block. Specify an image component.
func (bt *BubbleTeam) Hero() *linebot.ImageComponent {
	return nil
}

func (bt *BubbleTeam) teamInfo(label, value string) *ui.ExtBoxComponent {
	flexValue := 3
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
		Wrap:   true,
		Flex:   &flexValue,
	}
	ex, _ := regexp.Compile("^http(s)?://\\w+(\\.\\w+)+$")
	if ex.Match([]byte(value)) {
		valueComponent.Action = linebot.NewURIAction(bt.Team.Name, bt.Team.Website)
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

// Body block. Specify a box component.
func (bt *BubbleTeam) Body() *ui.ExtBoxComponent {
	const (
		TextTeamFounded    = "Founded"
		TextTeamVenue      = "Venue"
		TextTeamCountry    = "Country"
		TextTeamClubColors = "Colors"
		TextTeamWebsite    = "Website"
		TextTeamAddress    = "Address"
	)
	bodyContents := []linebot.FlexComponent{}
	bodyContents = append(bodyContents, bt.teamInfo(TextTeamFounded, strconv.Itoa(bt.Team.Founded)))
	bodyContents = append(bodyContents, bt.teamInfo(TextTeamVenue, bt.Team.Venue))
	bodyContents = append(bodyContents, bt.teamInfo(TextTeamCountry, bt.Team.Area.Name))
	bodyContents = append(bodyContents, bt.teamInfo(TextTeamClubColors, bt.Team.ClubColors))
	bodyContents = append(bodyContents, bt.teamInfo(TextTeamAddress, bt.Team.Address))
	bodyContents = append(bodyContents, bt.teamInfo(TextTeamWebsite, bt.Team.Website))
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: bodyContents,
		},
	}
}

// Footer block. Specify a box component.
func (bt *BubbleTeam) Footer() *ui.ExtBoxComponent {
	return nil
}

// Styles of each block. Specify a bubble style object.
func (bt *BubbleTeam) Styles() *linebot.BubbleStyle {
	return &linebot.BubbleStyle{
		Body: &linebot.BlockStyle{
			Separator:      true,
			SeparatorColor: ColorAmber,
		},
	}
}

// TeamContents function generates CarouselContainer for particular team.
func (sepakbola *SepakBola) TeamContents(competition *fbd.Competition, teamID int) *ui.ExtCarouselContainer {
	if _, ok := sepakbola.Competitions[competition.ID].Teams[teamID]; !ok {
		return ui.SomethingWrongContents()
	}
	var bubbles []*ui.ExtBubbleContainer
	bubble := BubbleTeam{
		Competition: competition,
		Team:        sepakbola.Competitions[competition.ID].Teams[teamID],
	}
	bubbles = append(bubbles, ui.Bubble(&bubble))
	contents := ui.ExtCarouselContainer{
		CarouselContainer: linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
		},
		Contents: bubbles,
	}
	return &contents
}

// TeamMessage function generates FlexMessage for particular team.
func (sepakbola *SepakBola) TeamMessage(competition *fbd.Competition, teamID int) linebot.SendingMessage {
	squadData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionSquad,
		Params: map[string]interface{}{
			"id":      competition.ID,
			"teamID":  teamID,
			"squadID": 0,
		},
	})
	altText := competition.Teams[teamID].Name
	contents := sepakbola.TeamContents(competition, teamID)
	return linebot.NewFlexMessage(altText, contents).WithQuickReplies(linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton(
			"",
			linebot.NewPostbackAction(TextSquad, string(squadData), "", "")),
	))
}

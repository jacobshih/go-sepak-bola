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
	"sort"

	"github.com/line/line-bot-sdk-go/linebot"
)

// BubbleTeams contains competition and teams information for creating a bubble
// container of teams of the particular competition.
type BubbleTeams struct {
	Competition *fbd.Competition
	Teams       map[int]*fbd.Team
}

// Type is FlexContainerTypeBubble.
func (bs *BubbleTeams) Type() linebot.FlexContainerType {
	return linebot.FlexContainerTypeBubble
}

// Direction is text directionality and the order of components in horizontal
// boxes in the container. Specify one of the following values:
// 	ltr: Left to right
//	rtl: Right to left
func (bs *BubbleTeams) Direction() linebot.FlexBubbleDirectionType {
	return linebot.FlexBubbleDirectionTypeLTR
}

// Header block. Specify a box component.
func (bs *BubbleTeams) Header() *ui.ExtBoxComponent {
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
func (bs *BubbleTeams) Hero() *linebot.ImageComponent {
	return &linebot.ImageComponent{
		Type: linebot.FlexComponentTypeImage,
		URL:  bs.Competition.EmblemURL,
		Size: linebot.FlexImageSizeTypeFull,
	}
}

// Body block. Specify a box component.
func (bs *BubbleTeams) Body() *ui.ExtBoxComponent {
	flexCrest := 1
	flexTeam := 4
	bodyContents := []linebot.FlexComponent{}
	keys := []int{}
	for id := range bs.Teams {
		keys = append(keys, id)
	}
	sort.Ints(keys)
	for id := range keys {
		team := bs.Teams[keys[id]]
		teamData, _ := json.Marshal(
			&appdata.PostData{
				Category: PkgName,
				Action:   ActionTeam,
				Params: map[string]interface{}{
					"id": team.ID,
				},
			})
		box := ui.ExtBoxComponent{
			BoxComponent: linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Margin: linebot.FlexComponentMarginTypeXs,
				Contents: []linebot.FlexComponent{
					&linebot.ImageComponent{
						Type: linebot.FlexComponentTypeImage,
						Flex: &flexCrest,
						URL:  team.CrestURL,
						Size: linebot.FlexImageSizeTypeXxs,
						Action: &linebot.PostbackAction{
							Data: string(teamData),
						},
					},
					&linebot.TextComponent{
						Type:    linebot.FlexComponentTypeText,
						Flex:    &flexTeam,
						Text:    team.ShortName,
						Margin:  linebot.FlexComponentMarginTypeSm,
						Size:    linebot.FlexTextSizeTypeXl,
						Color:   ColorDodgeBlue,
						Gravity: linebot.FlexComponentGravityTypeCenter,
						Action: &linebot.PostbackAction{
							Data: string(teamData),
						},
					},
				},
			},
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
func (bs *BubbleTeams) Footer() *ui.ExtBoxComponent {
	return nil
}

// Styles of each block. Specify a bubble style object.
func (bs *BubbleTeams) Styles() *linebot.BubbleStyle {
	return nil
}

// TeamsContents function generates CarouselContainer for teams.
func (sepakbola *SepakBola) TeamsContents(competition *fbd.Competition) *ui.ExtCarouselContainer {
	var bubbles []*ui.ExtBubbleContainer
	bubble := BubbleTeams{
		Competition: competition,
		Teams:       sepakbola.Competitions[competition.ID].Teams,
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

// TeamsMessage function generates FlexMessage for teams.
func (sepakbola *SepakBola) TeamsMessage(competition *fbd.Competition) *linebot.FlexMessage {
	altText := "Teams"
	contents := sepakbola.TeamsContents(competition)
	return linebot.NewFlexMessage(altText, contents)
}

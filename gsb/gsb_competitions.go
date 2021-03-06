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

	"github.com/line/line-bot-sdk-go/linebot"
)

// BubbleCompetition contains competition information for creating a bubble
// container of particular competition.
type BubbleCompetition struct {
	Competition *fbd.Competition
}

// Type is FlexContainerTypeBubble.
func (bc *BubbleCompetition) Type() linebot.FlexContainerType {
	return linebot.FlexContainerTypeBubble
}

// Direction is text directionality and the order of components in horizontal
// boxes in the container. Specify one of the following values:
// 	ltr: Left to right
//	rtl: Right to left
func (bc *BubbleCompetition) Direction() linebot.FlexBubbleDirectionType {
	return linebot.FlexBubbleDirectionTypeLTR
}

// Header block. Specify a box component.
func (bc *BubbleCompetition) Header() *ui.ExtBoxComponent {
	return &ui.ExtBoxComponent{
		BoxComponent: linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   bc.Competition.Name,
					Size:   linebot.FlexTextSizeTypeLg,
					Color:  ColorAmber,
					Margin: linebot.FlexComponentMarginTypeNone,
				},
			},
		},
	}
}

// Hero block. Specify an image component.
func (bc *BubbleCompetition) Hero() *linebot.ImageComponent {
	return &linebot.ImageComponent{
		Type:        linebot.FlexComponentTypeImage,
		URL:         bc.Competition.EmblemURL,
		Size:        linebot.FlexImageSizeType4xl,
		AspectRatio: linebot.FlexImageAspectRatioType1to1,
		AspectMode:  linebot.FlexImageAspectModeTypeCover,
	}
}

// Body block. Specify a box component.
func (bc *BubbleCompetition) Body() *ui.ExtBoxComponent {
	matchesData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionMatches,
		Params: map[string]interface{}{
			"id":        bc.Competition.ID,
			"code":      bc.Competition.Code,
			"name":      bc.Competition.Name,
			"emblemURL": bc.Competition.EmblemURL,
		},
	})
	standingsData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionStandings,
		Params: map[string]interface{}{
			"id":        bc.Competition.ID,
			"code":      bc.Competition.Code,
			"name":      bc.Competition.Name,
			"emblemURL": bc.Competition.EmblemURL,
		},
	})
	teamsData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionTeams,
		Params: map[string]interface{}{
			"id":        bc.Competition.ID,
			"code":      bc.Competition.Code,
			"name":      bc.Competition.Name,
			"emblemURL": bc.Competition.EmblemURL,
		},
	})

	bodyContents := []linebot.FlexComponent{}
	bodyContents = append(bodyContents, &linebot.ButtonComponent{
		Type:  linebot.FlexComponentTypeButton,
		Style: linebot.FlexButtonStyleTypeSecondary,
		Color: ColorOrange,
		Action: &linebot.PostbackAction{
			Label:       EmojiSquaredVS + " " + TextMatches,
			Data:        string(matchesData),
			Text:        "",
			DisplayText: "",
		},
	})
	bodyContents = append(bodyContents, &linebot.ButtonComponent{
		Type:  linebot.FlexComponentTypeButton,
		Style: linebot.FlexButtonStyleTypeSecondary,
		Color: ColorOrange,
		Action: &linebot.PostbackAction{
			Label:       EmojiTrophy + " " + TextStandings,
			Data:        string(standingsData),
			Text:        "",
			DisplayText: "",
		},
	})
	bodyContents = append(bodyContents, &linebot.ButtonComponent{
		Type:  linebot.FlexComponentTypeButton,
		Style: linebot.FlexButtonStyleTypeSecondary,
		Color: ColorOrange,
		Action: &linebot.PostbackAction{
			Label:       EmojiSoccer + " " + TextTeams,
			Data:        string(teamsData),
			Text:        "",
			DisplayText: "",
		},
	})
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
func (bc *BubbleCompetition) Footer() *ui.ExtBoxComponent {
	return nil
}

// Styles of each block. Specify a bubble style object.
func (bc *BubbleCompetition) Styles() *linebot.BubbleStyle {
	return &linebot.BubbleStyle{
		Body: &linebot.BlockStyle{
			Separator:      true,
			SeparatorColor: ColorAmber,
		},
	}
}

// CompetitionsContents function generates CarouselContainer for competitions menu.
func (sepakbola *SepakBola) CompetitionsContents() *ui.ExtCarouselContainer {
	var bubbles []*ui.ExtBubbleContainer
	for _, competition := range sepakbola.Competitions {
		bubble := BubbleCompetition{Competition: competition}
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

// CompetitionsMessage function generates FlexMessage for competitions menu.
func (sepakbola *SepakBola) CompetitionsMessage() *linebot.FlexMessage {
	altText := "Competitions"
	contents := sepakbola.CompetitionsContents()
	return linebot.NewFlexMessage(altText, contents)
}

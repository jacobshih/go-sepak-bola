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

package ui

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

// constants
const (
	URIComingSoon500x500        = "https://imgur.com/nGHRg4P.png"
	URIComingSoon1920x823       = "https://imgur.com/qwPOOR2.png"
	URIUnderConstruction256x256 = "https://imgur.com/rBII01U.png"

	TextComingSoon        = "Coming soon"
	TextUnderConstruction = "Under construction"

	ColorAcidGreen = "#b0bf1a"
	ColorAero      = "#7cb9e8"
)

// ComingSoonContents function generates ExtCarouselContainer for coming soon box.
func ComingSoonContents(text string) *ExtCarouselContainer {
	img := URIComingSoon500x500
	ratio := linebot.FlexImageAspectRatioType1to1
	if text == TextComingSoon {
		img = URIComingSoon1920x823
		ratio = linebot.FlexImageAspectRatioType("1920:823")
	} else if text == TextUnderConstruction {
		img = URIUnderConstruction256x256
		ratio = linebot.FlexImageAspectRatioType1to1
	} else {
		text = TextComingSoon
	}
	contents := ExtCarouselContainer{
		CarouselContainer: linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
		},
		Contents: []*ExtBubbleContainer{
			{
				Type:      linebot.FlexContainerTypeBubble,
				Direction: linebot.FlexBubbleDirectionTypeLTR,
				Header: &ExtBoxComponent{
					BoxComponent: linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Text:   text,
								Align:  linebot.FlexComponentAlignTypeCenter,
								Margin: linebot.FlexComponentMarginTypeNone,
								Size:   linebot.FlexTextSizeTypeLg,
								Color:  ColorAero,
							},
						},
					},
				},
				Hero: &linebot.ImageComponent{
					Type:        linebot.FlexComponentTypeImage,
					AspectRatio: ratio,
					AspectMode:  linebot.FlexImageAspectModeTypeCover,
					URL:         img,
					Size:        linebot.FlexImageSizeTypeFull,
				},
			},
		},
	}
	return &contents
}

// ComingSoonMessage function generates FlexMessage for coming soon message.
func ComingSoonMessage() *linebot.FlexMessage {
	contents := ComingSoonContents(TextComingSoon)
	return linebot.NewFlexMessage(TextComingSoon, contents)
}

// UnderConstructionMessage function generates FlexMessage for under
// construction message.
func UnderConstructionMessage() *linebot.FlexMessage {
	contents := ComingSoonContents(TextUnderConstruction)
	return linebot.NewFlexMessage(TextUnderConstruction, contents)
}

// Dump function dumps the instance of ExtCarouselContainer with json format.
func (ec *ExtCarouselContainer) Dump() {
	if json, err := ec.MarshalJSON(); err == nil {
		fmt.Printf("%s\n", json)
	}
}

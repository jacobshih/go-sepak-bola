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

// IBubble interface is used for bubble container.
type IBubble interface {
	Type() linebot.FlexContainerType
	Direction() linebot.FlexBubbleDirectionType
	Header() *linebot.BoxComponent
	Hero() *linebot.ImageComponent
	Body() *linebot.BoxComponent
	Footer() *linebot.BoxComponent
	Styles() *linebot.BubbleStyle
}

// Bubble function creates a bubble container from the instance that implements
// IBubble interface.
func Bubble(bubble IBubble) *linebot.BubbleContainer {
	return &linebot.BubbleContainer{
		Type:   bubble.Type(),
		Header: bubble.Header(),
		Hero:   bubble.Hero(),
		Body:   bubble.Body(),
		Footer: bubble.Footer(),
		Styles: bubble.Styles(),
	}
}

// DumpCarouselContainer function dumps the instance of CarouselContainer with
// json format.
func DumpCarouselContainer(cc *linebot.CarouselContainer) {
	if json, err := cc.MarshalJSON(); err == nil {
		fmt.Printf("%s\n", json)
	}
}

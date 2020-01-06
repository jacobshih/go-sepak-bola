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
	"encoding/json"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ExtBubbleContainer type
type ExtBubbleContainer struct {
	Type      linebot.FlexContainerType
	Direction linebot.FlexBubbleDirectionType
	Header    *ExtBoxComponent
	Hero      *linebot.ImageComponent
	Body      *ExtBoxComponent
	Footer    *ExtBoxComponent
	Styles    *linebot.BubbleStyle
}

// MarshalJSON method of ExtBubbleContainer
func (c *ExtBubbleContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      linebot.FlexContainerType       `json:"type"`
		Direction linebot.FlexBubbleDirectionType `json:"direction,omitempty"`
		Header    *ExtBoxComponent                `json:"header,omitempty"`
		Hero      *linebot.ImageComponent         `json:"hero,omitempty"`
		Body      *ExtBoxComponent                `json:"body,omitempty"`
		Footer    *ExtBoxComponent                `json:"footer,omitempty"`
		Styles    *linebot.BubbleStyle            `json:"styles,omitempty"`
	}{
		Type:      linebot.FlexContainerTypeBubble,
		Direction: c.Direction,
		Header:    c.Header,
		Hero:      c.Hero,
		Body:      c.Body,
		Footer:    c.Footer,
		Styles:    c.Styles,
	})
}

// ExtCarouselContainer type
type ExtCarouselContainer struct {
	// Type     linebot.FlexContainerType
	linebot.CarouselContainer
	Contents []*ExtBubbleContainer
}

// MarshalJSON method of ExtCarouselContainer
func (c *ExtCarouselContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type     linebot.FlexContainerType `json:"type"`
		Contents []*ExtBubbleContainer     `json:"contents"`
	}{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: c.Contents,
	})
}

// FlexContainer implements FlexContainer interface
func (*ExtBubbleContainer) FlexContainer() {}

// FlexContainer implements FlexContainer interface
func (*ExtCarouselContainer) FlexContainer() {}

// ExtBoxComponent type extends BoxComponent with some unsupported properties.
type ExtBoxComponent struct {
	linebot.BoxComponent
	// add the available properties that are not supported in linebot sdk.
	Width           string
	Height          string
	BackgroundColor string
	BorderColor     string
	BorderWidth     string // none, light, normal, medium, semi-bold, or bold.
	CornerRadius    string // none, xs, sm, md, lg, xl, or xxl.
	PaddingAll      string //  // none, xs, sm, md, lg, xl, or xxl.
	PaddingTop      string //  // none, xs, sm, md, lg, xl, or xxl.
	PaddingBottom   string //  // none, xs, sm, md, lg, xl, or xxl.
	PaddingStart    string //  // none, xs, sm, md, lg, xl, or xxl.
	PaddingEnd      string //  // none, xs, sm, md, lg, xl, or xxl.
}

// MarshalJSON method of ExtBoxComponent
func (c *ExtBoxComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type            linebot.FlexComponentType        `json:"type"`
		Layout          linebot.FlexBoxLayoutType        `json:"layout"`
		Contents        []linebot.FlexComponent          `json:"contents"`
		Flex            *int                             `json:"flex,omitempty"`
		Spacing         linebot.FlexComponentSpacingType `json:"spacing,omitempty"`
		Margin          linebot.FlexComponentMarginType  `json:"margin,omitempty"`
		Width           string                           `json:"width,omitempty"`
		Height          string                           `json:"height,omitempty"`
		BackgroundColor string                           `json:"backgroundColor,omitempty"`
		BorderColor     string                           `json:"borderColor,omitempty"`
		BorderWidth     string                           `json:"borderWidth,omitempty"`
		CornerRadius    string                           `json:"cornerRadius,omitempty"`
		PaddingAll      string                           `json:"paddingAll,omitempty"`
		PaddingTop      string                           `json:"paddingTop,omitempty"`
		PaddingBottom   string                           `json:"paddingBottom,omitempty"`
		PaddingStart    string                           `json:"paddingStart,omitempty"`
		PaddingEnd      string                           `json:"paddingEnd,omitempty"`
	}{
		Type:            linebot.FlexComponentTypeBox,
		Layout:          c.Layout,
		Contents:        c.Contents,
		Flex:            c.Flex,
		Spacing:         c.Spacing,
		Margin:          c.Margin,
		Width:           c.Width,
		Height:          c.Height,
		BackgroundColor: c.BackgroundColor,
		BorderColor:     c.BorderColor,
		BorderWidth:     c.BorderWidth,
		CornerRadius:    c.CornerRadius,
		PaddingAll:      c.PaddingAll,
		PaddingTop:      c.PaddingTop,
		PaddingBottom:   c.PaddingBottom,
		PaddingStart:    c.PaddingStart,
		PaddingEnd:      c.PaddingEnd,
	})
}

// FlexComponent implements FlexComponent interface
func (*ExtBoxComponent) FlexComponent() {}

// IBubble interface is used for bubble container.
type IBubble interface {
	Type() linebot.FlexContainerType
	Direction() linebot.FlexBubbleDirectionType
	Header() *ExtBoxComponent
	Hero() *linebot.ImageComponent
	Body() *ExtBoxComponent
	Footer() *ExtBoxComponent
	Styles() *linebot.BubbleStyle
}

// Bubble function creates a bubble container from the instance that implements
// IBubble interface.
func Bubble(bubble IBubble) *ExtBubbleContainer {
	return &ExtBubbleContainer{
		Type:   bubble.Type(),
		Header: bubble.Header(),
		Hero:   bubble.Hero(),
		Body:   bubble.Body(),
		Footer: bubble.Footer(),
		Styles: bubble.Styles(),
	}
}

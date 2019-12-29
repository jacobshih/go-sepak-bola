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

package main

import (
	"encoding/json"
	"fmt"
	"go-sepak-bola/gsb"
	"go-sepak-bola/internal/appdata"
	"go-sepak-bola/internal/fbd"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var fbdapi *fbd.FBDAPI
var sepakbola gsb.SepakBola

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	if err != nil {
		log.Println("Bot:", bot, " err:", err)
	}

	fbdapi = fbd.Instance()
	token := os.Getenv("FBDATA_TOKEN")
	fbdapi.SetToken(token)

	/*
	 * initialize sepakbola
	 */
	sepakbola.Init()
	sepakbola.GetCompetitions()

	/*
	 * perform mytest if environment variable MYTEST is set.
	 */
	if os.Getenv("MYTEST") == "y" {
		mytest()
	}

	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			handleEventForMessage(event)
		case linebot.EventTypeFollow:
			handleEventForFollow(event)
		case linebot.EventTypeUnfollow:
			handleEventForUnfollow(event)
		case linebot.EventTypeJoin:
			handleEventForJoin(event)
		case linebot.EventTypeLeave:
			handleEventForLeave(event)
		case linebot.EventTypeMemberJoined:
			handleEventForMemberJoined(event)
		case linebot.EventTypeMemberLeft:
			handleEventForMemberLeft(event)
		case linebot.EventTypePostback:
			handleEventForPostback(event)
		case linebot.EventTypeBeacon:
			handleEventForBeacon(event)
		case linebot.EventTypeAccountLink:
			handleEventForAccountLink(event)
		case linebot.EventTypeThings:
			handleEventForThings(event)
		}

		if event.Type == linebot.EventTypeMessage {
		}
	}
}

func printEventTypeAndSource(event *linebot.Event) {
	msg := ""
	msg += "[" + string(event.Type) + "] "
	msg += "(" + string(event.Source.Type) + ") "
	if len(event.Source.UserID) > 0 {
		msg += "user(" + event.Source.UserID + ")" + " "
	}
	if len(event.Source.GroupID) > 0 {
		msg += "group(" + event.Source.GroupID + ")" + " "
	}
	if len(event.Source.RoomID) > 0 {
		msg += "room(" + event.Source.RoomID + ")" + " "
	}
	fmt.Println(msg)
}

func handleEventForMessage(event *linebot.Event) {
	printEventTypeAndSource(event)
	var err error
	switch message := event.Message.(type) {
	case *linebot.StickerMessage:
		if message.PackageID == "2" && message.StickerID == "504" {
			msg := sepakbola.CompetitionsMessage()
			if _, err = bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
				log.Print(err)
			}
		}
	}
}

func handleEventForFollow(event *linebot.Event) {
	printEventTypeAndSource(event)
}
func handleEventForUnfollow(event *linebot.Event) {
	printEventTypeAndSource(event)
}
func handleEventForJoin(event *linebot.Event) {
	printEventTypeAndSource(event)
}
func handleEventForLeave(event *linebot.Event) {
	printEventTypeAndSource(event)
}
func handleEventForMemberJoined(event *linebot.Event) {
	printEventTypeAndSource(event)
}
func handleEventForMemberLeft(event *linebot.Event) {
	printEventTypeAndSource(event)
}
func handleEventForPostback(event *linebot.Event) {
	printEventTypeAndSource(event)
	var msg linebot.SendingMessage
	var postdata appdata.PostData
	if err := json.Unmarshal([]byte(event.Postback.Data), &postdata); err != nil {
		fmt.Println(err)
	} else {
		switch postdata.Category {
		case gsb.PkgName:
			msg = handleEventForPostbackOfGSB(event)
		}
		if msg != nil {
			if _, err = bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
				log.Print(err)
			}
			return
		}
	}
}

func handleEventForPostbackOfGSB(event *linebot.Event) (msg linebot.SendingMessage) {
	msg = nil
	var postdata appdata.PostData
	if err := json.Unmarshal([]byte(event.Postback.Data), &postdata); err != nil {
		log.Println(err)
	} else {
		dict := postdata.Params.(map[string]interface{})
		id := int(dict["id"].(float64))
		competition := sepakbola.Competitions[id]
		switch postdata.Action {
		case gsb.ActionMatches:
			msg = sepakbola.MatchdesTypeMessage(competition, 0)
		case gsb.ActionMatchday:
			currentMatchday := int(dict["currentMatchday"].(float64))
			msg = sepakbola.MatchdayMessage(competition, currentMatchday)
		case gsb.ActionAllMatches:
			fmt.Println("Action: ", gsb.ActionAllMatches)
		case gsb.ActionStandings:
			msg = sepakbola.StandingsMessage(competition)
		case gsb.ActionTeams:
			msg = sepakbola.TeamsMessage(competition)
		case gsb.ActionTeam:
			fmt.Println("Action: ", gsb.ActionTeam, "id: ", id)
		default:
		}
	}
	return msg
}

func handleEventForBeacon(event *linebot.Event) {
	printEventTypeAndSource(event)
}
func handleEventForAccountLink(event *linebot.Event) {
	printEventTypeAndSource(event)
}
func handleEventForThings(event *linebot.Event) {
	printEventTypeAndSource(event)
}

func handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	return nil
}

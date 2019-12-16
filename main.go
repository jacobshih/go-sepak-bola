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
	"strconv"

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
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				quota, err := bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				if message.PackageID == "2" && message.StickerID == "504" {
					msg := createMessageForMenu()
					if _, err = bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}

// pkg constants
const (
	EmojiSoccer          string = "\xe2\x9a\xbd"
	EmojiTrophy          string = "\xf0\x9f\x8f\x86"
	EmojiSquaredVS       string = "\xf0\x9f\x86\x9a"
	CategoryCompetition  string = "competition"
	ActionMatchday       string = "matchday"
	ActionStandings      string = "standings"
	ActionTeams          string = "teams"
	TextMatchday         string = "Matchday"
	TextStandings        string = "Standings"
	TextTeams            string = "Teams"
	ImagePremierLeague   string = "https://www.thesportsdb.com/images/media/league/badge/i6o0kh1549879062.png"
	ImageBundesliga      string = "https://www.thesportsdb.com/images/media/league/badge/0j55yv1534764799.png"
	ImagePrimeraDivision string = "https://www.thesportsdb.com/images/media/league/badge/7onmyv1534768460.png"
	TextPremierLeague    string = "Premier League"
	TextBundesliga       string = "Bundesliga"
	TextPrimeraDivision  string = "Primera Division"
	CodePremierLeague    string = "PL"
	CodeBundesliga       string = "BL1"
	CodePrimeraDivision  string = "PD"
)

func matchdayData(competition, code string) string {
	data, _ := json.Marshal(&appdata.PostData{
		Category: CategoryCompetition,
		Action:   ActionMatchday,
		Params:   nil,
	})
	return string(data)
}

func standingsData(competition, code string) string {
	data, _ := json.Marshal(&appdata.PostData{
		Category: CategoryCompetition,
		Action:   ActionStandings,
		Params:   nil,
	})
	return string(data)
}

func teamsData(competition, code string) string {
	data, _ := json.Marshal(&appdata.PostData{
		Category: CategoryCompetition,
		Action:   ActionTeams,
		Params:   nil,
	})
	return string(data)
}

func createMessageForMenu() linebot.SendingMessage {
	var msg linebot.SendingMessage
	var competition string
	var code string
	competition = "Premier League"
	code = "PL"
	column0 := linebot.NewCarouselColumn(
		ImagePremierLeague, TextPremierLeague, CodePremierLeague,
		linebot.NewPostbackAction(TextMatchday, matchdayData(competition, code), "", ""),
		linebot.NewPostbackAction(TextStandings, standingsData(competition, code), "", ""),
		linebot.NewPostbackAction(TextTeams, teamsData(competition, code), "", ""),
	)
	competition = "Bundesliga"
	code = "BL1"
	column1 := linebot.NewCarouselColumn(
		ImageBundesliga, TextBundesliga, CodeBundesliga,
		linebot.NewPostbackAction(TextMatchday, matchdayData(competition, code), "", ""),
		linebot.NewPostbackAction(TextStandings, standingsData(competition, code), "", ""),
		linebot.NewPostbackAction(TextTeams, teamsData(competition, code), "", ""),
	)
	competition = "Primera Division"
	code = "PD"
	column2 := linebot.NewCarouselColumn(
		ImagePrimeraDivision, TextPrimeraDivision, CodePrimeraDivision,
		linebot.NewPostbackAction(TextMatchday, matchdayData(competition, code), "", ""),
		linebot.NewPostbackAction(TextStandings, standingsData(competition, code), "", ""),
		linebot.NewPostbackAction(TextTeams, teamsData(competition, code), "", ""),
	)
	template := linebot.NewCarouselTemplate(
		column0,
		column1,
		column2,
	)
	altText := "go-sepak-bola"
	msg = linebot.NewTemplateMessage(altText, template)
	return msg
}

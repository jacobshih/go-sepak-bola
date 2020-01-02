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

	"github.com/line/line-bot-sdk-go/linebot"
)

// MatchdesTypeMessage function reply quick reply button message.
func (sepakbola *SepakBola) MatchdesTypeMessage(competition *fbd.Competition, matchday int) linebot.SendingMessage {
	allMatchesData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionAllMatches,
		Params:   nil,
	})
	matchdayData, _ := json.Marshal(&appdata.PostData{
		Category: PkgName,
		Action:   ActionMatchday,
		Params: map[string]interface{}{
			"id":              competition.ID,
			"currentMatchday": competition.Season.CurrentMatchday,
		},
	})

	return linebot.NewTextMessage(TextMatchesType).
		WithQuickReplies(linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton(
				"",
				linebot.NewPostbackAction(TextAllMatches, string(allMatchesData), "", TextAllMatches)),
			linebot.NewQuickReplyButton(
				"",
				linebot.NewPostbackAction(TextCurrentMatchday, string(matchdayData), "", TextCurrentMatchday)),
		))
}

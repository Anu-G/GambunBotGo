// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"gambunbot/gacha"
	"gambunbot/translate"

	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
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
					if strings.Contains(message.Text, "$apakah gacha") {
						if strings.Contains(message.Text, "$draw") && strings.Contains(message.Text, "$rate") {
							splitter := strings.Split(message.Text, "$")
							draw, _ := strconv.Atoi(strings.Split(splitter[2], " ")[1])
							rate, _ := strconv.Atoi(strings.Split(splitter[3], " ")[1])

							luckMessage, luck := gacha.GachaPercentage()
							simMessage := gacha.GachaSim(draw, rate, 1, luck)
							replyMessage := luckMessage + "\n" + simMessage

							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
								log.Print(err)
							}
						} else {
							replyMessage, _ := gacha.GachaPercentage()

							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
								log.Print(err)
							}
						}
					}

					if strings.Contains(message.Text, "$gacha sim") {
						splitter := strings.Split(message.Text, "$")
						draw, _ := strconv.Atoi(strings.Split(splitter[2], " ")[1])
						rate, _ := strconv.Atoi(strings.Split(splitter[3], " ")[1])

						replyMessage := gacha.GachaSim(draw, rate, 1, 0)

						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
							log.Print(err)
						}
					}

					if strings.Contains(message.Text, "$tl ") {
						splitCase := strings.Split(message.Text, " ")
						var toTrannslate []string
						replyMessage := ""

						switch splitCase[1] {
						case "jpen":
							toTrannslate = strings.Split(message.Text, "$tl jpen ")
							replyMessage = translate.TranslateJPtoEN(toTrannslate[1])
						case "jpid":
							toTrannslate = strings.Split(message.Text, "$tl jpid ")
							replyMessage = translate.TranslateJPtoID(toTrannslate[1])
						case "enjp":
							toTrannslate = strings.Split(message.Text, "$tl enjp ")
							replyMessage = translate.TranslateENtoJP(toTrannslate[1])
						case "idjp":
							toTrannslate = strings.Split(message.Text, "$tl idjp ")
							replyMessage = translate.TranslateIDtoJP(toTrannslate[1])
						}

						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
							log.Print(err)
						}
					}

					if strings.Contains(message.Text, "sticker") {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("446", "1989")).Do(); err != nil {
							log.Print(err)
						}
					}

					if strings.Contains(message.Text, "image") {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://i.imgur.com/vmdCE2r.png", "")).Do(); err != nil {
							log.Print(err)
						}
					}

					if strings.Contains(message.Text, "youtube") {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewVideoMessage("https://youtu.be/Cn7O5ibZfqc", "")).Do(); err != nil {
							log.Print(err)
						}
					}

					if strings.Contains(message.Text, "audio") {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewAudioMessage("http", 10)).Do(); err != nil {
							log.Print(err)
						}
					}

					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					// 	log.Print(err)
					// }

					// case *linebot.StickerMessage:
					// 	replyMessage := fmt.Sprintf(
					// 		"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					// 		log.Print(err)
					// 	}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

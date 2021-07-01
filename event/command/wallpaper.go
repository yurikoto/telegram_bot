package command

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
	"time"
)

// Wallpaper 返回随机壁纸
func Wallpaper(bot *telebot.Bot) {
	bot.Handle("/wallpaper", func(m *telebot.Message) {
		photo := &telebot.Photo{
			File: telebot.FromURL("https://v1.yurikoto.com/wallpaper?type=rand&id=" + strconv.FormatInt(time.Now().UnixNano(), 10)),
		}
		_, err := photo.Send(bot, m.Chat, &telebot.SendOptions{
			ReplyTo: m,
		})
		if err != nil {
			fmt.Printf("发送消息时发生了错误，错误信息： %s \n", err)
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
		}
	})
}

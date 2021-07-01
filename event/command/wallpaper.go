package command

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
	"time"
	"yurikoto.com/yurikoto-telegram-bot/middlewares"
)

// Wallpaper 返回随机壁纸
func Wallpaper(bot *telebot.Bot) {
	bot.Handle("/wallpaper", func(m *telebot.Message) {
		if !middlewares.RateLimit(m.Sender.Username) {
			msg := "您的调用次数已超限，请稍后再试"
			_, _ = bot.Send(m.Chat, msg)
			return
		}

		photo := &telebot.Photo{
			File: telebot.FromURL("https://v1.yurikoto.com/wallpaper?type=rand&id=" + strconv.FormatInt(time.Now().UnixNano(), 10)),
		}
		_, err := photo.Send(bot, m.Chat, &telebot.SendOptions{
			ReplyTo: m,
		})
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
		}
	})
}

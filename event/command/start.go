package command

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

// Start 用于响应 Telegram 要求的机器人 Start 指令
func Start(bot *telebot.Bot) {
	bot.Handle("/start", func(m *telebot.Message) {
		if !m.Private() { // 如果不是私发消息，不回复
			return
		}
		_, err := bot.Send(m.Sender, `欢迎您使用Yurikoto。
你可以...
使用 /about 以了解本机器人，
使用 /help  查看使用机器人的说明，
每位用户每小时限制调用100次`)
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
		}
	})
}

package event

import (
	"gopkg.in/tucnak/telebot.v2"
	"yurikoto.com/yurikoto-telegram-bot/event/command"
)

// RegisterEvent 定义了机器人事件处理的入口
func RegisterEvent(bot *telebot.Bot) {
	registerCommand(bot)
}

func registerCommand(bot *telebot.Bot) {
	command.Sentence(bot)
	command.Wallpaper(bot)
	command.Start(bot)
	command.About(bot)
	command.Ping(bot)
	command.Help(bot)
}

package event

import (
	"gopkg.in/tucnak/telebot.v2"
	"yurikoto.com/yurikoto-telegram-bot/event/command"
)

// RegisterEvent 定义了机器人事件处理的入口
func RegisterEvent(bot *telebot.Bot) {
	registerCommand(bot)
}

// TODO: telebot-v3更新后封装频次控制到中间件

func registerCommand(bot *telebot.Bot) {
	// bot.Use(RateLimit)
	command.Sentence(bot)
	command.Wallpaper(bot)
	command.Start(bot)
	command.About(bot)
	command.Ping(bot)
	command.Help(bot)
}

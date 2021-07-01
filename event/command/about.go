package command

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"time"
)

// About 返回关于信息
func About(b *telebot.Bot) {
	b.Handle("/about", func(m *telebot.Message) {
		_, err := b.Send(m.Chat, fmt.Sprintf(`Yurikoto 官方 Telegram 机器人。 目前仅提供简体中文支持。 提供台词、壁纸服务。
* 官方网站: https://yurikoto.com
* Telegram 群组: https://t.me/yurikoto_group
* 项目开源地址：https://github.com/yurikoto
--------------
当前服务器时间：%s`, time.Now().Format("2006年1月2日 15:04:05")))
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
		}
	})
}

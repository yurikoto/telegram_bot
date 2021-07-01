package command

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"time"
)

// Help 返回使用帮助
func Help(b *telebot.Bot) {
	b.Handle("/help", func(m *telebot.Message) {
		_, err := b.Send(m.Chat, fmt.Sprintf(`*以下是目前支持的指令：*
/about 关于机器人
/sentence 获取一条随机台词。
/help 获取机器人帮助信息。
/wallpaper 获取一张随机壁纸。
/ping 测试机器人连通性。
--------------
当前服务器时间：%s`, time.Now().Format("2006年1月2日 15:04:05")),
			&telebot.SendOptions{
				ParseMode: "markdown",
			},
		)
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
		}
	})
}

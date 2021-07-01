package command

import (
	"fmt"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"yurikoto.com/yurikoto-telegram-bot/middlewares"
	// "../event"
)

// Sentence 获取台词
func Sentence(b *telebot.Bot) {
	b.Handle("/sentence", func(m *telebot.Message) {
		if !middlewares.RateLimit(m.Sender.Username) {
			msg := "您的调用次数已超限，请稍后再试"
			_, _ = b.Send(m.Chat, msg)
			return
		}

		url := "https://v1.yurikoto.com/sentence"

		// 请求接口
		response, err := grequests.Get(url, nil)
		if err != nil {
			log.Errorf("尝试获取台词时出现错误，错误信息： %s\n", err)
			_, err = b.Send(m.Chat, "很抱歉，尝试获取发生错误。")
			if err != nil {
				log.Errorf("尝试发送消息时出现错误，错误信息：%s \n", err)
			}
			return
		}
		data := &YurikotoSentenceAPIV1Response{}
		err = response.JSON(data)
		if err != nil {
			log.Errorf("尝试解析台词时发生错误，错误信息： %s", err)
			_, err = b.Send(m.Chat, "很抱歉，尝试解析台词时发生错误。")
			if err != nil {
				log.Errorf("尝试发送消息时出现错误，错误信息：%s \n", err)
			}
			return
		}
		_, err = b.Reply(m, fmt.Sprintf(`%s ——「%s」`, data.Content, data.Source))
		if err != nil {
			log.Errorf("尝试发送消息时出现错误，错误信息：%s \n", err)
		}
	})
}

// YurikotoSentenceAPIV1Response 定义了台词接口的结构
type YurikotoSentenceAPIV1Response struct {
	ID      uint32 `json:"id"`
	Content string `json:"content"`
	Source  string `json:"source"`
	Status  string `json:"status"`
}

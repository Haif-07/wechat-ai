package plugins

import (
	"log"
	"strings"
	"wechat-ai/chat"

	"github.com/eatmoreapple/openwechat"
	"github.com/spf13/viper"
)

func GeneralChat(msg *openwechat.Message) {
	if msg.IsSendByGroup() && msg.IsAt() {
		sendUser, err := msg.Sender()
		if err != nil {
			log.Printf("get sender error: %v", err)
		}
		if sendGroup, ok := sendUser.AsGroup(); ok {
			group := viper.GetString("wechat.group")
			groups := strings.Split(group, ",")
			for _, g := range groups {
				if sendGroup.NickName == g {
					msgs := strings.Split(msg.Content, " ")
					// msgss := strings.Split(msgs[0], " ")
					comman := msgs[1]
					if comman == "解答" {
						response, err := chat.RunChatAi(msgs[2], viper.GetString("plugins.general-chat.prompt"))
						if err != nil {
							log.Printf("chat ai error: %v", err)
						}
						msg.ReplyText(response)
					}

				}
			}

		}
	}
}

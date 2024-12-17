package main

import (
	"fmt"

	"github.com/eatmoreapple/openwechat"
)

func main() {
	// bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
	// // 登陆
	// if err := bot.Login(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	bot := openwechat.DefaultBot(openwechat.Desktop)
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {

		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
		if msg.IsSendByGroup() && msg.IsAt() {
			//群里面呗@了

			sendUser, err := msg.Sender()
			if err != nil {
				fmt.Println(err)
				return
			}
			if sendGroup, ok := sendUser.AsGroup(); ok {
				if sendGroup.NickName == "测试群" {
					//将消息写进一个消息数组里面，容量为10条，有新的就更新，保存起来用于ai分析

				}
			}

		}
		fmt.Println(msg.Content)
	}
	if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		fmt.Println(err)
		return
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// // 获取所有的好友s

	// friends, err := self.Friends()
	// fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	fmt.Println(groups, err)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}

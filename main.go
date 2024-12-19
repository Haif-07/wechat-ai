package main

import (
	"log"
	"wechat-ai/config"
	"wechat-ai/plugins"

	"github.com/eatmoreapple/openwechat"
	"github.com/spf13/viper"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Printf("load config error: %v", err)
		return
	}

	var wcLogin openwechat.BotPreparer
	if viper.GetString("wechat.login") == "web" {
		wcLogin = openwechat.Normal
	} else {
		wcLogin = openwechat.Desktop
	}
	bot := openwechat.DefaultBot(wcLogin)
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		plugins.GeneralChat(msg)
	}
	//登陆
	if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		log.Printf("login error: %v", err)
		return
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Printf("get user error: %v", err)
		return
	}
	log.Printf("登陆成功: %s\n", self.NickName)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}

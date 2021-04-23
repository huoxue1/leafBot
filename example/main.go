package main

import (
	"fmt"
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/plugins"
	"os"
)

func init() {
	// 为bot添加weather响应器，命令为 ”/天气“ ,allies为命令别名，
	//参数格式为一个字符串数组，rule为一个结构体，响应前会先判断所以rules为true，weight为权重，block为是否阻断
	leafBot.AddCommandHandle(Weather, "/天气", nil, nil, 10, false)

	plugins.UseCreateQrCode()               //加载生成二维码插件
	plugins.UseDayImage()                   // 加载每日一图插件
	plugins.UseEchoHandle()                 // 加载echo插件
	plugins.UseMusicHandle()                // 加载音乐插件
	plugins.UseSetuHandle()                 // 加载涩图插件
	plugins.UseTranslateHandle()            // 加载翻译插件
	plugins.UseFlashImage(0)                // 加载闪照破解插件
	plugins.UseFlashImageToGroup(972264701) //加载闪照破解后发到对应群的插件

	leafBot.AddCommandHandle(func(event leafBot.Event, bot *leafBot.Bot, args []string) {
		if event.UserId == 3343780376 {
			for i := 0; i < 10; i++ {
				oneEvent, err := bot.GetOneEvent(leafBot.Rule{
					RuleCheck: func(event leafBot.Event, i ...interface{}) bool {
						if event.UserId == 3343780376 {
							return true
						}
						return false
					},
					Dates: nil,
				})
				if err != nil {
					bot.Send(event, fmt.Sprintf("这是第%v条信息:  "+err.Error(), i))
				} else {
					bot.Send(event, fmt.Sprintf("这是第%v条信息:  "+oneEvent.Message, i))
				}

			}
		}
	}, "/he", nil, nil, 10, false)
}

func main() {
	dir, _ := os.Getwd() // 获取当前路径
	if len(os.Args) > 1 {
		leafBot.LoadConfig(os.Args[1], leafBot.JSON)
	} else {
		leafBot.LoadConfig(dir+"/example/config.json", leafBot.JSON)
	}

	//拼接配置文件路径，并且加载配置文件
	leafBot.InitBots() //初始化Bot
}

/*
	event: bot的event，里面包含了事件的所有字段
	bot: 触发事件的bot指针
	args ： 命令的参数，为一个数组
*/
func Weather(event leafBot.Event, bot *leafBot.Bot, args []string) {
	m := map[string]string{"北京": "晴", "山东": "下雨"}
	// 调用发送消息的api，会根据messageType自动回复
	bot.SendMsg(event.MessageType, event.UserId, event.GroupId,
		args[0]+"的天气为"+m[args[0]],
		false)
}

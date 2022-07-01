package process

import (
	"MassUserCommunication/client/utils"
	"MassUserCommunication/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {

	fmt.Println("------------恭喜登录成功----------------------")
	fmt.Println("------------1: 显示用户在线列表---------------")
	fmt.Println("------------2：群发消息----------------------")
	fmt.Println("------------3: 发送消息----------------------")
	fmt.Println("------------4: 信息列表---------------------")
	fmt.Println("------------5: 退出系统----------------------")
	fmt.Println("请选择（1-5）:")

	var key int
	var context string
	var fromUserId int
	var toUserId int
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d \n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("你想对大家说些什么?")
		fmt.Scanf("%s \n", &context)
		smsProcess.sendGroupMes(context)
	case 3:
		outputOnlineUser()
		fmt.Print("请输入用户id:")
		fmt.Scanf("%d \n", &toUserId)
		fmt.Println("你想对大家说些什么?")
		fmt.Scanf("%s \n", &context)
		fromUserId = CurrentUser.UserId
		smsProcess.sendUserSms(fromUserId, toUserId, context)
	case 4:
		showSmsList()
	case 5:
		os.Exit(0)
	default:

	}
}

func serverProcessMes(conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err= ", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		case message.SmsMesOneType:
			addSmsList(&mes)
		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}
	}
}

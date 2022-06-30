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

	fmt.Println("------------恭喜xxx登录成功--------------------")
	fmt.Println("------------1: 显示用户在线列表---------------")
	fmt.Println("------------2：发送消息----------------------")
	fmt.Println("------------3: 信息列表---------------------")
	fmt.Println("------------4: 退出系统----------------------")
	fmt.Println("请选择（1-4）:")

	var key int
	fmt.Scanf("%d \n", &key)
	switch key {
	case 1:
		fmt.Println("显示用户在线列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		os.Exit(0)
	default:
		fmt.Println("输入数据格式不对")
	}
}

func serverProcessMes(conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err= ", err)
			return
		}
		fmt.Println("mes = %v \n", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}
	}
}

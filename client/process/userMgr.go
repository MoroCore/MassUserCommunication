package process

import (
	"MassUserCommunication/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

func updateUserStatus(mes *message.NotifyUserStatusMes) {

	user := &message.User{
		UserId:     mes.UserId,
		UserStatus: mes.Status,
	}
	onlineUsers[mes.UserId] = user
	outputOnlineUser()
}
func outputOnlineUser() {
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		fmt.Println("用户id : \t", id)
	}
}

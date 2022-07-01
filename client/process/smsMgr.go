package process

import (
	"MassUserCommunication/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {

	var smsMes message.SmsMes

	err := json.Unmarshal([]byte(mes.Data), &smsMes)

	if err != nil {
		fmt.Println(" json.Unmarshal err = ", err.Error())
		return
	}
	info := fmt.Sprintf("用户id:\t %d 对大家说:\t%s", smsMes.UserId, smsMes.Context)
	fmt.Println(info)
	fmt.Println()
}
func addSmsList(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println(" json.Unmarshal err = ", err.Error())
		return
	}
	smsList.SmsList = append(smsList.SmsList, smsMes)
}
func showSmsList() {

	fmt.Println("消息列表")
	for _, v := range smsList.SmsList {
		fmt.Printf("用户%d向你发的内容:%s\n", v.FromUserId, v.Context)
	}
}

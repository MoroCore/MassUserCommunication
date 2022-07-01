package process

import (
	"MassUserCommunication/client/utils"
	"MassUserCommunication/common/message"
	"encoding/json"
	"fmt"
)

var smsList SmsProcess

type SmsProcess struct {
	SmsList []message.SmsMes
}

func (this *SmsProcess) sendGroupMes(context string) (err error) {

	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Context = context
	smsMes.UserId = CurrentUser.UserId
	smsMes.UserStatus = CurrentUser.UserStatus

	marshal, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal(smsMes) fail ", err)
		return
	}
	mes.Data = string(marshal)

	bytes, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal(mes) fail ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurrentUser.Conn,
	}
	err = tf.WritePkg(bytes)
	if err != nil {
		fmt.Println("sendGroupMes  tf.WritePkg(bytes) fail ", err)
	}
	return
}
func (this *SmsProcess) sendUserSms(fromUserId int, toUserId int, context string) {

	var mes message.Message
	mes.Type = message.SmsMesOneType

	var smsMes message.SmsMes
	smsMes.Context = context
	smsMes.FromUserId = fromUserId
	smsMes.ToUserId = toUserId

	marshal, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendUserSms json.Marshal(smsMes) fail ", err)
		return
	}
	mes.Data = string(marshal)
	bytes, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("sendUserSms json.Marshal(mes) fail ", err)
		return
	}
	tf := &utils.Transfer{
		Conn: CurrentUser.Conn,
	}
	err = tf.WritePkg(bytes)
	if err != nil {
		fmt.Println("sendGroupMes  tf.WritePkg(bytes) fail ", err)
	}
	return
}

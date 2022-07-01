package processes

import (
	"MassUserCommunication/common/message"
	"MassUserCommunication/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {

	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal fail ", err)
		return
	}
	marshal, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal fail ", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if smsMes.UserId != id {
			this.sendMesEachOnlineUser(marshal, up.Coon)
		}
	}
	return
}

func (this *SmsProcess) sendMesEachOnlineUser(sms []byte, conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(sms)
	if err != nil {
		fmt.Println("转发消息失败 ", err)
		return
	}
}

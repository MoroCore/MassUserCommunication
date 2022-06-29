package processes

import (
	"MassUserCommunication/common/message"
	"MassUserCommunication/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

//处理和用户有关的请求
//登录  注册  注销  用户列表管理
type UserProcess struct {
	Coon net.Conn
}

//登录逻辑
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes
	//如果用户id=100 密码=123456 认为是合法的
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500 //该用户不存在
		loginResMes.Error = "该用户不存在"
	}
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail ", err)
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	tf := &utils.Transfer{
		Conn: this.Coon,
	}
	err = tf.WritePkg(data)
	return
}

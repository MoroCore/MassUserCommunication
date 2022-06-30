package processes

import (
	"MassUserCommunication/common/message"
	"MassUserCommunication/server/model"
	"MassUserCommunication/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Coon net.Conn
}

func (this *UserProcess) ServerPRocessRegister(mes *message.Message) (err error) {
	var register message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &register)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&register.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 500
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	marshal, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail ", err)
		return
	}
	resMes.Data = string(marshal)
	data, err := json.Marshal(resMes)
	tf := &utils.Transfer{
		Conn: this.Coon,
	}
	err = tf.WritePkg(data)
	return
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
	_, err = model.MyUserDao.LoginCheck(loginMes.UserId, loginMes.UserPwd)
	if err == nil {
		loginResMes.Code = 200
	}
	if err != nil {
		loginResMes.Code = 500
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

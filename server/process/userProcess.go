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
	Coon   net.Conn
	UserId int
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
	user, err := model.MyUserDao.LoginCheck(loginMes.UserId, loginMes.UserPwd)
	if err == nil {
		loginResMes.Code = 200
		this.UserId = user.UserId
		userMgr.addOnlineUser(this)
		this.NotifyOtherOnlineUser(user.UserId)
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserId = append(loginResMes.UserId, id)
		}
		fmt.Println("登录成功")
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
func (this *UserProcess) NotifyOtherOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
		//
	}
}
func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnLine
	marshal, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal fail ", err)
		return
	}
	mes.Data = string(marshal)

	bytes, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal fail ", err)
		return
	}
	tf := &utils.Transfer{
		Conn: this.Coon,
	}
	tf.WritePkg(bytes)
}

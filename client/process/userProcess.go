package process

import (
	"MassUserCommunication/client/utils"
	"MassUserCommunication/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(coon) err = ", err)
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
	} else {
		fmt.Println(registerResMes.Error)
	}
	os.Exit(0)
	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
	}

	defer conn.Close()

	//2: 构造Message
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println(" tf.WritePkg(data) fail ", err)
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(coon) err = ", err)
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("当前用户列表如下")
		for _, v := range loginResMes.UserId {
			if v != userId {
				user := &message.User{
					UserId:     v,
					UserStatus: message.UserOnLine,
				}
				onlineUsers[v] = user
			}
		}
		fmt.Print("\n\n")
		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

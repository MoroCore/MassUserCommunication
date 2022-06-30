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

	//var pkgLen uint32
	//pkgLen = uint32(len(data))
	//var bytes [4]byte
	//buf := bytes[:4]
	//binary.BigEndian.PutUint32(buf, pkgLen)
	//n, err := conn.Write(buf)
	//if n != 4 || err != nil {
	//	fmt.Println("conn.Write(buf) fail ", err)
	//}
	//_, err = conn.Write(data)
	//if err != nil {
	//	fmt.Println("conn.Write(data) fail", err)
	//	return
	//}
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
		fmt.Println("登录成功")
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

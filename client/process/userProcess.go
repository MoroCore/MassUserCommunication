package process

import (
	"MassUserCommunication/client/utils"
	"MassUserCommunication/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	//定协议
	//fmt.Printf("userId = %d,userPwd = %s\n", userId, userPwd)
	//return
	//1:连接服务器
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

	//conn.Write()发送的是byte[]  先发送切片的长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	buf := bytes[:4]
	binary.BigEndian.PutUint32(buf, pkgLen)
	n, err := conn.Write(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail ", err)
	}
	//fmt.Printf("客户端,发送消息的长度 = %d", len(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
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

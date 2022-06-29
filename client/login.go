package main

import "fmt"

/*
 * @GoName login
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 16:07
 */

func login(userId int, userPwd string) (err error) {

	//定协议
	fmt.Printf("userId = %d,userPwd = %s\n", userId, userPwd)
	return nil
	//conn, err := net.Dial("tcp", "localhost:8889")
	//if err != nil {
	//	fmt.Println("net.Dial err = ", err)
	//	return
	//}
	//defer conn.Close()
	//var mes message.Message
	//mes.Type = message.LoginMesType
	//
	//var loginMes message.LoginMes
	//loginMes.UserId = userId
	//loginMes.UserPwd = userPwd
	//
	//data, err := json.Marshal(loginMes)
	//if err != nil {
	//	fmt.Println("json.Marshal err = ", err)
	//	return
	//}
	//mes.Data = string(data)
	//data, err = json.Marshal(mes)
	//if err != nil {
	//	fmt.Println("json.Marshal err = ", err)
	//	return
	//}
	//var pkgLen uint32
	//pkgLen = uint32(len(data))
	//var buf [4]byte
	//binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//n, err := conn.Write(buf[:4])
	////发送消息的长度
	//if n != 4 || err != nil {
	//	fmt.Println("conn.Write(bytes) fail", err)
	//	return
	//}
	////发送消息的内容
	//_, err = conn.Write(data)
	//if err != nil {
	//	fmt.Println("conn.Write(data) fail", err)
	//	return
	//}
	//time.Sleep(time.Second * 20)
	//fmt.Printf("客户端,发送消息的长度=%d 内容=%s \n", len(data), string(data))
	//return
}

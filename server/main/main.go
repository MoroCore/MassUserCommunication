package main

import (
	"fmt"
	"net"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 16:11
 */
//func readPkg(conn net.Conn) (mes message.Message, err error) {
//
//	fmt.Println("客户端发来数据")
//	buf := make([]byte, 1024*4)
//	//如果客户端的conn关闭  read方法不阻塞了
//	//循环读取内容    err = io.EOF
//	_, err = conn.Read(buf[:4])
//	if err != nil {
//		fmt.Println("read pkg head error")
//		return
//	}
//	//fmt.Println("读到的长度buf", buf[:4])
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[0:4])
//
//	n, err := conn.Read(buf[:pkgLen])
//	if uint32(n) != pkgLen || err != nil {
//		fmt.Println("read pkg boby error ")
//		return
//	}
//	err = json.Unmarshal(buf[:pkgLen], &mes)
//	if err != nil {
//		fmt.Println("json.Unmarshal err = ", err)
//	}
//	return
//}
//func writePkg(conn net.Conn, data []byte) (err error) {
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var bytes [4]byte
//	buf := bytes[:4]
//	binary.BigEndian.PutUint32(buf, pkgLen)
//	n, err := conn.Write(buf)
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write(buf) fail ", err)
//	}
//	//fmt.Printf("客户端,发送消息的长度 = %d", len(data))
//	_, err = conn.Write(data)
//	if err != nil || n != int(pkgLen) {
//		fmt.Println("conn.Write(data) fail", err)
//		return
//	}
//	return
//}

//登录逻辑
//func serverProcessLogin(coon net.Conn, mes *message.Message) (err error) {
//
//	var loginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data), &loginMes)
//	if err != nil {
//		fmt.Println("json.Unmarshal fail err = ", err)
//		return
//	}
//	var resMes message.Message
//	resMes.Type = message.LoginResMesType
//
//	var loginResMes message.LoginResMes
//	//如果用户id=100 密码=123456 认为是合法的
//	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
//		loginResMes.Code = 200
//	} else {
//		loginResMes.Code = 500 //该用户不存在
//		loginResMes.Error = "该用户不存在"
//	}
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json.Marshal fail ", err)
//	}
//	resMes.Data = string(data)
//	data, err = json.Marshal(resMes)
//	err = writePkg(coon, data)
//	return
//}

//编写一个ServerProcessMes函数
//功能：根据客户端发送的消息种类不同  决定调用哪个函数
//func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
//
//	switch mes.Type {
//	case message.LoginMesType:
//		//处理登录逻辑
//		err = serverProcessLogin(conn, mes)
//	case message.RegisterMesType:
//	default:
//		fmt.Println("消息类型不存在,无法处理")
//	}
//	return nil
//}
func process(conn net.Conn) (err error) {

	defer conn.Close()

	processor := &Processor{
		Conn: conn,
	}
	err = processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端连接出问题")
		return
	}
	return
}
func main() {

	fmt.Println("服务器在8889端口监听.......")
	listener, err := net.Listen("tcp", "0.0.0.0:9001")
	defer listener.Close()
	if err != nil {
		//如果监听的端口已被占用
		fmt.Println("net.Listen err = ", err)
		return
	}
	for {
		fmt.Println("等待客户端来连接服务器......")
		conn, err := listener.Accept() //阻塞  登录客户端连接
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
		}
		go process(conn)
	}
}

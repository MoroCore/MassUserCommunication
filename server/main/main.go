package main

import (
	"fmt"
	"net"
	"time"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 16:11
 */

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

	initPool("192.168.10.10:6379", 16, 0, 300*time.Second)
	initUserDao()

	fmt.Println("服务器在9001端口监听.......")
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

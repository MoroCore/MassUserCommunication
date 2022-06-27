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

func main() {

	fmt.Println("服务器在8889端口监听.......")
	listener, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	for {
		fmt.Println("等待客户端来连接服务器......")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
		}
		go process(conn)
	}
}
func process(conn net.Conn) {

	defer conn.Close()

	for {
		buf := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}
		fmt.Println("读取到buf = ", buf[:4])
	}
}

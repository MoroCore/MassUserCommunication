package main

import (
	"MassUserCommunication/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
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
	listener, err := net.Listen("tcp", "0.0.0.0:9001")
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

		_, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务器段也退出")
				return
			} else {
				fmt.Println("readPkg err = ", err)
				return
			}
		}
		//err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}
func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

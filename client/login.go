package main

import (
	"fmt"
	"net"
)

/*
 * @GoName login
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 16:07
 */

func login(userId int, userPwd string) (err error) {

	//fmt.Printf("userId = %d  userPwd = %s \n", userId, userPwd)
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	defer conn.Close()

}

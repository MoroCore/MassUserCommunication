package main

import (
	"fmt"
	"os"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 15:47
 */

var userId int
var userPwd string

func main() {

	//接收用户的选择
	var key int
	//是否继续显示菜单
	var loop = true
	for loop {
		fmt.Println("------------------------------欢迎登录多人聊天系统------------------------------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Print("\t\t\t 请选择(1-3):")

		// \n 解决一些格式化信息
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误,请重新输入")
		}
	}
	//从主菜单退出  进入二级菜单
	if key == 1 {
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d \n", &userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s \n", &userPwd)
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("用户注册")
	}
}

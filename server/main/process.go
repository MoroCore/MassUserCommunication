package main

import (
	"MassUserCommunication/common/message"
	processes "MassUserCommunication/server/process"
	"MassUserCommunication/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		up := &processes.UserProcess{
			Coon: this.Conn,
		}
		up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &processes.UserProcess{
			Coon: this.Conn,
		}
		up.ServerPRocessRegister(mes)
	case message.SmsMesType:
		up := &processes.SmsProcess{}
		up.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在,无法处理")
	}
	return nil
}
func (this *Processor) process2() (err error) {

	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		pkg, err := tf.ReadPkg()
		if err == io.EOF {
			fmt.Println("readPkg err = ", err)
			fmt.Println("客户端退出了连接 服务端也退出。。。。")
			return err
		} else if err != nil {
			fmt.Println("readPkg err = ", err)
			return err
		}
		this.serverProcessMes(&pkg)
		if err != nil {
			return err
		}
	}
}

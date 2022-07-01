package model

import (
	"MassUserCommunication/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}

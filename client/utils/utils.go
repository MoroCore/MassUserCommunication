package utils

import (
	"MassUserCommunication/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//传输者
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//如果客户端的conn关闭  read方法不阻塞了
	//循环读取内容    err = io.EOF
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("read pkg head error")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("read pkg boby error ")
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
	}
	return
}
func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	buf := bytes[:4]
	binary.BigEndian.PutUint32(buf, pkgLen)
	n, err := this.Conn.Write(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail ", err)
	}
	_, err = this.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return
}

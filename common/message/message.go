package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnLine = iota
	UserOffLine
	UserBysyStatus
)

//网络上发送的Message消息  序列化
type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //消息的内容
}

type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户pws
	UserName string `json:"userName"` //用户名
}
type LoginResMes struct {
	Code   int    `json:"code"` //注册码  500 200
	UserId []int  `json:"userId"`
	Error  string `json:"error"` //
}
type RegisterMes struct {
	User User
}
type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
type SmsMes struct {
	Context string `json:"context"`
	User           //匿名结构体
}

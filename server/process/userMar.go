package processes

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (this *UserMgr) addOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}
func (this *UserMgr) deleteOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok { //当前用户不在线
		err = fmt.Errorf("用户%d 不存在", userId)
	}
	return
}

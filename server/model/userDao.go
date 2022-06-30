package model

import (
	"MassUserCommunication/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}
func (this *UserDao) getUserById(connn redis.Conn, id int) (user *message.User, err error) {

	res, err := connn.Do("HGet", "users", id)
	if err != nil {
		//没有找到用户
		if err == redis.ErrNil {
			err = ERROR_USER_EXISTS
		}
		return
	}
	user = &message.User{}
	s, err := redis.String(res, err)
	fmt.Println("s", s)
	err = json.Unmarshal([]byte(s), user)
	if err != nil {
		fmt.Println("json.Unmarshal fail", err)
		return
	}
	return
}
func (this *UserDao) LoginCheck(userId int, userPwd string) (user *message.User, err error) {

	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
func (this *UserDao) Register(user *message.User) (err error) {
	con := this.pool.Get()
	defer con.Close()
	_, err = this.getUserById(con, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = con.Do("Hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存用户出错")
		return
	}
	return
}

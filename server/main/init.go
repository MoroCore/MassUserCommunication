package main

import (
	"MassUserCommunication/server/model"
	"github.com/garyburd/redigo/redis"
	"time"
)

//初始化函数

var pool *redis.Pool

func initPool(adress string, maxIdke int, maxActive int, idleTimeOut time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdke,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeOut,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", adress)
		},
	}
}
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

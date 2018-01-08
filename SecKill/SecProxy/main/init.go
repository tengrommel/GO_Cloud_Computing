package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"time"
)

func initRedis() (err error) {
	_ := &redis.Pool{
		MaxIdle:secKillConf.redisConf.redisMaxIdle,
		MaxActive:secKillConf.redisConf.redisMaxActive,
		IdleTimeout: time.Duration(secKillConf.redisConf.redisIdleTimeout)*time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.redisConf.redisAddr)
		},
	}
	return
}

func initEtch() (err error) {
	return
}


func initSec() (err error) {

	err = initRedis()
	if err != nil{
		logs.Error("init redis failed, err: %v", err)
		return
	}

	err = initEtch()
	if err != nil{
		logs.Error("init etcd failed, err: %v", err)
		return
	}

	logs.Info("init sec succ")
	return
}

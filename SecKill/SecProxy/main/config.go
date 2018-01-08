package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
)

var (
	secKillConf = &SecSkillConf{}
)

type SecSkillConf struct {
	redisAddr string
	etcdAddr string
}

func initConfig() (err error) {
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")

	logs.Debug("read config succ, redis addr: v%", redisAddr)
	logs.Debug("read config succ, etcd addr: v%", etcdAddr)

	secKillConf.etcdAddr = etcdAddr
	secKillConf.redisAddr = redisAddr
	if len(redisAddr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return
	}
	return
}
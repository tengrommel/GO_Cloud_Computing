package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
)

var (
	secKillConf = &SecSkillConf{}
)

type RedisConf struct {
	redisAddr string
	redisMaxIdle int
	redisMaxActive int
	redisIdleTimeout int
}

type EtchConf struct {
	etcdAddr string
	timeout int
}

type SecSkillConf struct {
	redisConf RedisConf
	etcdConf EtchConf
}

func initConfig() (err error) {
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")

	logs.Debug("read config succ, redis addr: v%", redisAddr)
	logs.Debug("read config succ, etcd addr: v%", etcdAddr)

	secKillConf.etcdConf.etcdAddr = etcdAddr
	secKillConf.redisConf.redisAddr = redisAddr
	if len(redisAddr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return
	}
	redisMaxIdle, err := beego.AppConfig.Int("redis_max_idle")
	if err != nil{
		err = fmt.Errorf("init config failed, read redis_max_idle error:%v", err)
		return
	}

	redisMaxActive, err := beego.AppConfig.Int("redis_max_active")
	if err != nil{
		err = fmt.Errorf("init config failed, read redis_max_active error:%v", err)
		return
	}

	redisIdleTimeout, err := beego.AppConfig.Int("redis_idle_timeout")
	if err != nil{
		err = fmt.Errorf("init config failed, read redis_idle_timeout error:%v", err)
		return
	}

	secKillConf.redisConf.redisMaxIdle = redisMaxIdle
	secKillConf.redisConf.redisMaxActive = redisMaxActive
	secKillConf.redisConf.redisIdleTimeout = redisIdleTimeout

	etcdTimeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil{
		err = fmt.Errorf("init config failed, read etcd_timeout error:%v", err)
		return
	}

	secKillConf.etcdConf.timeout = etcdTimeout
	return
}
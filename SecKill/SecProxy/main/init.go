package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"time"
	etcd_client "github.com/coreos/etcd/clientv3"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
)

var (
	redisPool *redis.Pool
	etcdClient *etcd_client.Client
)

func initRedis() (err error) {
	redisPool = &redis.Pool{
		MaxIdle:secKillConf.redisConf.redisMaxIdle,
		MaxActive:secKillConf.redisConf.redisMaxActive,
		IdleTimeout: time.Duration(secKillConf.redisConf.redisIdleTimeout)*time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.redisConf.redisAddr)
		},
	}

	conn := redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil{
		 logs.Error("ping redis failed, err:%v\n", err)
		 return
	}

	return
}

func initEtcd() (err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints: []string{secKillConf.etcdConf.etcdAddr},
		DialTimeout: time.Duration(secKillConf.etcdConf.timeout) * time.Second,
	})
	if err != nil{
		logs.Error("connect etcd failed, err:", err)
		return
	}
	etcdClient = cli
	return
}

func converLogLevel(level string) int{
	switch (level) {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func initLogs() (err error) {
	config := make(map[string]interface{})
	config["filename"] = secKillConf.logPath
	config["level"] = converLogLevel(secKillConf.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil{
		fmt.Println("marshal failed, err:", err)
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

func loadSecConf() (err error) {
	key := fmt.Sprintf("%s/product", secKillConf.etcdConf.etcdSecKey)
	resp, err := etcdClient.Get(context.Background(), key)
	if err != nil{
		logs.Error("get [%s] from etcd failed, err :%v",key, err)
		return
	}

	var secProductInfo []SecProductInfoConf
	for k, v := range resp.Kvs{
		logs.Debug("key[%s] value[%s]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil{
			logs.Error("Unmarshal sec product info failed, err%v", err)
			return
		}
		logs.Debug("sec info conf is [%v]", secProductInfo)
	}
	secKillConf.secProductInfo = secProductInfo
	return
}


func initSec() (err error) {
	err = initLogs()
	if err != nil{
		logs.Error("init logger failed, err: %v", err)
	}

	err = initRedis()
	if err != nil{
		logs.Error("init redis failed, err: %v", err)
		return
	}

	err = initEtcd()
	if err != nil{
		logs.Error("init etcd failed, err: %v", err)
		return
	}

	err = loadSecConf()
	if err != nil{
		logs.Error("load sec conf failed, err: %v", err)
		return
	}

	err = initSecProductWatcher()
	if err != nil{
		return
	}

	logs.Info("init sec succ")
	return
}

func initSecProductWatcher () error {
	return nil
}

package main

import (
	"github.com/astaxie/beego"
	_ "github.com/tengrommel/GO_Cloud_Computing/SecKill/SecProxy/router"
)



func main()  {
	err := initConfig()
	if err!= nil{
		panic(err)
		return
	}

	err = initSec()
	if err != nil{
		panic(err)
		return
	}
	beego.Run()
}
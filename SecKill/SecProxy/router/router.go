package router

import (
	"github.com/astaxie/beego"
	"github.com/tengrommel/GO_Cloud_Computing/SecKill/SecProxy/controller"
	"github.com/astaxie/beego/logs"

)

func init()  {
	logs.Debug("init router")
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}

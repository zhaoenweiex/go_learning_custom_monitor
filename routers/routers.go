package routers

import (
	"go_learning_custom_monitor/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/monitor", &controllers.MonitorController{})
}
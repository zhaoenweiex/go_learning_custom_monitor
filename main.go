package main

import (
	_ "go_learning_custom_monitor/routers"
	"github.com/astaxie/beego"
)
func main(){
	beego.SetStaticPath("/","static")
	beego.Run()
}
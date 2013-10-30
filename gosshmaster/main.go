package main

import (
	"github.com/shxsun/gossh/gosshmaster/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}


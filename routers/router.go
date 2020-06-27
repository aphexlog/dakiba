package routers

import (
	"github.com/aphexlog/dakiba/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

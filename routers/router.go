package routers

import (
	"github.com/ruoo/gocms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

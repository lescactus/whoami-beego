package routers

import (
	"github.com/astaxie/beego"
	"github.com/lescactus/whoami-beego/controllers"
)

func init() {
	beego.SetStaticPath("static/css", "static/css")
	beego.SetStaticPath("static/img", "static/img")

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/ip", &controllers.IPController{})
	beego.Router("/port", &controllers.PortController{})
	beego.Router("/lang", &controllers.LangController{})
	beego.Router("/ua", &controllers.UserAgentController{})
	beego.Router("/raw/go", &controllers.RawGoController{})
	beego.Router("/raw/json", &controllers.RawJSONController{})
	beego.Router("/raw/yaml", &controllers.RawYAMLController{})
}

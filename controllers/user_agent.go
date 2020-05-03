package controllers

import (
	"github.com/astaxie/beego"
)

type UserAgentController struct {
	beego.Controller
}

func (c *UserAgentController) Get() {
	c.Ctx.Output.Body([]byte(c.Ctx.Input.Header("User-Agent")))
}

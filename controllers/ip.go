package controllers

import (
	"github.com/astaxie/beego"
)

type IPController struct {
	beego.Controller
}

func (c *IPController) Get() {
	c.Ctx.Output.Body([]byte(c.Ctx.Input.IP()))
}

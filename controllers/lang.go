package controllers

import (
	"github.com/astaxie/beego"
)

type LangController struct {
	beego.Controller
}

func (c *LangController) Get() {
	c.Ctx.Output.Body([]byte(c.Ctx.Input.Header("Accept-Language")))
}

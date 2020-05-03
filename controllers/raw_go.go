package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type RawGoController struct {
	beego.Controller
}

func (c *RawGoController) Get() {
	headers := c.Ctx.Request.Header
	if c.Ctx.Input.Header("Host") != "" {
		headers["Host"] = []string{c.Ctx.Input.Header("Host")}
	}
	c.Ctx.Output.Body([]byte(fmt.Sprintln(headers)))
}

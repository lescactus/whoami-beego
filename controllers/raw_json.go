package controllers

import (
	"encoding/json"
	"regexp"

	"github.com/astaxie/beego"
)

type RawJSONController struct {
	beego.Controller
}

func (c *RawJSONController) Get() {
	headers := c.Ctx.Request.Header
	if c.Ctx.Input.Header("Host") != "" {
		headers["Host"] = []string{c.Ctx.Input.Header("Host")}
	}

	for header := range headers {
		matched, _ := regexp.Match(headersToRemoveRegex, []byte(header))
		if matched {
			delete(headers, header)
		}
	}

	j, err := json.Marshal(headers)
	if err != nil {
		beego.Error("Error while marshling json headers" + err.Error())
		c.Abort("503")
	}

	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.Output.Body([]byte(j))
}

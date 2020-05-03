package controllers

import (
	"regexp"

	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
)

type RawYAMLController struct {
	beego.Controller
}

func (c *RawYAMLController) Get() {
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

	y, err := yaml.Marshal(headers)
	if err != nil {
		beego.Error("Error while marshling yaml headers" + err.Error())
		c.Abort("503")
	}

	c.Ctx.Output.Body([]byte(y))
}

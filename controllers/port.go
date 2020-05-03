package controllers

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type PortController struct {
	beego.Controller
}

func (c *PortController) Get() {
	port, err := getClientRemotePort(c.Ctx.Request.RemoteAddr)
	if err != nil {
		beego.Error("Error while getting client remote port")
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error while getting remote port"))
		return
	}
	c.Ctx.Output.Body([]byte(strconv.Itoa(port)))
}

// Get the client remote port from a string of type http.Request.RemoteAddr
// Port is an integer
func getClientRemotePort(r string) (int, error) {
	_, p, err := net.SplitHostPort(r)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	// Cast port from string to integer
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	return port, nil
}

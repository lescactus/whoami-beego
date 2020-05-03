package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/lescactus/whoami-beego/models"
	"gopkg.in/yaml.v2"
)

var (
	// Remove all 'x-' or 'X-' HTTP headers (Ex: 'X-Forwarded-For')
	headersToRemoveRegex = "(^((forwarded|Forwarded).*)|((x|X)-(.*))$)"

	// GeoIP API URL (https://freegeoip.app/)
	urlGeoIP = "https://freegeoip.app/json/"
)

type IndexController struct {
	beego.Controller
}

func initBrowser(c *IndexController, getLocation bool) (*models.Browser, error) {
	var browser models.Browser
	var err error

	browser.IP = c.Ctx.Input.IP()

	browser.Port, err = getClientRemotePort(c.Ctx.Request.RemoteAddr)
	if err != nil {
		beego.Error("Error while getting client remote port")
		browser.Port = 0
	}

	browser.Host = c.Ctx.Request.Host
	browser.Headers = c.Ctx.Request.Header
	browser.Headers["Host"] = []string{browser.Host}

	for header := range browser.Headers {
		matched, _ := regexp.Match(headersToRemoveRegex, []byte(header))
		if matched {
			delete(browser.Headers, header)
		}
	}

	browser.URL = c.Ctx.Request.URL
	browser.Lang = c.Ctx.Input.Header("Accept-Language")
	browser.UserAgent = c.Ctx.Input.Header("User-Agent")

	j, err := json.Marshal(browser.Headers)
	if err == nil {
		browser.JSON = string(j)
	}

	y, err := yaml.Marshal(browser.Headers)
	if err == nil {
		browser.YAML = string(y)
	}

	if getLocation {
		browser.Location, err = getLocationInfo(browser.IP)
		if err != nil {
			beego.Error("Can't get GeoIP informations for IP %s", browser.IP)
		}
	}
	return &browser, err
}

// From the IP of the client, call the free GeoIP database https://freegeoip.app/json/<ip>
// to get GeoIP infos, such as country name, city name or coordinates.
// Return *Location
func getLocationInfo(ip string) (*models.Location, error) {
	// New http request to https://freegeoip.app/json/<ip>
	req, _ := http.NewRequest("GET", urlGeoIP+ip, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	beego.Info("Calling: %s\n", urlGeoIP+ip)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Unmarshal json response to Location structure
	var location models.Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		return nil, err
	}
	return &location, nil

}

func (c *IndexController) Get() {
	browser, err := initBrowser(c, true)
	if err != nil {
		beego.Error("An error occured while getting browser informations" + err.Error())
		c.Abort("503")
	}
	c.Data["Browser"] = browser
	c.TplName = "index.html"
}

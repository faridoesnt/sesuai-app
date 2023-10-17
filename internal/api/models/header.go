package models

import (
	"Sesuai/pkg/autils"
	"github.com/kataras/iris/v12"
	"strings"
)

type Headers struct {
	User          string
	IP            string
	Header        string
	Params        string
	Endpoint      string
	ID            string
	Authorization string
	Device        string
	OS            string
	OSVersion     string
	DateTime      string
}

func (p *Headers) InitParams(c iris.Context) {
	if !strings.Contains(c.Path(), "v0") {
		result := ""

		for key, values := range c.FormValues() {
			if key != "photo" && key != "image" {
				result += key + ": " + values[0] + ", "
			}
		}

		result = strings.TrimRight(result, ", ")

		p.Params = "{" + result + "}"
	}
}

func (p *Headers) InitHeader(c iris.Context) {
	headers := c.Request().Header

	result := ""
	for key, values := range headers {
		for _, value := range values {
			if key != "Accept" && key != "User-Agent" && key != "Accept-Encoding" && key != "Postman-Token" && key != "Connection" {
				result += key + ": " + value + ", "
			}
		}
	}

	result = strings.TrimRight(result, ", ")

	p.Header = "{" + result + "}"
}

func (p *Headers) GetDateTime() (datetime string) {
	return autils.ParseTo24HourFormat(p.DateTime)
}

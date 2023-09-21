package models

import (
	"Sesuai/pkg/autils"
	"github.com/kataras/iris/v12"
	"strings"
)

type Headers struct {
	IP            string
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

func (p *Headers) GetDateTime() (datetime string) {
	return autils.ParseTo24HourFormat(p.DateTime)
}

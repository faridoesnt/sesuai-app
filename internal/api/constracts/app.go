package constracts

import "github.com/kataras/iris/v12"

type App struct {
	Config      map[string]string
	Datasources *Datasources
	Iris        *iris.Application
	Services    *Services
}

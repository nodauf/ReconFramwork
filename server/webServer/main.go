package main

import (
	"html/template"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/nodauf/ReconFramwork/server/webServer/routers"
)

func main() {
	Run()
}

func Run() {
	beego.AddFuncMap("toJS", toJS)
	beego.Run()

}

func toJS(s string) template.JS {
	return template.JS(s)
}

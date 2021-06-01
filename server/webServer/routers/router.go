package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/nodauf/ReconFramwork/server/webServer/controllers"
	"github.com/nodauf/ReconFramwork/server/webServer/filters"
)

func init() {
	beego.Router("/", &controllers.LoginController{})

	beego.InsertFilter("/*", beego.BeforeRouter, filters.Data)

}

package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/nodauf/ReconFramwork/server/webServer/controllers"
	"github.com/nodauf/ReconFramwork/server/webServer/filters"
)

func init() {
	beego.Router("/", &controllers.LoginController{}, "get,post:Login")

	beego.Router("/recon/", &controllers.ReconController{}, "get:Dashboard")

	beego.Router("/recon/tasks/list", &controllers.ReconController{}, "get:ListTasks")
	beego.Router("/recon/tasks/run/:taskName", &controllers.ReconController{}, "get,post:RunTask")
	beego.Router("/recon/tasks/edit/:taskName", &controllers.ReconController{}, "get,post:EditTask")

	beego.Router("/recon/workflows/list", &controllers.ReconController{}, "get:ListWorkflows")
	beego.Router("/recon/workflows/run/:workflowName", &controllers.ReconController{}, "get,post:RunWorkflow")
	beego.Router("/recon/workflows/edit/:workflowName", &controllers.ReconController{}, "get,post:EditWorkflow")

	beego.InsertFilter("/*", beego.BeforeRouter, filters.Data)
	//beego.InsertFilter("/recon/*", beego.BeforeRouter, filters.IsLogin)

}

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

	beego.Router("/recon/results/overview", &controllers.ReconController{}, "get:OverviewResults")
	beego.Router("/recon/results/listAll", &controllers.ReconController{}, "get:ListAllResults")
	beego.Router("/recon/results/web/list", &controllers.ReconController{}, "get:ListResultsWeb")
	beego.Router("/recon/results/web/details/:ip/:port/:task", &controllers.ReconController{}, "get:DetailsResultsWeb")
	beego.Router("/recon/results/tree/:ip", &controllers.ReconController{}, "get:TreeResults")

	beego.Router("/recon/targets/delete/:target", &controllers.ReconController{}, "get:DeleteTarget")

	beego.InsertFilter("/*", beego.BeforeRouter, filters.Data)
	//beego.InsertFilter("/recon/*", beego.BeforeRouter, filters.IsLogin)

}

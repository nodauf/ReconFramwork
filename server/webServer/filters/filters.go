package filters

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func Data(ctx *context.Context) {
	ctx.Input.SetData("WebsiteName", "ReconFramwork")
}

func IsLogin(ctx *context.Context) {
	if value := ctx.Input.Session("userid"); value == nil {
		ctx.Redirect(302, beego.URLFor("LoginController.Login"))
	}
}

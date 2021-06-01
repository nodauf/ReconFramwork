package filters

import (
	"github.com/beego/beego/v2/server/web/context"
)

func Data(ctx *context.Context) {
	ctx.Input.SetData("WebsiteName", "ReconFramwork")
}

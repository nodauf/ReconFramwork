package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/nodauf/ReconFramwork/server/server/config"
	"github.com/nodauf/ReconFramwork/server/server/db"
)

func (c *ReconController) DeleteTarget() {
	flash := web.ReadFromRequest(&c.Controller)
	targetString := c.Ctx.Input.Param(":target")
	if targetString != "" {
		target := db.GetTarget(targetString)
		if target != nil {
			if db.DeleteTarget(target) {
				flash.Success("Target " + target.GetTarget() + " has been deleted")
				flash.Store(&c.Controller)

			} else {

				flash.Error("Unable to delete the target " + target.GetTarget())
				flash.Store(&c.Controller)
			}
		} else {
			flash.Error("Target " + target.GetTarget() + " not found in the database")
			flash.Store(&c.Controller)
		}
	} else {
		flash.Error("host parameter not found")
		flash.Store(&c.Controller)
	}
	c.Data["Tasks"] = config.Config.Command
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/tasksList.tpl"
	c.Data["DataTables"] = true
}

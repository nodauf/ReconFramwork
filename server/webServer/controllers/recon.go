package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ReconController struct {
	beego.Controller
}

func (c *ReconController) Dashboard() {
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/dashboard.tpl"
}


package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/nodauf/ReconFramwork/server/server/db"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username != "" && password != "" && db.UserExist(username, password) {
		// Do something
	} else {
		c.Data["Error"] = "Wrong credentials"
		c.TplName = "login.tpl"
	}
}

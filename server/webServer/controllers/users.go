package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/nodauf/ReconFramwork/server/server/db"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "login.tpl"
	} else if c.Ctx.Request.Method == "POST" {
		c.ValidateLogin()
	}
}

func (c *LoginController) ValidateLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username != "" && password != "" {
		if userID, exist := db.UserExist(username, password); exist {
			c.SetSession("userid", userID)
			c.Redirect(c.URLFor("ReconController.Dashboard"), 302)
		}
	}
	c.Data["Error"] = "Wrong credentials"
	c.TplName = "login.tpl"

}

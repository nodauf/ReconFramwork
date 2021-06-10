package controllers

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/nodauf/ReconFramwork/server/server/config"
	"github.com/nodauf/ReconFramwork/server/server/orchestrator"
	"gopkg.in/yaml.v3"
)

func (c *ReconController) ListTasks() {
	c.Data["Tasks"] = config.Config.Command
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/tasks/tasksList.tpl"
	c.Data["DataTables"] = true
}

func (c *ReconController) RunTask() {
	flash := web.ReadFromRequest(&c.Controller)
	taskName := c.Ctx.Input.Param(":taskName")
	if taskName != "" {
		if _, ok := config.Config.Command[taskName]; ok {
			c.Data["TaskName"] = taskName
		}
	}
	// Execute the task
	if c.Data["TaskName"] != nil && c.Ctx.Request.Method == "POST" {
		targets := c.GetStrings("targets[]")
		recurseOnSubdomain := c.GetString("recurseOnSubdomain")
		if len(targets) == 0 {
			flash.Error("No target were specified")
			flash.Store(&c.Controller)
		} else {
			for _, target := range targets {
				orchestratorOptions := orchestrator.NewOptions()
				orchestratorOptions.Target = target
				orchestratorOptions.Task = taskName
				if strings.ToLower(recurseOnSubdomain) == "on" {
					orchestratorOptions.RecurseOnSubdomain = true
				}
				orchestratorOptions.Wg.Add(1)
				go orchestratorOptions.RunTask()
				c.Data["Toastr"] = "Task has been successfully launch"
			}
		}

	} else if c.Data["TaskName"] == nil {
		// If the task name in the URL was incorrect or missing
		flash.Error("Task not found")
		flash.Store(&c.Controller)
	}

	c.Data["Select2"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/tasks/taskRun.tpl"
}

func (c *ReconController) EditTask() {
	flash := web.ReadFromRequest(&c.Controller)
	taskName := c.Ctx.Input.Param(":taskName")
	if taskName != "" {
		if _, ok := config.Config.Command[taskName]; ok {
			value, _ := yaml.Marshal(config.Config.Command[taskName])
			c.Data["Yaml"] = string(value)
			c.Data["TaskName"] = taskName
		}
	}
	if c.Data["Yaml"] == nil {
		flash.Error("Task not found")
		flash.Store(&c.Controller)
	}

	c.Data["CodeMirror"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/tasks/taskEdit.tpl"
}

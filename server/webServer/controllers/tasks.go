package controllers

import (
	"strings"

	"github.com/nodauf/ReconFramwork/server/server/config"
	"github.com/nodauf/ReconFramwork/server/server/orchestrator"
	"gopkg.in/yaml.v3"
)

func (c *ReconController) ListTasks() {
	c.Data["Tasks"] = config.Config.Command
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/tasksList.tpl"
	c.Data["DataTables"] = true
}

func (c *ReconController) RunTask() {
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
			c.Data["Error"] = "No target were specified"
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
			}
		}

	} else if c.Data["TaskName"] == nil {
		// If the task name in the URL was incorrect or missing
		c.Data["Error"] = "Task not found"
	}

	c.Data["Select2"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/taskRun.tpl"
}

func (c *ReconController) EditTask() {
	taskName := c.Ctx.Input.Param(":taskName")
	if taskName != "" {
		if _, ok := config.Config.Command[taskName]; ok {
			value, _ := yaml.Marshal(config.Config.Command[taskName])
			c.Data["Yaml"] = string(value)
			c.Data["TaskName"] = taskName
		}
	}
	if c.Data["Yaml"] == nil {
		c.Data["Error"] = "Task not found"
	}

	c.Data["CodeMirror"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/taskEdit.tpl"
}

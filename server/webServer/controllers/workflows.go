package controllers

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/nodauf/ReconFramwork/server/server/config"
	"github.com/nodauf/ReconFramwork/server/server/orchestrator"
	"gopkg.in/yaml.v3"
)

func (c *ReconController) ListWorkflows() {
	c.Data["Workflows"] = config.Config.Workflow
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/workflows/workflowsList.tpl"
	c.Data["DataTables"] = true
}

func (c *ReconController) RunWorkflow() {
	flash := web.ReadFromRequest(&c.Controller)
	workflowName := c.Ctx.Input.Param(":workflowName")
	if workflowName != "" {
		if _, ok := config.Config.Workflow[workflowName]; ok {
			c.Data["WorkflowName"] = workflowName
		}
	}
	// Execute the task
	if c.Data["WorkflowName"] != nil && c.Ctx.Request.Method == "POST" {
		targets := c.GetStrings("targets[]")
		recurseOnSubdomain := c.GetString("recurseOnSubdomain")
		if len(targets) == 0 {
			flash.Error("No target were specified")
			flash.Store(&c.Controller)
		} else {
			for _, target := range targets {
				orchestratorOptions := orchestrator.NewOptions()
				orchestratorOptions.Target = target
				orchestratorOptions.Workflow = workflowName
				if strings.ToLower(recurseOnSubdomain) == "on" {
					orchestratorOptions.RecurseOnSubdomain = true
				}
				orchestratorOptions.Wg.Add(1)
				go orchestratorOptions.RunWorkflow()
				c.Data["Toastr"] = "Workflow has been successfully launch"
			}
		}

	} else if c.Data["WorkflowName"] == nil {
		// If the workflow name in the URL was incorrect or missing
		flash.Error("Workflow not found")
		flash.Store(&c.Controller)
	}

	c.Data["Select2"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/workflows/workflowRun.tpl"
}

func (c *ReconController) EditWorkflow() {
	flash := web.ReadFromRequest(&c.Controller)
	workflowName := c.Ctx.Input.Param(":workflowName")
	if workflowName != "" {
		if _, ok := config.Config.Workflow[workflowName]; ok {
			value, _ := yaml.Marshal(config.Config.Workflow[workflowName])
			c.Data["Yaml"] = string(value)
			c.Data["WorkflowName"] = workflowName
		}
	}
	if c.Data["Yaml"] == nil {
		flash.Error("Workflow not found")
		flash.Store(&c.Controller)
	}
	c.Data["CodeMirror"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/workflows/workflowEdit.tpl"
}

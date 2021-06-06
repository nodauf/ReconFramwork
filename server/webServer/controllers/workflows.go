package controllers

import (
	"strings"

	"github.com/nodauf/ReconFramwork/server/server/config"
	"github.com/nodauf/ReconFramwork/server/server/orchestrator"
	"gopkg.in/yaml.v3"
)

func (c *ReconController) ListWorkflows() {
	c.Data["Workflows"] = config.Config.Workflow
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/workflowsList.tpl"
	c.Data["DataTables"] = true
}

func (c *ReconController) RunWorkflow() {
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
			c.Data["Error"] = "No target were specified"
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
		c.Data["Error"] = "Workflow not found"
	}

	c.Data["Select2"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/workflowRun.tpl"
}

func (c *ReconController) EditWorkflow() {
	workflowName := c.Ctx.Input.Param(":workflowName")
	if workflowName != "" {
		if _, ok := config.Config.Workflow[workflowName]; ok {
			value, _ := yaml.Marshal(config.Config.Workflow[workflowName])
			c.Data["Yaml"] = string(value)
			c.Data["WorkflowName"] = workflowName
		}
	}
	if c.Data["Yaml"] == nil {
		c.Data["Error"] = "Workflow not found"
	}
	c.Data["CodeMirror"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/workflowEdit.tpl"
}

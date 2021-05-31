package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nodauf/ReconFramwork/server/API/models"
)

func RunTask(c *gin.Context) {
	var params models.RunParams
	if err := c.BindJSON(&params); err != nil && params.Task != "" {
		c.AbortWithStatus(400)
		return
	}
	OptionsOrchestrator.Target = params.Target
	OptionsOrchestrator.Task = params.Task
	OptionsOrchestrator.RecurseOnSubdomain = params.Options.RecurseOnSubdomain
	OptionsOrchestrator.RunTask()
}

func RunWorkflow(c *gin.Context) {
	var params models.RunParams
	if err := c.BindJSON(&params); err != nil && params.Workflow != "" {
		c.AbortWithStatus(400)
		return
	}
	OptionsOrchestrator.RunWorkflow()
}

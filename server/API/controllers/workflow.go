package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nodauf/ReconFramwork/server/server/config"
	"gopkg.in/yaml.v2"
)

func ListWorkflow(c *gin.Context) {
	c.JSON(200, config.Config.Workflow)
}

func ViewWorkflow(c *gin.Context) {
	workflowName := c.Param("workflowName")
	yamlByte, _ := yaml.Marshal(config.Config.Workflow[workflowName])
	fmt.Println(yamlByte)
}

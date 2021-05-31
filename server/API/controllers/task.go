package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nodauf/ReconFramwork/server/server/config"
	"gopkg.in/yaml.v2"
)

func ListTask(c *gin.Context) {
	c.JSON(200, config.Config.Command)
}

func ViewTask(c *gin.Context) {
	taskName := c.Param("taskName")
	yamlByte, _ := yaml.Marshal(config.Config.Command[taskName])
	c.String(http.StatusOK, string(yamlByte))
}

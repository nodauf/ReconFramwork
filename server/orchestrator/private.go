package orchestrator

import (
	"strings"

	"github.com/nodauf/ReconFramwork/server/models"
	"github.com/nodauf/ReconFramwork/utils"
)

func preProcessingTemplate(template models.Command, target, service string) string {
	var cmd string
	if service != "" {
		targetAndPort := target
		target = strings.Split(targetAndPort, ":")[0]
		port := strings.Split(targetAndPort, ":")[1]
		cmd = strings.ReplaceAll(template.Service[service], "<target>", target)
		cmd = strings.ReplaceAll(cmd, "<port>", port)
	} else {
		cmd = strings.ReplaceAll(template.Cmd, "<target>", target)
	}
	for variable, value := range template.Variable {
		cmd = strings.ReplaceAll(cmd, "<"+variable+">", value)
	}
	randString := utils.RandomString(10)
	cmd = strings.ReplaceAll(cmd, "<randstring>", randString)
	return cmd
}

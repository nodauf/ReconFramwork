package orchestrator

import (
	"strings"

	"github.com/nodauf/ReconFramwork/server/models"
)

func preProcessingTemplate(template models.Command, target string) string {
	var cmd string
	cmd = strings.ReplaceAll(template.Cmd, "<target>", target)
	for variable, value := range template.Variable {
		cmd = strings.ReplaceAll(cmd, variable, value)
	}
	return cmd
}

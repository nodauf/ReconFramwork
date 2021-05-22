package orchestrator

import (
	"fmt"
	"strings"
	"sync"

	"github.com/RichardKnop/machinery/v1"
	"github.com/nodauf/ReconFramwork/server/models"
	"github.com/nodauf/ReconFramwork/utils"
)

func preProcessingTemplate(template models.Command, target, service string) string {
	var cmd string
	if service != "" {
		targetAndPort := target
		target = strings.Split(targetAndPort, ":")[0]
		port := strings.Split(targetAndPort, ":")[1]
		cmd = strings.ReplaceAll(template.Cmd, "<target>", target)
		cmd = strings.ReplaceAll(cmd, "<port>", port)
		for variable, value := range template.Service[service].Variable {
			cmd = strings.ReplaceAll(cmd, "<"+variable+">", value)
		}
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

func hasService(target models.Target, serviceCommand map[string]models.CommandService) map[string]string {
	return target.HasService(serviceCommand)
}

func (option *Options) recurseOnSubdomain(wg *sync.WaitGroup, server *machinery.Server, taskWorkflowName string, target models.Target, taskOrWorkflow string) {
	if target.HasSubdomain() {
		fmt.Println("recurse on subdomains")
		if taskOrWorkflow == "task" {
			wg.Add(1)
			go option.RunTask(wg, server, taskOrWorkflow, target.GetTarget())
		} else {
			go option.RunWorkflow(wg, server, taskOrWorkflow, target.GetTarget())
		}
	}
}

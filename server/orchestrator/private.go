package orchestrator

import (
	"fmt"
	"strings"
	"sync"

	"github.com/RichardKnop/machinery/v1"
	"github.com/nodauf/ReconFramwork/server/models"
	"github.com/nodauf/ReconFramwork/utils"
)

func preProcessingTemplate(template models.Command, target, service string) (string, string) {
	var cmd string
	var machineryTask string
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
	if template.CustomTask == "" {
		machineryTask = "runCmd"
	} else {
		machineryTask = template.CustomTask
	}

	return cmd, machineryTask
}

func hasService(target models.Target, serviceCommand map[string]models.CommandService) map[string]string {
	return target.HasService(serviceCommand)
}

func (options *Options) recurseOnSubdomain(wg *sync.WaitGroup, server *machinery.Server, taskWorkflowName string, target models.Target, taskOrWorkflow string) {
	if subdomains := target.GetSubdomain(); len(subdomains) > 0 {
		fmt.Println("recurse on subdomains")
		options.RecurseOnSubdomain = false
		for _, subdomain := range subdomains {
			if taskOrWorkflow == "task" {
				wg.Add(1)
				go options.RunTask(wg, server, taskWorkflowName, subdomain)
			} else {
				go options.RunWorkflow(wg, server, taskWorkflowName, subdomain)
			}
		}
	}
}

package orchestrator

import (
	"strings"

	"github.com/nodauf/ReconFramwork/server/server/db"
	"github.com/nodauf/ReconFramwork/server/server/models"
	modelsConfig "github.com/nodauf/ReconFramwork/server/server/models/config"
	"github.com/nodauf/ReconFramwork/server/server/utils"
)

func preProcessingTemplate(template modelsConfig.Command, target, service string) (string, string) {
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

func (options Options) recurseOnSubdomain(target models.Target, taskOrWorkflow string) {

	if subdomains := db.GetSubdomain(target); len(subdomains) > 0 {
		options.RecurseOnSubdomain = false
		for _, subdomain := range subdomains {
			if taskOrWorkflow == "task" {
				options.Wg.Add(1)
				var optionsSubdomain Options
				optionsSubdomain = options
				optionsSubdomain.Target = subdomain
				go optionsSubdomain.RunTask()
			} else {

				var optionsSubdomain Options
				optionsSubdomain = options
				optionsSubdomain.Target = subdomain
				go optionsSubdomain.RunWorkflow()
			}
		}
	}
}

func (options Options) runOnAllDomains(target models.Target) {
	if domains := target.GetDomain(); len(domains) > 0 {
		options.RunOnAllDomains = false
		for _, domain := range domains {
			options.Wg.Add(1)
			var optionsSubdomain Options
			optionsSubdomain = options
			optionsSubdomain.Target = domain
			go optionsSubdomain.RunTask()
		}
	}
}

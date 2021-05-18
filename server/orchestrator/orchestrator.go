package orchestrator

import (
	"sync"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/config"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/models/database"
	"github.com/nodauf/ReconFramwork/utils"
)

func RunTask(wg *sync.WaitGroup, server *machinery.Server, task, target string) {
	defer wg.Done()
	targetType := utils.ParseList(config.Config.Command[task].Target)
	// service is part of the targets
	_, targetServiceConfig := utils.StringInSlice("service", targetType)

	// If the destination is a network
	if utils.IsNetwork(target) {
		hosts, err := utils.HostsFromNetwork(target)
		if err != nil {
			log.ERROR.Println(err)
			return
		}
		// Execute the command for each host of the network
		for _, host := range hosts {
			RunTask(wg, server, task, host)
		}

	} else {
		host := db.GetHost(target)
		if host.Address == "" {
			var host database.Host
			if utils.IsIP(target) {
				host.Address = target
			} else {
				host.Hostname = target
			}
			db.AddOrUpdateHost(host)
		}
		// If the target is the all host no need to specified port or service
		if _, ok := utils.StringInSlice("host", targetType); ok {
			//cmd := strings.ReplaceAll(config.Config.Command[task].Cmd, "<target>", target)
			parser := config.Config.Command[task].ParserFunction
			cmd := preProcessingTemplate(config.Config.Command[task], target, "")
			executeCommands(server, target, cmd, parser, task)
			// If the target is a service and the host has the service in the database from a previous scan
		} else if targetServiceDB := db.HostHasService(target, config.Config.Command[task].Service); targetServiceConfig && len(targetServiceDB) > 0 {
			parser := config.Config.Command[task].ParserFunction
			for service, targetAndPort := range targetServiceDB {
				cmd := preProcessingTemplate(config.Config.Command[task], targetAndPort, service)
				executeCommands(server, target, cmd, parser, task)
			}
			// If the template target the domain, for exemple subdomain enumeration
		} else if _, ok := utils.StringInSlice("domain", targetType); ok {
			parser := config.Config.Command[task].ParserFunction
			cmd := preProcessingTemplate(config.Config.Command[task], target, "")
			executeCommands(server, target, cmd, parser, task)
		} else {
			log.DEBUG.Println(task)
			log.DEBUG.Println(config.Config.Command)
			log.DEBUG.Println(targetServiceDB)
			log.ERROR.Println("Impossible to execute the task. The host is not found or the host has not the service targeted")
		}
	}
}

func RunWorkflow(wg *sync.WaitGroup, server *machinery.Server, workflowString, target string) {
	defer wg.Done()
	// If the destination is a network
	if utils.IsNetwork(target) {
		hosts, err := utils.HostsFromNetwork(target)
		if err != nil {
			log.ERROR.Println(err)
			return
		}
		// Execute the command for each host of the network
		for _, host := range hosts {
			wg.Add(1)
			go RunWorkflow(wg, server, workflowString, host)
		}

		// If the target is the all host no need to specified port or service
	} else {
		workflow, ok := config.Config.Workflow[workflowString]
		if ok {
			//fmt.Println(config.Config.Workflow)
			for _, task := range workflow.Commands {
				wg.Add(1)
				RunTask(wg, server, task, target)
			}
		} else {
			log.ERROR.Println("Workflow " + workflowString + " not found")
		}
	}
}

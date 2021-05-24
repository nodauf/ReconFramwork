package orchestrator

import (
	"sync"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/config"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/models"
	"github.com/nodauf/ReconFramwork/utils"
)

type Options struct {
	RecurseOnSubdomain bool
}

func (options Options) RunTask(wg *sync.WaitGroup, server *machinery.Server, task, target string) {
	defer wg.Done()
	log.DEBUG.Println("Running task " + task + " over target " + target)
	var targetObject models.Target
	targetType := utils.ParseList(config.Config.Command[task].Target)
	// service is part of the targets, to know if we can use this template for this host
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
			options.RunTask(wg, server, task, host)
		}

	} else {
		targetObject = db.AddOrUpdateTarget(target)
		//we process for subdomain
		if options.RecurseOnSubdomain {
			options.recurseOnSubdomain(wg, server, task, targetObject, "task")
		}
		// If the target is the all host no need to specified port or service
		if _, ok := utils.StringInSlice("host", targetType); ok {
			//cmd := strings.ReplaceAll(config.Config.Command[task].Cmd, "<target>", target)
			parser := config.Config.Command[task].ParserFunction
			cmd := preProcessingTemplate(config.Config.Command[task], target, "")
			executeCommands(server, target, cmd, parser, task)
			// If the target is a service and the host has the service in the database from a previous scan
			//} else if targetServiceDB := db.HostHasService(target, config.Config.Command[task].Service); targetServiceConfig && len(targetServiceDB) > 0 {
		} else if targetServiceDB := hasService(targetObject, config.Config.Command[task].Service); targetServiceConfig && len(targetServiceDB) > 0 {
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
			//log.DEBUG.Println(config.Config.Command)
			log.DEBUG.Println(targetServiceDB)
			log.ERROR.Println("Impossible to execute the task " + task + ". The host is not found or the host has not the service targeted")
		}

	}
}

func (options Options) RunWorkflow(wg *sync.WaitGroup, server *machinery.Server, workflowString, target string) {
	//defer wg.Done()

	// If the destination is a network
	if utils.IsNetwork(target) {
		hosts, err := utils.HostsFromNetwork(target)
		if err != nil {
			log.ERROR.Println(err)
			return
		}
		// Execute the command for each host of the network
		for _, host := range hosts {
			//wg.Add(1)
			go options.RunWorkflow(wg, server, workflowString, host)
		}

		// If the target is the all host no need to specified port or service
	} else {
		targetObject := db.AddOrUpdateTarget(target)
		workflow, ok := config.Config.Workflow[workflowString]
		if ok {
			//fmt.Println(config.Config.Workflow)
			for _, task := range workflow.Commands {
				log.INFO.Println("Running task: " + task)
				// wg.Done is done at the end of each tasks
				wg.Add(1)
				if workflow.Options.ParallelizeTasks {
					go options.RunTask(wg, server, task, target)
				} else {
					options.RunTask(wg, server, task, target)
				}
			}
		} else {
			log.ERROR.Println("Workflow " + workflowString + " not found")
		}
		//Now we have process for this host we process for subdomain

		if options.RecurseOnSubdomain {
			options.recurseOnSubdomain(wg, server, workflowString, targetObject, "workflow")
		}
	}
}

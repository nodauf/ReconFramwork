package orchestrator

import (
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/config"
	"github.com/nodauf/ReconFramwork/server/server/db"
	"github.com/nodauf/ReconFramwork/server/server/models"
	"github.com/nodauf/ReconFramwork/utils"
)

func (options Options) RunTask() {
	defer options.Wg.Done()
	log.DEBUG.Println("Running task " + options.Task + " over target " + options.Target)

	targetType := utils.ParseList(config.Config.Command[options.Task].Target)
	// service is part of the targets, to know if we can use this template for this host
	_, targetServiceConfig := utils.StringInSlice("service", targetType)

	// If the destination is a network
	if utils.IsNetwork(options.Target) {
		hosts, err := utils.HostsFromNetwork(options.Target)
		if err != nil {
			log.ERROR.Println(err)
			return
		}
		// Execute the command for each host of the network
		for _, host := range hosts {
			var optionsHost Options
			optionsHost = options
			optionsHost.Target = host
			options.Wg.Add(1)
			go optionsHost.RunTask()
		}

	} else {
		targetObject := models.CreateTarget(options.Target)
		targetObject = db.AddOrUpdateTarget(targetObject)
		//we process for subdomain
		if options.RecurseOnSubdomain {
			options.recurseOnSubdomain(targetObject, "task")
		}
		if options.RunOnAllDomains {
			options.runOnAllDomains(targetObject)
		}
		// If the target is the all host no need to specified port or service
		if _, ok := utils.StringInSlice("host", targetType); ok {
			//cmd := strings.ReplaceAll(config.Config.Command[task].Cmd, "<target>", target)
			parser := config.Config.Command[options.Task].ParserFunction
			cmd, machineryTask := preProcessingTemplate(config.Config.Command[options.Task], options.Target, "")
			executeCommands(options.Server, options.Target, cmd, parser, options.Task, machineryTask)
			// If the target is a service and the host has the service in the database from a previous scan
			//} else if targetServiceDB := db.HostHasService(target, config.Config.Command[task].Service); targetServiceConfig && len(targetServiceDB) > 0 {
		} else if targetServiceDB := targetObject.HasService(config.Config.Command[options.Task].Service); targetServiceConfig && len(targetServiceDB) > 0 {
			parser := config.Config.Command[options.Task].ParserFunction
			for service, targetAndPort := range targetServiceDB {
				cmd, machineryTask := preProcessingTemplate(config.Config.Command[options.Task], targetAndPort, service)
				executeCommands(options.Server, options.Target, cmd, parser, options.Task, machineryTask)
			}
			// If the template target the domain, for exemple subdomain enumeration
		} else if _, ok := utils.StringInSlice("domain", targetType); ok {
			parser := config.Config.Command[options.Task].ParserFunction
			cmd, machineryTask := preProcessingTemplate(config.Config.Command[options.Task], options.Target, "")
			executeCommands(options.Server, options.Target, cmd, parser, options.Task, machineryTask)
		} else {
			log.DEBUG.Println(options.Task)
			log.DEBUG.Println(targetType)
			//log.DEBUG.Println(config.Config.Command)
			log.DEBUG.Println(targetServiceDB)
			log.ERROR.Println("Impossible to execute the task " + options.Task + ". The host " + options.Target + " has not the service targeted or the task does not exist")
		}

	}
	log.DEBUG.Println("Task " + options.Task + " pour la target " + options.Target + " fini")
}

func (options Options) RunWorkflow() {
	defer options.Wg.Done()

	// If the destination is a network
	if utils.IsNetwork(options.Target) {
		hosts, err := utils.HostsFromNetwork(options.Target)
		if err != nil {
			log.ERROR.Println(err)
			return
		}
		// Execute the command for each host of the network
		for _, host := range hosts {
			//wg.Add(1)
			var optionsHost Options
			optionsHost = options
			optionsHost.Target = host
			go optionsHost.RunWorkflow()
		}

		// If the target is the all host no need to specified port or service
	} else {

		targetObject := models.CreateTarget(options.Target)
		targetObject = db.AddOrUpdateTarget(targetObject)
		workflow, ok := config.Config.Workflow[options.Workflow]
		if ok {
			//fmt.Println(config.Config.Workflow)
			for _, task := range workflow.Commands {
				log.INFO.Println("Running task: " + task)
				// wg.Done is done at the end of each tasks
				options.Wg.Add(1)
				options.Task = task
				if workflow.Options.ParallelizeTasks {
					go options.RunTask()
				} else {
					options.RunTask()
				}
			}
		} else {
			log.ERROR.Println("Workflow " + options.Workflow + " not found")
		}
		//Now we have process for this host we process for subdomain

		if options.RecurseOnSubdomain {
			options.recurseOnSubdomain(targetObject, "workflow")
		}
	}
}

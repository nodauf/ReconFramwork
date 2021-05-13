package orchestrator

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/config"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/utils"
)

func RunTask(server *machinery.Server, task, target string) {
	targetType := utils.ParseList(config.Config.Command[task].Target)
	_, targetServiceConfig := utils.StringInSlice("service", targetType)

	// If the destination is a network
	if utils.IsNetwork(target) {
		hosts, err := utils.HostsFromNetwork(target)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Execute the command for each host of the network
		for _, host := range hosts {
			RunTask(server, task, host)
		}

		// If the target is the all host no need to specified port or service
	} else if _, ok := utils.StringInSlice("host", targetType); ok {
		//cmd := strings.ReplaceAll(config.Config.Command[task].Cmd, "<target>", target)
		parser := config.Config.Command[task].ParserFunction
		cmd := preProcessingTemplate(config.Config.Command[task], target)
		ExecuteCommands(server, cmd, parser)
		// If the target is a service and the host has the service in the database from a previous scan
	} else if targetServiceDB := db.HostHasService(target, config.Config.Command[task].Service); targetServiceConfig && len(targetServiceDB) > 0 {
		parser := config.Config.Command[task].ParserFunction
		for _, target := range targetServiceDB {
			cmd := preProcessingTemplate(config.Config.Command[task], target)
			ExecuteCommands(server, cmd, parser)
		}
	} else {
		fmt.Println(targetServiceDB)
		fmt.Println(targetServiceConfig)
		log.ERROR.Println("Impossible to execute the task")
	}
}

func RunWorkflow(workflow, target string) {

}

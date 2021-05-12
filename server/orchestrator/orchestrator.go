package orchestrator

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/nodauf/ReconFramwork/server/config"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/utils"
)

func RunTask(server *machinery.Server, task, target string) {
	targetType := utils.ParseList(config.Config.Command[task].Target)
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
	} else if utils.StringInSlice("host", targetType) {
		//cmd := strings.ReplaceAll(config.Config.Command[task].Cmd, "<target>", target)
		parser := config.Config.Command[task].ParserFunction
		cmd := preProcessingTemplate(config.Config.Command[task], target)
		ExecuteCommands(server, cmd, parser)

	} else if targetService := db.HostHasService(target, config.Config.Command[task].Service); utils.StringInSlice("service", targetType) && len(targetService) > 0 {
		parser := config.Config.Command[task].ParserFunction
		for _, target := range targetService {
			cmd := preProcessingTemplate(config.Config.Command[task], target)
			ExecuteCommands(server, cmd, parser)
		}
	}
}

func RunWorkflow(workflow, target string) {

}

package prompt

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/RichardKnop/machinery/v1/log"
	prompt "github.com/c-bata/go-prompt"
	"github.com/nodauf/ReconFramwork/server/config"
	"github.com/nodauf/ReconFramwork/server/orchestrator"
	"github.com/nodauf/ReconFramwork/utils"
)

var optionsOrchestrator orchestrator.Options

func handleExit() {
	// workaround for the bug https://github.com/c-bata/go-prompt/issues/147
	rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()
}

func executor(in string) {
	command := strings.Split(in, " ")
	first := command[0]
	switch strings.ToLower(first) {
	case "show":
		optionsOrchestrator.PrintOptions()
	case "set":
		if len(command) > 1 {
			second := command[1]
			value := command[2]
			switch strings.ToLower(second) {
			case "target":
				optionsOrchestrator.Target = value
			case "task":
				optionsOrchestrator.Task = value
			case "workflow":
				optionsOrchestrator.Workflow = value
			case "recurseonsubdomain":
				valueBool, err := strconv.ParseBool(value)
				if err == nil {
					optionsOrchestrator.RecurseOnSubdomain = valueBool
				} else {
					log.ERROR.Println("Can't convert " + value + " to boolean")
				}
			}
		} else {
			//Help for command set
		}
	case "list":
		if len(command) > 1 {
			second := command[1]
			switch strings.ToLower(second) {
			case "tasks":
				config.PrintTasks()
			case "workflows":
				config.PrintWorkflows()
			}
		} else {
			//Help for command set
		}
	case "run":
		if len(command) > 1 {
			second := command[1]
			switch strings.ToLower(second) {
			case "task":
				optionsOrchestrator.Wg.Add(1)
				go optionsOrchestrator.RunTask()
			case "workflow":
				optionsOrchestrator.Wg.Add(1)
				go optionsOrchestrator.RunWorkflow()
			}
		} else {
			//Help for command run
		}
	case "users":
		if len(command) > 1 {
			second := command[1]
			switch strings.ToLower(second) {
			case "add":
			case "delete":
			}
		} else {
			//Help for command user
		}
	case "search":
		if len(command) > 1 {
			config.SearchTasks(command[1])
		} else {
			//Help for command set
		}
	case "help":
	case "exit":
		handleExit()
		os.Exit(0)
	// Not doing anything for just a new line
	case "":
	default:
		fmt.Println("Invalid command")
	}
}

// Prompt run the custom prompt to manage sessions
func Prompt() {
	defer handleExit()
	log.INFO.Println("Starting the server")
	server, err := utils.GetMachineryServer()
	if err != nil {
		log.ERROR.Fatalln(err)
	}
	var wg sync.WaitGroup

	optionsOrchestrator.Server = server
	optionsOrchestrator.Wg = &wg

	p := prompt.New(
		executor,
		complete,
		prompt.OptionPrefix("ReconFramwork> "),
		prompt.OptionPrefixTextColor(prompt.Red),
		prompt.OptionTitle("ReconFramwork"),
	)
	p.Run()
	wg.Wait()
}

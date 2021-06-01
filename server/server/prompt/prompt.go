package prompt

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	prompt "github.com/c-bata/go-prompt"
	"github.com/nodauf/ReconFramwork/server/server/config"
	"github.com/nodauf/ReconFramwork/server/server/db"
	"github.com/nodauf/ReconFramwork/server/server/orchestrator"
)

var optionsOrchestrator orchestrator.Options

func handleExit() {
	// workaround for the bug https://github.com/c-bata/go-prompt/issues/147
	rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()
	optionsOrchestrator.Wg.Done()
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
				if len(command) == 4 {
					username := command[2]
					username = strings.ToLower(username)
					password := command[3]
					db.AddUser(username, password)
				} else {
					// Help for add user
				}
			case "delete":
				if len(command) == 3 {
					username := command[2]
					username = strings.ToLower(username)
					db.DeleteUser(username)
				} else {
					// Help for delete user
				}
			}
		} else {
			//Help for command user and show all users
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
func Prompt(options orchestrator.Options) {
	optionsOrchestrator = options
	defer handleExit()
	log.INFO.Println("Starting the server")

	p := prompt.New(
		executor,
		complete,
		prompt.OptionPrefix("ReconFramwork> "),
		prompt.OptionPrefixTextColor(prompt.Red),
		prompt.OptionTitle("ReconFramwork"),
	)
	p.Run()
	/*optionsOrchestrator.Target = "localhost"
	optionsOrchestrator.Task = "nmap-quick-scan"
	optionsOrchestrator.Wg.Add(1)
	go optionsOrchestrator.RunTask()*/
}

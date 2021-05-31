package prompt

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/nodauf/ReconFramwork/server/server/config"
)

var commands = []prompt.Suggest{
	{Text: "show", Description: "Show the variable set with set command"},
	{Text: "set", Description: "Set a value to a variable "},
	{Text: "list", Description: "List tasks and workflows"},
	{Text: "run", Description: "Run the task or workflow against the target defined with set command"},
	{Text: "users", Description: "Manage users (default print them)"},
	{Text: "search", Description: "Search among the tasks"},
	{Text: "help", Description: "Help menu"},

	{Text: "exit", Description: "Exit this program"},
}

// set subcommand

var setSubCommand = []prompt.Suggest{
	{Text: "target", Description: "Set target variable"},
	{Text: "task", Description: "Set task variable"},
	{Text: "workflow", Description: "Set workflow variable"},
	{Text: "recurseOnSubdomain", Description: "Manage recurseOnSubdomain option"},
}

// list subcommand

var listSubCommand = []prompt.Suggest{
	{Text: "tasks", Description: "List tasks"},
	{Text: "workflows", Description: "List workflows"},
}

var tasksListSubCommand = []prompt.Suggest{}

var workflowsListSubCommand = []prompt.Suggest{}

// run subcommand

var runSubCommand = []prompt.Suggest{
	{Text: "task", Description: "Run task"},
	{Text: "workflow", Description: "Run workflow"},
}

// users subcommand

var usersSubCommand = []prompt.Suggest{
	{Text: "add", Description: "Add a new user"},
	{Text: "delete", Description: "Delete a user"},
}

// Help subcommand

var helpSubCommand = []prompt.Suggest{}

func LoadCompleter() {
	for taskName, command := range config.Config.Command {
		tasksListSubCommand = append(tasksListSubCommand, prompt.Suggest{Text: taskName, Description: command.Description})
	}
	for workflowName, workflow := range config.Config.Workflow {
		workflowsListSubCommand = append(workflowsListSubCommand, prompt.Suggest{Text: workflowName, Description: workflow.Description})
	}
}

func complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")

	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}
	first := strings.ToLower(args[0])
	switch first {
	case "set":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(setSubCommand, second, true)
		} else if len(args) > 2 {
			switch second {
			case "task":
				third := args[2]
				if len(args) == 3 {
					return prompt.FilterHasPrefix(tasksListSubCommand, third, true)
				}
			case "workflow":
				third := args[2]
				if len(args) == 3 {
					return prompt.FilterHasPrefix(workflowsListSubCommand, third, true)
				}
			}
		}
	case "list":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(listSubCommand, second, true)
		}
	case "run":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(runSubCommand, second, true)
		}
	case "users":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(usersSubCommand, second, true)
		}
	case "help":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(helpSubCommand, second, true)
		}
	}
	return []prompt.Suggest{}
}

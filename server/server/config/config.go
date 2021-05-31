package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/fatih/color"
	"github.com/karrick/godirwalk"
	"gopkg.in/yaml.v2"

	modelsConfig "github.com/nodauf/ReconFramwork/server/server/models/config"
)

var Config struct {
	Command  map[string]modelsConfig.Command
	Workflow map[string]modelsConfig.Workflow
}

func init() {
	Config.Command = make(map[string]modelsConfig.Command)
	Config.Workflow = make(map[string]modelsConfig.Workflow)
}

func SearchTasks(searchString string) {
	for taskName, command := range Config.Command {
		if strings.Contains(taskName, searchString) || strings.Contains(command.Description, searchString) {
			fmt.Println(" - " + color.YellowString(taskName) + " => " + command.Description)
		}
	}
}

func LoadConfig() {
	err := getTemplateFiles("./server/config/")
	if err != nil {
		log.FATAL.Fatalln(err)
	}

}

func PrintTasks() {
	for taskName, command := range Config.Command {
		fmt.Println(" - " + color.YellowString(taskName) + " => " + command.Description)
	}
}

func PrintWorkflows() {
	for workflowName, workflow := range Config.Workflow {
		fmt.Println(" - " + color.YellowString(workflowName) + " => " + workflow.Description)
	}
}

func getTemplateFiles(filePath string) error {
	err := godirwalk.Walk(filePath, &godirwalk.Options{
		Callback: func(filePath string, de *godirwalk.Dirent) error {
			var err error
			if !de.IsDir() && strings.HasSuffix(filePath, ".yaml") {
				// if parent directory is commands or workflow
				pathDir := path.Dir(filePath)
				if strings.Contains(pathDir, "commands") {
					err = loadTemplateCommands(filePath)
				} else if strings.Contains(pathDir, "workflow") {
					err = loadTemplateWorkflows(filePath)
				}
			}

			return err
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})
	return err
}

func loadTemplateCommands(filePath string) error {
	var command = &modelsConfig.Command{}
	//command.Parser = reflect.ValueOf("parsers.ParseSmbMap").Interface().(models.Parser)
	command.Name = "test"
	log.INFO.Println("Loading template " + filePath)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	//fmt.Println(string(data))
	err = yaml.NewDecoder(bytes.NewReader(data)).Decode(command)
	if err != nil {
		return err
	}
	Config.Command[command.Name] = *command
	//Config.Command = append(Config.Command, *command)
	//fmt.Printf("%#v \n", Config.Command)
	return nil
}

func loadTemplateWorkflows(filePath string) error {
	var workflow = &modelsConfig.Workflow{}
	log.INFO.Println("Loading workflow " + filePath)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	//fmt.Println(string(data))
	err = yaml.NewDecoder(bytes.NewReader(data)).Decode(workflow)
	if err != nil {
		return err
	}

	Config.Workflow[workflow.Name] = *workflow
	//Config.Workflow = append(Config.Workflow, *workflow)
	return nil
}

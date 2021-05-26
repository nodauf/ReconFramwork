package main

import (
	machinery "github.com/RichardKnop/machinery/v1"

	"github.com/nodauf/ReconFramwork/server/config"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/prompt"
)

var taskserver *machinery.Server

func init() {
	config.LoadConfig()
	prompt.LoadCompleter()

	db.Init()
}

func main() {
	/*var wg sync.WaitGroup
	var options orchestrator.Options
	//options.RecurseOnSubdomain = true
	target := "audon.fr"
	task := "ffuf"
	//task := "Domain from certificat"
	workflow := "discovery http"
	db.Init()
	/*var host database.Host
	host.Address = "127.0.0.1"
	var port database.Port
	port.Port = 22
	var portComment database.PortComment
	portComment.Comment = "test ssh"
	port.PortComment = append(port.PortComment, portComment)
	host.Ports = append(host.Ports, port)
	port.Port = 123
	host.Ports = append(host.Ports, port)

	db.AddOrUpdateHost(host)

	var host2 database.Host
	host2 = db.GetHost("127.0.0.1")
	empJSON, _ := json.MarshalIndent(host2, "", "  ")
	fmt.Println(string(empJSON))
	host2.Ports[0].Port = 111

	db.AddOrUpdateHost(host2)*/
	/*	log.INFO.Println("Starting the server")
		server, err := utils.GetMachineryServer()
		if err != nil {
			log.ERROR.Fatalln(err)
		}
		wg.Add(1)
		options.Target = target
		options.Wg = &wg
		options.Server = server
		options.Task = task
		options.Workflow = workflow
		go options.RunTask()
		//go options.ConsumeEndedTasks(server, &wg)
		//go options.RunWorkflow(&wg, server, workflow, target)

		wg.Wait()*/
	//orchestrator.ExecuteCommands(server, task, target)
	prompt.Prompt()
}

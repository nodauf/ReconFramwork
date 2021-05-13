package main

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"

	"github.com/nodauf/ReconFramwork/server/config"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/orchestrator"
	"github.com/nodauf/ReconFramwork/utils"
)

var taskserver *machinery.Server

func init() {
	config.LoadConfig()
}

func main() {
	target := "127.0.0.1"
	//task := "nmap quick scan"
	//task := "ffuf"
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
	log.INFO.Println("Starting the server")
	server, err := utils.GetMachineryServer()
	if err != nil {
		log.ERROR.Fatalln(err)
	}
	//orchestrator.RunTask(server, task, target)
	orchestrator.RunWorkflow(server, workflow, target)
	//orchestrator.ExecuteCommands(server, task, target)

}

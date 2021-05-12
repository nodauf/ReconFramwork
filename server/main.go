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
	task := "nmap quick scan"
	db.Init()

	log.INFO.Println("Starting the server")
	server, err := utils.GetMachineryServer()
	if err != nil {
		log.ERROR.Fatalln(err)
	}
	orchestrator.RunTask(server, task, target)
	//orchestrator.ExecuteCommands(server, task, target)

}

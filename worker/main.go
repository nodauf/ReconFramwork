package main

import (
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/utils"
)

func main() {
	log.INFO.Println("Starting the worker")
	server, err := utils.GetMachineryServer()
	if err != nil {
		log.ERROR.Fatalln(err)
	}

	worker := server.NewWorker("machinery_worker", 10)
	if err := worker.Launch(); err != nil {
		log.ERROR.Fatalln(err)
	}
}

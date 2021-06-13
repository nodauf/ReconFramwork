package utils

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/tasks"
	customTasks "github.com/nodauf/ReconFramwork/tasks/custom"
)

func GetMachineryServer() (*machinery.Server, error) {
	log.INFO.Println("initing task server")

	server, err := machinery.NewServer(&config.Config{
		Broker:        "redis://localhost:6379",
		ResultBackend: "redis://localhost:6379",
	})
	if err == nil {

		err = server.RegisterTasks(map[string]interface{}{
			"runCmd":         tasks.RunCmd,
			"DomainFromCert": customTasks.DomainFromCert,
		})
	}
	return server, err

}

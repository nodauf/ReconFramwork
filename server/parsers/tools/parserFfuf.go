package parsersTools

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/models/parsers"
	"github.com/nodauf/ReconFramwork/utils"
)

func (parse Parser) ParseFfuf(taskName, cmdline, stdout, stderr string) bool {
	var ffuf modelsParsers.Ffuf
	// Slice to store all the port discover (port in the db) for this host
	var portList []string
	err := json.Unmarshal([]byte(stdout), &ffuf)

	// If there is something found by ffuf
	if err == nil && len(ffuf.Results) > 0 {
		commandline := strings.Split(ffuf.Commandline, "/")[2]
		target := strings.Split(commandline, ":")[0]
		port := strings.Split(commandline, ":")[1]
		if host := db.GetHost(target); host.Address != "" {
			for _, portHost := range host.Ports {
				portList = append(portList, strconv.Itoa(portHost.Port))
			}

			if index, ok := utils.StringInSlice(port, portList); ok {
				// Update the object from the database
				//Convert to ffuf to json, it will contains only the necessary fields
				commandOutput, _ := json.Marshal(ffuf)
				portComment := modelsDatabases.PortComment{Task: taskName, CommandOutput: string(commandOutput)}
				host.Ports[index].PortComment = append(host.Ports[index].PortComment, portComment)
				db.AddOrUpdateHost(&host)
			} else {
				log.ERROR.Println("Something went wrong. Port " + port + " not found for the Host " + target + " in the database.")
			}

		} else {
			log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
		}
	}
	return true
}

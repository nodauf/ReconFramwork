package parsers

import (
	"encoding/json"
	"strconv"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/models/parsers"
	"github.com/nodauf/ReconFramwork/utils"
)

func (parse Parser) ParseNikto(taskName, cmdline, stdout, stderr string) bool {
	var nikto modelsParsers.Nikto
	// Slice to store all the port discover (port in the db) for this host
	var portList []string
	err := json.Unmarshal([]byte(stdout), &nikto)

	if err == nil && len(nikto.Vulnerabilities) > 0 {
		target := nikto.IP
		port := nikto.Port
		if host := db.GetHost(target); host.Address != "" {
			for _, portHost := range host.Ports {
				portList = append(portList, strconv.Itoa(portHost.Port))
			}

			if index, ok := utils.StringInSlice(port, portList); ok {
				// Update the object from the database
				//Convert to nikto to json, it will contains only the necessary fields
				commandOutput, _ := json.Marshal(nikto)
				portComment := database.PortComment{Task: taskName, CommandOutput: string(commandOutput)}
				host.Ports[index].PortComment = append(host.Ports[index].PortComment, portComment)
				db.AddOrUpdateHost(host)
			} else {
				log.ERROR.Println("Something went wrong. Port " + port + " not found for the Host " + target + " in the database.")
			}

		} else {
			log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
		}
	} else {
		log.INFO.Println("Nothing found by nikto")
	}
	return true
}

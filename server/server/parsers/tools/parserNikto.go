package parsersTools

import (
	"encoding/json"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/server/models/parsers"
)

func (parse Parser) ParseNikto(taskName, cmdline, stdout, stderr string) bool {
	var nikto modelsParsers.Nikto
	err := json.Unmarshal([]byte(stdout), &nikto)

	if err == nil && len(nikto.Vulnerabilities) > 0 {
		target := nikto.IP
		port := nikto.Port
		if host := db.GetHostWherePort(target, port); host.Address != "" {
			// Can only have on port with this value
			index := 0
			// Update the object from the database
			//Convert to nikto to json, it will contains only the necessary fields
			commandOutput, _ := json.Marshal(nikto)
			portComment := modelsDatabases.PortComment{Task: taskName, CommandOutput: string(commandOutput)}
			host.Ports[index].PortComment = append(host.Ports[index].PortComment, portComment)
			db.AddOrUpdateHost(&host)

		} else {
			log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
		}
	} else {
		log.INFO.Println("Nothing found by nikto")
	}
	return true
}

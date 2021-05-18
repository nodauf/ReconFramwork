package parsers

import (
	"encoding/json"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/models/parsers"
)

func (parse Parser) ParseNuclei(taskName, cmdline, stdout, stderr string) bool {
	var nuclei modelsParsers.Nuclei
	var nucleiFinding modelsParsers.NucleiFinding
	for _, finding := range strings.Split(stdout, "\n") {

		err := json.Unmarshal([]byte(finding), &nucleiFinding)
		if err == nil {
			nuclei.Findings = append(nuclei.Findings, nucleiFinding)
		}
	}
	if len(nuclei.Findings) > 0 {
		target := nuclei.Findings[0].IP
		port := strings.Split(nuclei.Findings[0].Host, ":")[2]
		if host := db.GetHostWherePort(target, port); host.Address != "" {
			// Can only have on port with this value
			index := 0
			// Update the object from the database
			//Convert to nikto to json, it will contains only the necessary fields
			commandOutput, _ := json.Marshal(nuclei)
			portComment := database.PortComment{Task: taskName, CommandOutput: string(commandOutput)}
			host.Ports[index].PortComment = append(host.Ports[index].PortComment, portComment)
			db.AddOrUpdateHost(host)

		} else {
			log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
		}
	} else {
		log.INFO.Println("Nothing found by nuclei")
	}
	return true
}

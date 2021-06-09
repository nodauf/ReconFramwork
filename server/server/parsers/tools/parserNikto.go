package parsersTools

import (
	"encoding/json"
	"strconv"

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
		port, _ := strconv.Atoi(nikto.Port)

		if targetObject := db.GetTarget(target); targetObject != nil {
			commandOutput, _ := json.Marshal(nikto)
			portComment := modelsDatabases.PortComment{Task: taskName, CommandOutput: string(commandOutput)}
			err := db.AddPortComment(targetObject, port, portComment)
			if err != nil {
				log.ERROR.Println(err)
			}

		} else {
			log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
		}
	} else {
		log.INFO.Println("Nothing found by nikto")
	}
	return true
}

func (parse Parser) PrintOutputNikto(data string) (string, bool) {
	html := false
	var nikto modelsParsers.Nikto
	var output string

	json.Unmarshal([]byte(data), &nikto)
	for _, finding := range nikto.Vulnerabilities {
		output += "+ " + finding.OSVDB + ": " + finding.URL + ": " + finding.Msg + "\n"
	}
	return output, html
}

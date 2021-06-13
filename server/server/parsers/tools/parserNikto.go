package parsersTools

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/db"
	"github.com/nodauf/ReconFramwork/server/server/models"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/server/models/parsers"
)

func (parse Parser) ParseNikto(outputBytes []byte) bool {
	var nikto modelsParsers.Nikto
	var output models.Output
	json.Unmarshal(outputBytes, &output)

	err := json.Unmarshal([]byte(output.Stdout), &nikto)

	if err == nil && len(nikto.Vulnerabilities) > 0 {
		target := nikto.IP
		port, _ := strconv.Atoi(nikto.Port)

		if targetObject := db.GetTarget(target); targetObject != nil {
			var comment string
			commandOutput, _ := json.Marshal(nikto)
			if strings.Contains(string(commandOutput), "Directory listing found") {
				comment = "Directory indexing found"
			}
			portComment := modelsDatabases.PortComment{Task: output.TaskName, CommandOutput: string(commandOutput), Comment: comment}
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

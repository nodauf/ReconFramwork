package parsersTools

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/models/parsers"
)

func (parse Parser) ParseFfuf(taskName, cmdline, stdout, stderr string) bool {
	var ffuf modelsParsers.Ffuf
	err := json.Unmarshal([]byte(stdout), &ffuf)

	// If there is something found by ffuf
	if err == nil && len(ffuf.Results) > 0 {
		commandline := strings.Split(ffuf.Commandline, "/")[2]
		target := strings.Split(commandline, ":")[0]
		port, _ := strconv.Atoi(strings.Split(commandline, ":")[1])
		if targetObject := db.GetTarget(target); targetObject != nil {
			// Update the object from the database
			//Convert to ffuf to json, it will contains only the necessary fields
			commandOutput, _ := json.Marshal(ffuf)
			portComment := modelsDatabases.PortComment{Task: taskName, CommandOutput: string(commandOutput)}

			err := db.AddPortComment(targetObject, port, portComment)
			if err != nil {
				log.ERROR.Println(err)
			}
			//db.AddOrUpdateHost(&host)

		} else {
			log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
		}
	}
	return true
}

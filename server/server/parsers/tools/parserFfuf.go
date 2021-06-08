package parsersTools

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/server/models/parsers"
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

func (parse Parser) PrintOutputFfuf(data string) string {
	var output string
	var ffuf modelsParsers.Ffuf
	json.Unmarshal([]byte(data), &ffuf)
	for _, result := range ffuf.Results {
		output += result.URL + "\t" + "[Status: " + strconv.Itoa(result.Status) + ", Size: " + strconv.Itoa(result.Length) + ", Lines: " + strconv.Itoa(result.Lines) + "]\n"
	}
	return output
}

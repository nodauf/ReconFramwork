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

func (parse Parser) ParseFfuf(outputBytes []byte) bool {
	var ffuf modelsParsers.Ffuf
	var output models.Output
	json.Unmarshal(outputBytes, &output)
	err := json.Unmarshal([]byte(output.Stdout), &ffuf)

	// If there is something found by ffuf
	if err == nil && len(ffuf.Results) > 0 {
		commandline := strings.Split(ffuf.Commandline, "/")[2]
		target := strings.Split(commandline, ":")[0]
		port, _ := strconv.Atoi(strings.Split(commandline, ":")[1])
		if targetObject := db.GetTarget(target); targetObject != nil {
			var comment string
			comment = strconv.Itoa(len(ffuf.Results)) + " element found"
			// Update the object from the database
			//Convert to ffuf to json, it will contains only the necessary fields
			commandOutput, _ := json.Marshal(ffuf)
			portComment := modelsDatabases.PortComment{Task: output.TaskName, CommandOutput: string(commandOutput), Comment: comment}

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

func (parse Parser) PrintOutputFfuf(data string) (string, bool) {
	html := false
	var output string
	var ffuf modelsParsers.Ffuf
	json.Unmarshal([]byte(data), &ffuf)
	for _, result := range ffuf.Results {
		output += result.URL + "\t" + "[Status: " + strconv.Itoa(result.Status) + ", Size: " + strconv.Itoa(result.Length) + ", Lines: " + strconv.Itoa(result.Lines) + "]\n"
	}
	return output, html
}

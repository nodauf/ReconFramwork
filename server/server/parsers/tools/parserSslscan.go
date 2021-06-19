package parsersTools

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/db"
	"github.com/nodauf/ReconFramwork/server/server/models"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
)

func (parse Parser) ParseSslscan(outputBytes []byte) bool {
	var output models.Output
	json.Unmarshal(outputBytes, &output)
	if output.Error == "" && output.Stderr == "" {
		targetWithPort := strings.Split(output.Cmd, " ")[1]
		target := strings.Split(targetWithPort, ":")[0]
		port, _ := strconv.Atoi(strings.Split(targetWithPort, ":")[1])
		if targetObject := db.GetTarget(target); targetObject != nil {
			portComment := modelsDatabases.PortComment{Task: output.TaskName, CommandOutput: output.Stdout}
			err := db.AddPortComment(targetObject, port, portComment)
			if err != nil {
				log.ERROR.Println(err)
			}
		}
	}

	return true
}

func (parse Parser) PrintOutputSslscan(data string) (string, bool) {
	html := true
	output := data
	output = strings.ReplaceAll(output, "[31m", "<span style='color:red'>")
	output = strings.ReplaceAll(output, "[32m", "<span style='color:mediumseagreen'>")
	output = strings.ReplaceAll(output, "[33m", "<span style='color:yellow'>")
	output = strings.ReplaceAll(output, "[1;34m", "<span style='color:violet'>")
	output = strings.ReplaceAll(output, "[0m", "</span>")
	return output, html
}

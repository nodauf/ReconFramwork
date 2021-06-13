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

func (parse Parser) ParseGowitness(outputBytes []byte) bool {
	var output models.Output
	json.Unmarshal(outputBytes, &output)
	if len(output.Stdout) > 0 {
		hostWithPort := strings.Split(strings.Split(output.Cmd, "/")[2], " ")[0]
		target := strings.Split(hostWithPort, ":")[0]
		port, _ := strconv.Atoi(strings.Split(hostWithPort, ":")[1])
		if targetObject := db.GetTarget(target); targetObject != nil {
			portComment := modelsDatabases.PortComment{Task: output.TaskName, CommandOutput: output.Stdout}
			err := db.AddPortComment(targetObject, port, portComment)
			if err != nil {
				log.ERROR.Println(err)
			}

		} else {
			log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
		}
	}
	return true
}

func (parse Parser) PrintOutputGowitness(data string) (string, bool) {
	html := true
	output := `<img src="data:image/png;base64,` + data + `" alt="Image gowitness" class="img-fluid" />`
	return string(output), html
}

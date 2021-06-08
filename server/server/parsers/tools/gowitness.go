package parsersTools

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
)

func (parse Parser) ParseGowitness(taskName, cmdline, stdout, stderr string) bool {
	if len(stdout) > 0 {
		hostWithPort := strings.Split(strings.Split(cmdline, "/")[2], " ")[0]
		target := strings.Split(hostWithPort, ":")[0]
		port, _ := strconv.Atoi(strings.Split(hostWithPort, ":")[1])
		if targetObject := db.GetTarget(target); targetObject != nil {
			portComment := modelsDatabases.PortComment{Task: taskName, CommandOutput: stdout}
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
	fmt.Println(output)
	return string(output), html
}

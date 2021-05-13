package parsers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/models/parsers"
	"github.com/nodauf/ReconFramwork/utils"
)

func (parse Parser) ParseFfuf(stdout, stderr string) bool {
	/*fmt.Println("parsing ffuf ...")
	fmt.Println("stdout: " + stdout)
	fmt.Println("stderr: " + stderr)*/
	var ffuf modelsParsers.Ffuf
	var portList []string
	json.Unmarshal([]byte(stdout), &ffuf)

	// If there is something found by ffuf
	if len(ffuf.Results) > 0 {
		commandline := strings.Split(ffuf.Commandline, "/")[2]
		fmt.Println(commandline)
		target := strings.Split(commandline, ":")[0]
		port := strings.Split(commandline, ":")[1]
		if host := db.GetHost(target); host.Address != "" {
			for _, portHost := range host.Ports {
				portList = append(portList, strconv.Itoa(portHost.Port))
			}

			if index, ok := utils.StringInSlice(port, portList); ok {
				// Update the object from the database
				commandOutput, _ := json.Marshal(ffuf)
				portComment := database.PortComment{Tool: "ffuf", CommandOutput: string(commandOutput)}
				host.Ports[index].PortComment = append(host.Ports[index].PortComment, portComment)
				db.AddOrUpdateHost(host)
			} else {
				log.ERROR.Println("Something went wrong. Port " + port + " not found for the Host " + target + " in the database.")
			}

		} else {
			log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
		}
	}
	return true
}

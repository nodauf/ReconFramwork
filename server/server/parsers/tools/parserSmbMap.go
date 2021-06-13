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

func (parse Parser) ParseSmbmap(outputBytes []byte) bool {
	//regexIPPort := `\[\+\] IP: (?P<IP>[A-Za-z1-9\.]+):(?P<port>\d+) +Name:`
	//regexIPPort := `(?s)\[\+\] IP: (?P<IP>[A-Za-z0-9\.]+):(?P<port>[0-9]+).+Name:`

	//r := regexp.MustCompile(regexIPPort)
	//resultRegex := r.FindStringSubmatch(stdout)

	var output models.Output
	json.Unmarshal(outputBytes, &output)
	target := strings.Split(output.Cmd, " ")[2]
	port, _ := strconv.Atoi(strings.Split(output.Cmd, " ")[4])
	if targetObject := db.GetTarget(target); targetObject != nil {
		comment := ""
		if strings.Contains(output.Stdout, "READ") {
			comment = "Anonymous READ allowed\n"
		}
		if strings.Contains(output.Stdout, "WRITE") {
			comment += "Anonymous WRITE allowed\n"
		}

		portComment := modelsDatabases.PortComment{Task: output.TaskName, CommandOutput: output.Stdout, Comment: comment}
		err := db.AddPortComment(targetObject, port, portComment)
		if err != nil {
			log.ERROR.Println(err)
		}

	} else {
		log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
	}
	return true
}

func (parse Parser) PrintOutputSmbMap(data string) (string, bool) {
	html := true
	output := data
	output = strings.ReplaceAll(output, "[32mREAD, WRITE[0m", "<span style='color:mediumseagreen'>READ, WRITE</span>")
	output = strings.ReplaceAll(output, "[31mNO ACCESS[0m", "<span style='color:red'>NO ACCESS</span>")
	return output, html
}

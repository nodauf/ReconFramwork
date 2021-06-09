package parsersTools

import (
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
)

func (parse Parser) ParseSmbmap(taskName, cmdline, stdout, stderr string) bool {
	//regexIPPort := `\[\+\] IP: (?P<IP>[A-Za-z1-9\.]+):(?P<port>\d+) +Name:`
	//regexIPPort := `(?s)\[\+\] IP: (?P<IP>[A-Za-z0-9\.]+):(?P<port>[0-9]+).+Name:`

	//r := regexp.MustCompile(regexIPPort)
	//resultRegex := r.FindStringSubmatch(stdout)
	target := strings.Split(cmdline, " ")[2]
	port, _ := strconv.Atoi(strings.Split(cmdline, " ")[4])
	if targetObject := db.GetTarget(target); targetObject != nil {
		comment := ""
		if strings.Contains(stdout, "READ") {
			comment = "Anonymous READ allowed\n"
		}
		if strings.Contains(stdout, "WRITE") {
			comment += "Anonymous WRITE allowed\n"
		}

		portComment := modelsDatabases.PortComment{Task: taskName, CommandOutput: stdout, Comment: comment}
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

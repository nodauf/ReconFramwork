package parsersTools

import (
	"regexp"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
)

func (parse Parser) ParseSmbmap(taskName, cmdline, stdout, stderr string) bool {
	//regexIPPort := `\[\+\] IP: (?P<IP>[A-Za-z1-9\.]+):(?P<port>\d+) +Name:`
	regexIPPort := `(?s)\[\+\] IP: (?P<IP>[A-Za-z0-9\.]+):(?P<port>[0-9]+).+Name:`

	r := regexp.MustCompile(regexIPPort)
	resultRegex := r.FindStringSubmatch(stdout)
	target := resultRegex[1]
	port := resultRegex[2]
	if host := db.GetHostWherePort(target, port); host.Address != "" {
		// Can only have on port with this value
		index := 0
		comment := ""
		if strings.Contains(stdout, "READ") {
			comment = "Anonymous READ allowed\n"
		}
		if strings.Contains(stdout, "WRITE") {
			comment += "Anonymous WRITE allowed\n"
		}
		portComment := modelsDatabases.PortComment{Task: taskName, CommandOutput: string(stdout), Comment: comment}
		host.Ports[index].PortComment = append(host.Ports[index].PortComment, portComment)
		db.AddOrUpdateHost(&host)

	} else {
		log.ERROR.Println("Something went wrong. Host " + target + " not found in the database or the port was not open for this host")
	}
	return true
}

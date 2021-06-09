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

func (parse Parser) ParseNuclei(taskName, cmdline, stdout, stderr string) bool {
	var nuclei modelsParsers.Nuclei
	var nucleiFinding modelsParsers.NucleiFinding
	for _, finding := range strings.Split(stdout, "\n") {

		err := json.Unmarshal([]byte(finding), &nucleiFinding)
		if err == nil {
			nuclei.Findings = append(nuclei.Findings, nucleiFinding)
		}
	}
	if len(nuclei.Findings) > 0 {
		target := nuclei.Findings[0].IP
		port, _ := strconv.Atoi(strings.Split(nuclei.Findings[0].Host, ":")[2])
		if targetObject := db.GetTarget(target); targetObject != nil {
			commandOutput, _ := json.Marshal(nuclei)
			portComment := modelsDatabases.PortComment{Task: taskName, CommandOutput: string(commandOutput)}
			err := db.AddPortComment(targetObject, port, portComment)
			if err != nil {
				log.ERROR.Println(err)
			}

		} else {
			log.ERROR.Println("Something went wrong. Host " + target + " not found in the database")
		}
	} else {
		log.INFO.Println("Nothing found by nuclei")
	}
	return true
}

func (parse Parser) PrintOutputNuclei(data string) (string, bool) {
	html := true
	var nuclei modelsParsers.Nuclei
	var output string
	json.Unmarshal([]byte(data), &nuclei)
	for _, finding := range nuclei.Findings {
		var severity string
		var timestamp string
		var templateID string
		var typeString string

		if finding.Info.Severity == "info" {
			severity = "<span style='color:cyan'>" + finding.Info.Severity + "</span>"
		} else if finding.Info.Severity == "low" {
			severity = "<span style='color:mediumseagreen'>" + finding.Info.Severity + "</span>"
		} else if finding.Info.Severity == "medium" {
			severity = "<span style='color:yellow'>" + finding.Info.Severity + "</span>"
		} else if finding.Info.Severity == "high" {
			severity = "<span style='color:orange'>" + finding.Info.Severity + "</span>"
		} else if finding.Info.Severity == "critical" {
			severity = "<span style='color:red'>" + finding.Info.Severity + "</span>"
		}
		timestamp = "<span style='color:darkCyan'>" + finding.Timestamp + "</span>"
		templateID = "<span style='color:mediumseagreen'>" + finding.TemplateID + "</span>"
		typeString = "<span style='color:dodgerblue'>" + finding.Type + "</span>"

		output += "[" + timestamp + "]\t[" + templateID + "]\t[" + typeString + "]\t[" + severity + "]\t" + finding.Matched + "\n"
	}
	return output, html
}

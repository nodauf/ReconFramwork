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

func (parse Parser) ParseNuclei(outputBytes []byte) bool {
	var output models.Output
	json.Unmarshal(outputBytes, &output)
	var nuclei modelsParsers.Nuclei
	var nucleiFinding modelsParsers.NucleiFinding
	for _, finding := range strings.Split(output.Stdout, "\n") {

		err := json.Unmarshal([]byte(finding), &nucleiFinding)
		if err == nil {
			nuclei.Findings = append(nuclei.Findings, nucleiFinding)
		}
	}
	if len(nuclei.Findings) > 0 {
		target := nuclei.Findings[0].IP
		port, _ := strconv.Atoi(strings.Split(nuclei.Findings[0].Host, ":")[2])
		if targetObject := db.GetTarget(target); targetObject != nil {
			comment := stats(nuclei)
			comment += "<br />\n" + getInterestingTech(nuclei)
			commandOutput, _ := json.Marshal(nuclei)
			portComment := modelsDatabases.PortComment{Task: output.TaskName, CommandOutput: string(commandOutput), Comment: comment}
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

func stats(findings modelsParsers.Nuclei) string {
	var infoNb, lowNb, mediumNb, highNb, criticalNb int
	var comment string
	for _, finding := range findings.Findings {
		if finding.Info.Severity == "info" {
			infoNb++
		} else if finding.Info.Severity == "low" {
			lowNb++
		} else if finding.Info.Severity == "medium" {
			mediumNb++
		} else if finding.Info.Severity == "high" {
			highNb++
		} else if finding.Info.Severity == "critical" {
			criticalNb++
		}
	}
	comment = "Findings severity: <span style='color:cyan'>" + strconv.Itoa(infoNb) + " info</span>, <span style='color:mediumseagreen'>" + strconv.Itoa(lowNb) + " low</span>, <span style='color:yellow'>" + strconv.Itoa(mediumNb) + " medium</span>, <span style='color:orange'>" + strconv.Itoa(highNb) + " high</span>, <span style='color:red'>" + strconv.Itoa(criticalNb) + " critical</span>."
	return comment
}

func getInterestingTech(findings modelsParsers.Nuclei) string {
	var comment string

	for _, finding := range findings.Findings {
		// Two awesome templates to detect A LOT of technologies
		if finding.TemplateID == "tech-detect" {
			comment += finding.MatcherName
		} else if finding.TemplateID == "favicon-detection" {
			comment += finding.MatcherName + "<br />\n"
		}
	}

	return comment
}

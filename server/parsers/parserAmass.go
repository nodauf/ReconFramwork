package parsers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/models/parsers"
)

func (parse Parser) ParseAmass(taskName, cmdline, stdout, stderr string) bool {
	var amass modelsParsers.Amass
	var amassDomain modelsParsers.AmassDomain
	for _, domain := range strings.Split(stdout, "\n") {

		err := json.Unmarshal([]byte(domain), &amassDomain)
		if err == nil {
			amass.Domains = append(amass.Domains, amassDomain)
		}
	}
	fmt.Println(parse.Job)
	if len(amass.Domains) > 0 {
		var host database.Host
		host = parse.Job.Host
		if host.ID != 0 {
			for _, domain := range amass.Domains {
				var domainDB database.Domain

				domainDB.Domain = domain.Name
				for _, address := range domain.Addresses {
					var host database.Host
					host.Address = address.IP
					domainDB.Host = append(domainDB.Host, host)
				}
				db.AddDomain(domainDB)
			}

		} else {
			log.ERROR.Println("Something went wrong. Host " + host.Address + " not found in the database")
		}
	} else {
		log.INFO.Println("Nothing found by nuclei")
	}
	return true
}

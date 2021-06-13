package parsersTools

import (
	"encoding/json"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/db"
	"github.com/nodauf/ReconFramwork/server/server/models"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/server/models/parsers"
)

func (parse Parser) ParseAmass(outputBytes []byte) bool {
	var amass modelsParsers.Amass
	var amassDomain modelsParsers.AmassDomain
	var output models.Output
	json.Unmarshal(outputBytes, &output)
	for _, domain := range strings.Split(output.Stdout, "\n") {

		err := json.Unmarshal([]byte(domain), &amassDomain)
		if err == nil {
			amass.Domains = append(amass.Domains, amassDomain)
		}
	}

	if len(amass.Domains) > 0 {
		var host modelsDatabases.Host
		host = parse.Job.Host
		if host.ID != 0 {
			// If the job is attached to a host
			for _, domain := range amass.Domains {
				var domainDB modelsDatabases.Domain

				domainDB.Domain = domain.Name
				for _, address := range domain.Addresses {
					var host modelsDatabases.Host
					host.Address = address.IP
					domainDB.Host = append(domainDB.Host, host)
				}
				db.AddDomain(domainDB)
			}
		} else if parse.Job.Domain.ID != 0 {
			// If the job is attached to a domain
			for _, domain := range amass.Domains {
				var domainDB modelsDatabases.Domain

				domainDB.Domain = domain.Name
				for _, address := range domain.Addresses {
					var host modelsDatabases.Host
					host.Address = address.IP
					domainDB.Host = append(domainDB.Host, host)
				}
				//domainDB.SubdomainOf = &parse.Job.Domain

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

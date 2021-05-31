package parsersCustoms

import (
	"fmt"
	"strings"

	"github.com/nodauf/ReconFramwork/server/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
)

func (parse Parser) ParseDomainFromCert(taskName, cmdline, stdout, stderr string) bool {
	hostnameWithIP := make(map[string]string)
	// Convert the string to a map to be sure there is no duplicate host
	for _, line := range strings.Split(stdout, "\n") {
		hostname := strings.Split(line, ":")[0]
		ip := strings.Split(line, ":")[1]
		hostnameWithIP[hostname] = ip
	}
	for hostname, ip := range hostnameWithIP {
		var domain modelsDatabases.Domain
		var host modelsDatabases.Host

		host.Address = ip
		domain.Domain = hostname
		domain.Host = append(domain.Host, host)

		db.AddDomain(domain)
	}
	fmt.Println("Parse DomainFromCert")
	return true
}

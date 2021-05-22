package parsers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/nodauf/ReconFramwork/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/models/parsers"
	"github.com/nodauf/ReconFramwork/utils"
)

func (parse Parser) ParseNmap(taskName, cmdline, stdout, stderr string) bool {
	fmt.Println("parse nmap")
	fmt.Println(stdout)
	var nmap modelsParsers.Nmaprun
	var portList []string
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal([]byte(stdout), &nmap)
	//fmt.Printf("%#v \n", nmap)
	//empJSON, _ := json.MarshalIndent(nmap, "", "  ")
	//fmt.Println(string(empJSON))

	// Get the object host from the database if it exists
	var host modelsDatabases.Host
	if host = db.GetHost(nmap.Host.Address.Addr); host.Address != "" {
		for _, port := range host.Ports {
			portList = append(portList, strconv.Itoa(port.Port))
		}

	}
	host.Address = nmap.Host.Address.Addr

	var ports []modelsDatabases.Port
	for _, portNmap := range nmap.Host.Ports.Port {
		if portNmap.State.State == "open" {
			var port modelsDatabases.Port
			// If we retrieve the object from the database it may already have some ports
			if index, ok := utils.StringInSlice(portNmap.Portid, portList); ok {
				port = host.Ports[index]
			} else {

				port.Port, _ = strconv.Atoi(portNmap.Portid)
				port.Service = portNmap.Service.Name
				port.Version = portNmap.Service.Version
			}

			var portComment modelsDatabases.PortComment
			portComment.CommandOutput = portNmap.Script.Output
			portComment.Task = taskName
			port.PortComment = append(port.PortComment, portComment)

			ports = append(ports, port)
		}
	}
	host.Ports = append(host.Ports, ports...)
	t, _ := json.MarshalIndent(host, "", " ")
	fmt.Println(string(t))

	// Workaround attach an existing domain if the host does not exist does not work
	db.AddOrUpdateHost(&host)
	if len(nmap.Host.Hostnames.Hostname) > 0 {
		var domain modelsDatabases.Domain
		domain.Domain = nmap.Host.Hostnames.Hostname[0].Name
		host.Domain = append(host.Domain, domain)
		db.AddOrUpdateHost(&host)
	}

	//fmt.Println("stdout: " + stdout)
	//fmt.Println("stderr: " + stderr)
	if stderr != "" {
		return false
	}
	return true
}

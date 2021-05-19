package parsers

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/models/parsers"
	"github.com/nodauf/ReconFramwork/utils"
)

func (parse Parser) ParseNmap(taskName, cmdline, stdout, stderr string) bool {
	fmt.Println("parse nmap")
	var nmap modelsParsers.Nmaprun
	var portList []string
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal([]byte(stdout), &nmap)
	//fmt.Printf("%#v \n", nmap)
	//empJSON, _ := json.MarshalIndent(nmap, "", "  ")
	//fmt.Println(string(empJSON))

	// Get the object host from the database if it exists
	var host database.Host
	if host = db.GetHost(nmap.Host.Address.Addr); host.Address != "" {
		for _, port := range host.Ports {
			portList = append(portList, strconv.Itoa(port.Port))
		}

	} else {
		host.Address = nmap.Host.Address.Addr
		if len(nmap.Host.Hostnames.Hostname) > 0 {
			var domain database.Domain
			domain.Domain = nmap.Host.Hostnames.Hostname[0].Name
			host.Domain = append(host.Domain, domain)
		}
	}

	var ports []database.Port
	for _, portNmap := range nmap.Host.Ports.Port {

		var port database.Port
		// If we retrieve the object from the database it may already have some ports
		if index, ok := utils.StringInSlice(portNmap.Portid, portList); ok {
			port = host.Ports[index]
		} else {

			port.Port, _ = strconv.Atoi(portNmap.Portid)
			port.Service = portNmap.Service.Name
			port.Version = portNmap.Service.Version
		}

		var portComment database.PortComment
		portComment.CommandOutput = portNmap.Script.Output
		portComment.Task = taskName
		port.PortComment = append(port.PortComment, portComment)

		ports = append(ports, port)
	}
	host.Ports = append(host.Ports, ports...)
	db.AddOrUpdateHost(&host)

	//fmt.Println("stdout: " + stdout)
	//fmt.Println("stderr: " + stderr)
	if stderr != "" {
		return false
	}
	return true
}

package parsers

import (
	"encoding/xml"
	"strconv"

	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/models/database"
	modelsParsers "github.com/nodauf/ReconFramwork/server/models/parsers"
	"github.com/nodauf/ReconFramwork/utils"
)

func (parse Parser) ParseNmap(stdout, stderr string) bool {
	tool := "nmap"
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
		host.Hostname = nmap.Host.Hostnames.Hostname[0].Name
	}

	var ports []database.Port
	for _, portNmap := range nmap.Host.Ports.Port {
		if utils.StringInSlice(portNmap.Portid, portList) {
			break
		}
		var port database.Port
		var portComment database.PortComment
		port.Port, _ = strconv.Atoi(portNmap.Portid)
		port.Service = portNmap.Service.Name
		port.Version = portNmap.Service.Version
		portComment.Comment = portNmap.Script.Output
		portComment.Tool = tool
		port.PortComment = append(port.PortComment, portComment)

		ports = append(ports, port)
	}
	host.Ports = ports

	db.AddOrUpdateHost(host)

	//fmt.Println("stdout: " + stdout)
	//fmt.Println("stderr: " + stderr)
	if stderr != "" {
		return false
	}
	return true
}

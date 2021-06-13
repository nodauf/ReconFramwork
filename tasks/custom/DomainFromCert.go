package customTasks

import (
	"crypto/tls"
	"encoding/json"
	"net"
	"regexp"

	"github.com/nodauf/ReconFramwork/server/server/models"
)

func DomainFromCert(taskName, cmd string) ([]byte, error) {
	var certificates string
	var output models.Output
	output.TaskName = taskName
	output.Cmd = cmd

	conn, err := tls.Dial("tcp", cmd, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		output.Error = err
		outputBytes, _ := json.Marshal(output)
		return outputBytes, nil
	}
	state := conn.ConnectionState()

	conn.Close()

	// Only analyze the website's certificate
	websiteCerts := state.PeerCertificates[0]
	regex, err := regexp.Compile("^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?$")
	if regex.MatchString(websiteCerts.Subject.CommonName) {
		certificates += websiteCerts.Subject.CommonName + ":" + resolveHostname(websiteCerts.Subject.CommonName)
	}

	for _, websiteCert := range websiteCerts.DNSNames {
		if regex.MatchString(websiteCert) {
			certificates += "\n" + websiteCert + ":" + resolveHostname(websiteCert)
		}

	}
	output.Stdout = certificates
	outputBytes, _ := json.Marshal(output)
	return outputBytes, nil
}

func resolveHostname(hostname string) string {
	var ip string
	addrs, err := net.LookupIP(hostname)
	if err == nil {
		for _, addr := range addrs {
			ip += addr.String() + ","
		}
	}
	// Return ip except the last character which is a comma
	return ip[:len(ip)-1]
}

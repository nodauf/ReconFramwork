package customTasks

import (
	"crypto/tls"
	"net"
	"regexp"
)

func DomainFromCert(taskName, cmd string) (string, string, string, string, error) {
	var certificates string

	conn, err := tls.Dial("tcp", cmd, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return taskName, cmd, "", "", err
	}
	state := conn.ConnectionState()
	if err != nil {
		return taskName, cmd, "", "", err
	}

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
	return taskName, cmd, certificates, "", nil
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

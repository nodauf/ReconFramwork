package parsersCustoms

import "fmt"

func (parse Parser) ParseDomainFromCert(taskName, cmdline, stdout, stderr string) bool {
	fmt.Println(stdout)
	fmt.Println("Parse DomainFromCert")
	return true
}

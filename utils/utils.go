package utils

import (
	"net"
	"strings"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/tasks"
)

func GetMachineryServer() (*machinery.Server, error) {
	log.INFO.Println("initing task server")

	server, err := machinery.NewServer(&config.Config{
		Broker:        "redis://localhost:6379",
		ResultBackend: "redis://localhost:6379",
	})
	if err == nil {

		err = server.RegisterTasks(map[string]interface{}{
			"a":      tasks.HelloWorld,
			"runcmd": tasks.RunCmd,
		})
	}
	return server, err

}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsNetwork(value string) bool {
	_, _, errParseCIDR := net.ParseCIDR(value)
	ip := net.ParseIP(value)
	// If this is a valid CIDR and not a valid IP, it's a network
	if errParseCIDR == nil && ip == nil {
		return true
	}
	return false
}

func HostsFromNetwork(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// remove network address and broadcast address
	lenIPs := len(ips)
	switch {
	case lenIPs < 2:
		return ips, nil

	default:
		return ips[1 : len(ips)-1], nil
	}
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func ParseList(list string) []string {
	var finalItems []string
	items := strings.Split(list, ",")
	for _, item := range items {
		item = strings.Trim(item, " ")
		finalItems = append(finalItems, item)
	}
	return finalItems
}

package utils

import (
	"math/rand"
	"net"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func StringInSlice(a string, list []string) (int, bool) {
	for i, b := range list {
		if b == a {
			return i, true
		}
	}
	return -1, false
}

func IntInSlice(a int, list []int) (int, bool) {
	for i, b := range list {
		if b == a {
			return i, true
		}
	}
	return -1, false
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

func IsIP(value string) bool {
	ip := net.ParseIP(value)

	if ip == nil {
		return false
	}
	return true
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

// Credits https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandomString(n int) string {
	letterIdxBits := 6
	letterIdxMask := int64(1<<letterIdxBits - 1)
	letterIdxMax := 63 / letterIdxBits
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var src = rand.NewSource(time.Now().UnixNano())

	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()

}

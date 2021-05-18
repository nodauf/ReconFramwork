package modelsParsers

type Amass struct {
	Domains []AmassDomain
}

type AmassDomain struct {
	Addresses []struct {
		Asn  int64  `json:"asn"`
		Cidr string `json:"cidr"`
		Desc string `json:"desc"`
		IP   string `json:"ip"`
	} `json:"addresses"`
	Domain  string   `json:"domain"`
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
	Tag     string   `json:"tag"`
}

package modelsParsers

type Nikto struct {
	Host            string `json:"host"`
	IP              string `json:"ip"`
	Port            string `json:"port"`
	Banner          string `json:"banner"`
	Vulnerabilities []struct {
		ID     string `json:"id"`
		OSVDB  string `json:"OSVDB"`
		Method string `json:"method"`
		URL    string `json:"url"`
		Msg    string `json:"msg"`
	} `json:"vulnerabilities"`
	Msg string `json:"msg"`
}

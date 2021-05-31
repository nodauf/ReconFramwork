package modelsParsers

type Nuclei struct {
	Findings []NucleiFinding
}

type NucleiFinding struct {
	Host string `json:"host"`
	Info struct {
		Author   string `json:"author"`
		Name     string `json:"name"`
		Severity string `json:"severity"`
		Tags     string `json:"tags"`
	} `json:"info"`
	IP         string `json:"ip"`
	Matched    string `json:"matched"`
	TemplateID string `json:"templateID"`
	Timestamp  string `json:"timestamp"`
	Type       string `json:"type"`
}

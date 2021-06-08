package modelsBeego

type Result struct {
	Address string
	Domain  []string
	NbPorts int
}

type ResultWeb struct {
	Address string
	Domain  []string
	WebPort int
	Task    string
}

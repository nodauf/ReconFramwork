package modelsBeego

type ResultTask struct {
	Address string
	Domain  []string
	NbPorts int
	Port    int
	Task    string
	Comment string
}

type ResultJob struct {
	ID            uint
	Target        string
	ExecutionTime string
	MachineryTask string
	MachineryArgs string
	Processed     bool
}

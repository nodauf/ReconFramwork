package models

type Workflow struct {
	Name     string         `yaml:"name"`
	Tags     []string       `yaml:"tags,omitempty"`
	Commands []string       `yaml:"commands,omitempty"`
	Options  WorkflowOption `yaml:"options,omitempty"`
}

type WorkflowOption struct {
	ParallelizeTasks bool `yaml:"parallelizeTasks"`
}

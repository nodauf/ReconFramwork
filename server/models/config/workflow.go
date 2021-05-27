package modelsConfig

type Workflow struct {
	Name        string         `yaml:"name"`
	Description string         `yaml:"description"`
	Tags        []string       `yaml:"tags,omitempty"`
	Commands    []string       `yaml:"commands,omitempty"`
	Options     WorkflowOption `yaml:"options,omitempty"`
}

type WorkflowOption struct {
	ParallelizeTasks bool `yaml:"parallelizeTasks"`
}

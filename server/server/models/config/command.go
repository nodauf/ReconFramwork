package modelsConfig

type Command struct {
	Name           string                    `yaml:"name"`
	Description    string                    `yaml:"description"`
	Tags           []string                  `yaml:"tags,omitempty"`
	Target         string                    `yaml:"target"`
	Service        map[string]CommandService `yaml:"service,omitempty"`
	Port           string                    `yaml:"port,omitempty"`
	Cmd            string                    `yaml:"cmd"`
	Variable       map[string]string         `yaml:"variable,omitempty"`
	Regex          []string                  `yaml:"regex,omitempty"`
	RegexSuccess   string                    `yaml:"regexSuccess,omitempty"`
	ParserFunction string                    `yaml:"parserFunction,omitempty"`
	PrintFunction  string                    `yaml:"printFunction,omitempty"`
	CustomTask     string                    `yaml:"customTask"`
	//Parser         Parser
}

type CommandService struct {
	Variable map[string]string `yaml:"variable,omitempty"`
}

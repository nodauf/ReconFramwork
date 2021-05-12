package models

type Command struct {
	Name           string            `yaml:"name"`
	Tags           []string          `yaml:"tags,omitempty"`
	Target         string            `yaml:"target"`
	Service        string            `yaml:"service,omitempty"`
	Port           string            `yaml:"port,omitempty"`
	Cmd            string            `yaml:"cmd"`
	Variable       map[string]string `yaml:"variable,omitempty"`
	Regex          []string          `yaml:"regex,omitempty"`
	RegexSuccess   string            `yaml:"regexSuccess,omitempty"`
	ParserFunction string            `yaml:"parserFunction,omitempty"`
	//Parser         Parser
}

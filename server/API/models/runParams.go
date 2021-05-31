package models

type RunParams struct {
	Target   string `json:"target" binding:"required"`
	Task     string `json:"task"`
	Workflow string `json:"workflow"`
	Options  struct {
		RecurseOnSubdomain bool `json:"RecurseOnSubdomain"`
	} `json:"options"`
}

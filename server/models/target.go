package models

import (
	modelsConfig "github.com/nodauf/ReconFramwork/server/models/config"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/models/database"
	"github.com/nodauf/ReconFramwork/utils"
)

type Target interface {
	HasService(map[string]modelsConfig.CommandService) map[string]string
	HasPort(int) int
	AddPortComment(int, modelsDatabases.PortComment) ([]modelsDatabases.Host, error)
	GetSubdomain() []string
	GetTarget() string
}

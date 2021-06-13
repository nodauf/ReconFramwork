package models

import (
	modelsConfig "github.com/nodauf/ReconFramwork/server/server/models/config"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	"github.com/nodauf/ReconFramwork/server/server/utils"
)

type Target interface {
	HasService(map[string]modelsConfig.CommandService) map[string]string
	HasPort(int) int
	AddPortComment(int, modelsDatabases.PortComment) ([]modelsDatabases.Host, error)
	GetDomain() []string
	GetTarget() string
}

func CreateTarget(target string) Target {
	var targetObject Target
	if utils.IsIP(target) {
		host := &modelsDatabases.Host{}
		host.Address = target
		targetObject = host

		// Otherwise this is a domain and we add it in domain table
	} else {
		domain := &modelsDatabases.Domain{}
		domain.Domain = target
		targetObject = domain
	}

	return targetObject

}

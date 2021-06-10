package db

import (
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
)

func GetPortComment(ip, port, task string) string {
	var host modelsDatabases.Host
	db.Preload("Ports", "port = ?", port).
		Preload("Ports.PortComment", "task = ?", task).
		Where("address = ?", ip).
		First(&host)
	return host.Ports[0].PortComment[0].CommandOutput
}

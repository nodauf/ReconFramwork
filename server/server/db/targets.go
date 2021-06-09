package db

import (
	"errors"
	"strconv"

	"github.com/nodauf/ReconFramwork/server/server/models"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	"github.com/nodauf/ReconFramwork/utils"
)

func GetTarget(target string) models.Target {
	var targetObject models.Target
	host := GetHost(target)
	domain := GetDomain(target)
	// If there is nothing in the datbase for this target
	if host.Address != "" || domain.Domain != "" {
		if host.Address != "" {
			targetObject = &host
		} else {
			targetObject = &domain
		}
	}
	return targetObject
}

func GetTargets() []string {
	var results []string
	targets := []struct {
		Target string
	}{}
	db.Distinct("address as target").Table("hosts").Scan(&targets)
	for _, target := range targets {
		results = append(results, target.Target)
	}
	db.Distinct("domain as target").Table("domains").Scan(&targets)
	for _, target := range targets {
		results = append(results, target.Target)
	}
	return results
}

func AddOrUpdateTarget(target models.Target) models.Target {
	var targetToReturn models.Target
	if utils.IsIP(target.GetTarget()) {
		host := AddOrUpdateHost(target.(*modelsDatabases.Host))
		targetToReturn = &host
	} else {
		domain := AddOrUpdateDomain(target.(*modelsDatabases.Domain))
		targetToReturn = &domain
	}
	return targetToReturn
}

func AddPortComment(targetObject models.Target, port int, portComment modelsDatabases.PortComment) error {

	if index := targetObject.HasPort(port); index != -1 {
		targetList, err := targetObject.AddPortComment(port, portComment)
		if err != nil {
			return err
		}
		for _, target := range targetList {
			AddOrUpdateHost(&target)
		}
	} else {
		return errors.New("The target " + targetObject.GetTarget() + " has not the port " + strconv.Itoa(port))
	}
	return nil
}

func DeleteTarget(target models.Target) bool {
	var returnValue bool
	if utils.IsIP(target.GetTarget()) {
		returnValue = DeleteHost(target.(*modelsDatabases.Host))
	} else {
		returnValue = DeleteDomain(target.(*modelsDatabases.Domain))
	}
	return returnValue
}

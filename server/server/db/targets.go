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

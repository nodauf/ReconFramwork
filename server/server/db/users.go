package db

import (
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	"github.com/nodauf/ReconFramwork/server/server/utils"
)

func AddUser(username, password string) (bool, error) {
	var user modelsDatabases.User
	user.Username = username
	user.Password = utils.HashPassword(password)
	result := db.Create(&user)
	if result.Error == nil {
		return true, nil
	}
	return false, result.Error
}

func DeleteUser(username string) (bool, error) {
	var user modelsDatabases.User
	result := db.Where("username = ?", username).First(&user)
	if result.RowsAffected == 1 || result.Error != nil {
		db.Unscoped().Delete(user)
		return true, nil
	} else {
		return false, result.Error
	}
}

func UserExist(username, password string) (uint, bool) {
	var user modelsDatabases.User
	result := db.Where("username = ?", username).First(&user)
	if result.RowsAffected == 1 || result.Error != nil {
		if utils.ComparePassword(user.Password, password) {
			return user.ID, true
		}
	}
	return 0, false

}

package db

import (
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/ReconFramwork?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//conn.Logger = conn.Logger.LogMode(logger.Info)
	conn.Logger = conn.Logger.LogMode(logger.Silent)
	// Auto Migrate
	conn.AutoMigrate(&modelsDatabases.Host{}, &modelsDatabases.Port{}, &modelsDatabases.PortComment{}, &modelsDatabases.Job{}, &modelsDatabases.Domain{}, &modelsDatabases.User{}) //, &database.HostsPorts{})
	// The foreign key is not needed by gorm and we don't need it as NULL != NULL, it will break unique index
	conn.Migrator().DropConstraint(&modelsDatabases.PortComment{}, "fk_domains_port_comment")
	// Set table options
	//db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&database.Host{})
	/*conn.Debug().Migrator().CreateConstraint(&database.Host{}, "Ports")
	conn.Debug().Migrator().CreateConstraint(&database.Port{}, "Hosts")
	conn.Debug().Migrator().CreateConstraint(&database.Port{}, "PortComment")
	conn.Debug().Migrator().CreateConstraint(&database.PortComment{}, "Host")
	conn.Debug().Migrator().CreateConstraint(&database.PortComment{}, "Port")*/
	db = conn
}

package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"tkktrd/golang-gin-realworld-example-app/config"
)

func NewMysqlConnect(config *config.Config) (*gorm.DB, error){
	dbConnectionInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySQL.MySQLUser,
		config.MySQL.MySQLPassword,
		config.MySQL.MySQLHost,
		config.MySQL.MySQLPort,
		config.MySQL.MySQLDatabaseName)

	dialector := mysql.Open(dbConnectionInfo)

	var err error
	for index := 0; index < config.MySQL.MySQLConnectRetries; index++ {
		time.Sleep(time.Second * time.Duration(index * 2))

		if db, err := gorm.Open(dialector, &gorm.Config{}); err == nil {
			return db, nil
		}
	}
	return nil, fmt.Errorf("failed to connect to database, %v", err)
}


package main

import (
	"log"
	"tkktrd/golang-gin-realworld-example-app/config"
	"tkktrd/golang-gin-realworld-example-app/internal/app"
	"tkktrd/golang-gin-realworld-example-app/pkg/db/mysql"
)

func main() {

	// Read and parse configuration files.
	conf, err := config.LoadConfig("config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	db, err := mysql.NewMysqlConnect(conf)
	if err != nil {
		log.Fatalf("MySQL init: %s", err)
	}

	// TODO: error 
	app := app.NewApp(conf, db)

	app.Run()
}
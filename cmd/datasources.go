package main

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/constracts"
	"Sesuai/pkg/alog"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

func InitDatasource() {
	var err error
	var dbWriter *sqlx.DB
	var dbReader *sqlx.DB

	dsWriter, dsReader := parseDs()

	if dbWriter, err = sqlx.Connect(app.Config[constants.DbDialeg], dsWriter); err == nil {
		dbWriter.SetConnMaxLifetime(time.Duration(1) * time.Second)
		dbWriter.SetMaxOpenConns(10)
		dbWriter.SetMaxIdleConns(10)

		alog.Logger.Printf("Initalizing Writer DB: Pass")
	} else {
		alog.Logger.Printf("error while connecting to writer db: %s", err)
	}

	if dbReader, err = sqlx.Connect(app.Config[constants.DbDialeg], dsReader); err == nil {
		dbReader.SetConnMaxLifetime(time.Duration(1) * time.Second)
		dbReader.SetMaxOpenConns(10)
		dbReader.SetMaxIdleConns(10)

		alog.Logger.Printf("Initalizing Reader DB: Pass")
	} else {
		alog.Logger.Printf("error while connecting to reader db: %s", err)
	}

	ds := &constracts.Datasources{
		WriterDB: dbWriter,
		ReaderDB: dbReader,
	}

	app.Datasources = ds

	return
}

func parseDs() (dsWriter, dsReader string) {
	hostWriter := app.Config[constants.DbHostWriter]
	hostReader := app.Config[constants.DbHostReader]
	port := app.Config[constants.DbPort]
	user := app.Config[constants.DbUser]
	pass := app.Config[constants.DbPass]
	name := app.Config[constants.DbName]

	dsWriter = fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, hostWriter, port, name)
	dsReader = fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, hostReader, port, name)

	return
}

package models

import (
	"database/sql"
	//"fmt"
	//"errors"
	//stat "github.com/asiainfoLDP/datafoundry_appmarket/statistics"
	//"github.com/asiainfoLDP/datahub_commons/log"
)

type DatabaseUpgrader_4 struct {
	DatabaseUpgrader_Base
}

func newDatabaseUpgrader_4() *DatabaseUpgrader_4 {
	updater := &DatabaseUpgrader_4{}

	updater.currentTableCreationSqlFile = "initdb_v003.sql"

	updater.oldVersion = 4
	updater.newVersion = 5

	return updater
}

func (upgrader DatabaseUpgrader_4) Upgrade(db *sql.DB) error {
	return nil
}

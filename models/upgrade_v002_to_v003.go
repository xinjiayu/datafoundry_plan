package models

import (
	"database/sql"
	//"fmt"
	//"errors"

	//stat "github.com/asiainfoLDP/datahub_commons/statistics"
	"github.com/asiainfoLDP/datahub_commons/log"
)



type DatabaseUpgrader_3 struct {
	DatabaseUpgrader_Base
	
	AlterSQL string
}

func newDatabaseUpgrader_3() *DatabaseUpgrader_3 {
	updater := &DatabaseUpgrader_3{}
	
	updater.oldVersion = 3
	updater.newVersion = 4
	updater.AlterSQL = `
		ALTER TABLE DF_PLAN
		MODIFY CUSTOM_PRICES VARCHAR(4096) NOT NULL DEFAULT '' COMMENT 'nested json'
	`
	
	return updater
}

func (upgrader DatabaseUpgrader_3) Upgrade (db *sql.DB) error {
	
	log.DefaultLogger().Info("DatabaseUpgrader_3 started ... ") 
	
	// ...
	
	log.DefaultLogger().Info("DatabaseUpgrader_3 alter tables ... ") 
	
	_, _ = db.Exec(upgrader.AlterSQL)
	//if err != nil {
	//	return err
	//}
	//n, _ := result.RowsAffected()
	
	log.DefaultLogger().Info("DatabaseUpgrader_3, alter tables done. ")
	
	// 
	
	return nil
}

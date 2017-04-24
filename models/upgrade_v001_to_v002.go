package models

import (
	"database/sql"
	//"fmt"
	//"errors"

	//stat "github.com/asiainfoLDP/datahub_commons/statistics"
	"github.com/asiainfoLDP/datahub_commons/log"
)



type DatabaseUpgrader_1 struct {
	DatabaseUpgrader_Base
	
	AlterSQL string
}

func newDatabaseUpgrader_1() *DatabaseUpgrader_1 {
	updater := &DatabaseUpgrader_1{}
	
	updater.oldVersion = 1
	updater.newVersion = 2
	updater.AlterSQL = `
		ALTER TABLE DF_PLAN
		ADD COLUMN CUSTOM_PRICES VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'nested json' AFTER PRICE
	`
	
	return updater
}   

func (upgrader DatabaseUpgrader_1) Upgrade (db *sql.DB) error {
	
	log.DefaultLogger().Info("DatabaseUpgrader_1 started ... ") 
	
	// ...
	
	log.DefaultLogger().Info("DatabaseUpgrader_1 alter tables ... ") 
	
	_, _ = db.Exec(upgrader.AlterSQL)
	//if err != nil {
	//	return err
	//}
	//n, _ := result.RowsAffected()
	
	log.DefaultLogger().Info("DatabaseUpgrader_1, alter tables done. ")
	
	// 
	
	return nil
}

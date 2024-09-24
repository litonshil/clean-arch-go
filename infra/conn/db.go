package conn

import (
	"clean-arch/config"
	"fmt"
	"time"

	//commonLogger "github.com/klikit/utils/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var db *gorm.DB

func ConnectDB() {
	masterConf := config.DB().Master
	replicaConf := config.DB().Replica1

	//commonLogger.Info("connecting to mysql at ", masterConf.Host, ":", masterConf.Port, "...")
	fmt.Println("connecting to mysql at ", masterConf.Host, ":", masterConf.Port, "...")

	logMode := logger.Silent
	if masterConf.Debug {
		logMode = logger.Info
	}

	masterDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		masterConf.Username, masterConf.Password, masterConf.Host, masterConf.Port, masterConf.Name)
	replicaDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", replicaConf.Username,
		replicaConf.Password, replicaConf.Host, replicaConf.Port, replicaConf.Name)

	fmt.Printf(masterDSN)
	dB, err := gorm.Open(mysql.Open(masterDSN), &gorm.Config{
		PrepareStmt: masterConf.PrepareStmt,
		Logger:      logger.Default.LogMode(logMode),
	})
	if err != nil {
		panic(err)
	}

	dB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{
			mysql.Open(replicaDSN),
		},
		Policy: dbresolver.RandomPolicy{},
	}))

	sqlDb, err := dB.DB()
	if err != nil {
		panic(err)
	}

	if masterConf.MaxIdleConn != 0 {
		sqlDb.SetMaxIdleConns(masterConf.MaxIdleConn)
	}
	if masterConf.MaxOpenConn != 0 {
		sqlDb.SetMaxOpenConns(masterConf.MaxOpenConn)
	}
	if masterConf.MaxLifeTime != 0 {
		sqlDb.SetConnMaxLifetime(masterConf.MaxLifeTime * time.Second)
	}

	db = dB
	//commonLogger.Info("mysql connection successful...")
	fmt.Println("mysql connection successful...")
}

func Db() *gorm.DB {
	return db
}

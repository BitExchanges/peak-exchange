package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	envConfig "peak-exchange/config"
	"strings"
	"time"
)

var (
	MainDb, BackupDb *gorm.DB
)

type GormDB struct {
	*gorm.DB
	gdbDone bool
}

func getConnectionString(config *ConfigEnv, name string) string {
	host := config.Get(envConfig.CurrentEnv.Model+"."+name+".host", "")
	port := config.Get(envConfig.CurrentEnv.Model+"."+name+".port", "")
	user := config.Get(envConfig.CurrentEnv.Model+"."+name+".username", "")
	passwd := config.Get(envConfig.CurrentEnv.Model+"."+name+".password", "")
	dbname := config.Get(envConfig.CurrentEnv.Model+"."+name+".database", "")
	protocol := config.Get(envConfig.CurrentEnv.Model+"."+name+".protocol", "tcp")
	dbargs := config.Get(envConfig.CurrentEnv.Model+"."+name+".dbargs", " ")
	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, passwd, protocol, host, port, dbname, dbargs)
}

// 初始化主数据库
func InitMainDB() {
	config := getDatabaseConfig()
	var connStr = getConnectionString(config, "main")
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(config.GetInt(envConfig.CurrentEnv.Model+".main.pool", 5))
	db.DB().SetMaxOpenConns(config.GetInt(envConfig.CurrentEnv.Model+".main.maxopen", 0))
	du, _ := time.ParseDuration(config.Get(envConfig.CurrentEnv.Model+".main.timeout", "3600") + "s")
	db.DB().SetConnMaxLifetime(du)
	//开启日志
	//db.LogMode(true)
	//不使用复数
	db.SingularTable(true)
	db.Exec("set transaction isolation level repeatable read")
	MainDb = db
}

// 初始化备份数据库
func InitBakupDB() {
	config := getDatabaseConfig()
	var connstring = getConnectionString(config, "backup")
	db, err := gorm.Open(connstring)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(config.GetInt(envConfig.CurrentEnv.Model+".backup.pool", 5))
	db.DB().SetMaxOpenConns(config.GetInt(envConfig.CurrentEnv.Model+".backup.maxopen", 0))
	du, _ := time.ParseDuration(config.Get(envConfig.CurrentEnv.Model+".backup.timeout", "3600") + "s")
	db.DB().SetConnMaxLifetime(du)
	db.Exec("set transaction isolation level repeatable read")
	BackupDb = db
}

func CloseMainDB() {
	MainDb.Close()
}

func CloseBackupDB() {
	BackupDb.Close()
}

func MainDbBegin() *GormDB {
	txn := MainDb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	return &GormDB{txn, false}
}

func BackupDbBegin() *GormDB {
	txn := BackupDb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	return &GormDB{txn, false}
}

func (c *GormDB) DbCommit() {
	if c.gdbDone {
		return
	}
	tx := c.Commit()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}

func (c *GormDB) DbRollback() {
	if c.gdbDone {
		return
	}
	tx := c.Rollback()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrConnDone {
		panic(err)
	}
}

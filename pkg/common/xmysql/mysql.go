package xmysql

import (
	"database/sql"
	"errors"
	"fmt"
	logger "github.com/ipfs/go-log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rcsv/pkg/conf"
)

var (
	log                      = logger.Logger("xmysql")
	ERR_DB_INSTANCE_IS_EMPTY = errors.New("database instance is empty")
)

var (
	cli *MysqlClient
)

type MysqlClient struct {
	db  *gorm.DB
	cfg conf.Mysql
}

func NewMysqlClient(cfg conf.Mysql) *MysqlClient {
	cli = &MysqlClient{cfg: cfg}
	cli.db, _ = ConnectDB(cfg)
	return cli
}

func GetDB() *gorm.DB {
	if cli.db == nil {
		cli.db, _ = ConnectDB(cli.cfg)
	}
	return cli.db
}

func GetTX() *gorm.DB {
	return GetDB().Begin()
}

func Transaction(handle func(tx *gorm.DB) (err error)) (err error) {
	var (
		db *gorm.DB
	)
	db = GetDB()
	if db == nil {
		err = ERR_DB_INSTANCE_IS_EMPTY
		return
	}
	tx := db.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err = handle(tx)
	if err != nil {
		err = tx.Rollback().Error
		return
	}
	err = tx.Commit().Error
	return
}

func ConnectDB(cfg conf.Mysql) (db *gorm.DB, err error) {
	var (
		args  string
		opts  *gorm.Config
		sqlDB *sql.DB
	)
	args = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Address,
		cfg.Db)

	opts = &gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            false,
	}

	db, err = gorm.Open(mysql.Open(args), opts)
	if err != nil {
		log.Error(err.Error())
		return
	}
	db = db.Debug()

	sqlDB, err = db.DB()
	if err != nil {
		log.Error(err.Error())
		return
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	return
}

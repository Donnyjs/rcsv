package main

import (
	logger "github.com/ipfs/go-log"
	"rcsv/apps/rcsv/config"
	"rcsv/apps/rcsv/internal/server"
	"rcsv/pkg/commands"
	"rcsv/pkg/common/xmysql"
	"rcsv/pkg/common/xoss"
	"rcsv/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
	xoss.NewOssClient(conf.Oss)
}

func main() {
	logger.SetLogLevel("*", "INFO")
	commands.Run(server.NewServer())
}

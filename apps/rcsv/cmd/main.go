package main

import (
	"rcsv/apps/rcsv/internal/config"
	"rcsv/apps/rcsv/internal/server"
	"rcsv/pkg/commands"
	"rcsv/pkg/common/xmysql"
	"rcsv/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	commands.Run(server.NewServer())
}

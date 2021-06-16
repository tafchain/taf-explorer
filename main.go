package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"tafexplorer/db"
	"tafexplorer/log"
	_ "tafexplorer/routers"
	"tafexplorer/taf_server"
)

// Version版本信息
var (
	Version = "unknown"
	Build   = "2021-03-09"
)

func main() {
	err := db.InitMgo()
	if err != nil {
		panic(err)
	}
	err = db.InitRedis()
	if err != nil {
		panic(err)
	}
	taf_server.InitTafServer()
	log.InitLogger()
	logs.Info("System init done. Version:[%s] Build time:[%s]\n", Version, Build)
	beego.SetStaticPath("/swagger", "swagger")
	beego.Run()
}










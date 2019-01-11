package main

import (
	//	"fmt"
	"mime"
	"os"

	"./models"
	_ "./routers"
	"./util"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	//初始化
	initialize()

	if beego.AppConfig.String("runmode") == "dev" {
		beego.BConfig.Listen.EnableAdmin = true
		//orm.Debug = true
	}
}

func initialize() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	mime.AddExtensionType(".css", "text/css")
	models.Connect()
	//router()
	//beego.AddFuncMap("stringsToJson", StringsToJson)
}

func dbinit() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}

func main() {
	//dbinit()
	Init()
	util.WritePidFile()
	beego.Run()

}

package main

import (
	"fmt"
	_ "zonjli_blog/routers"
	_ "zonjli_blog/utils"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	username, _ := beego.AppConfig.String("username")
	password, _ := beego.AppConfig.String("password")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	database, _ := beego.AppConfig.String("database")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, database)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", datasource)

	fmt.Println("数据连接成功")

	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)

	if err != nil {
		panic(err)
	}

}

func main() {
	orm.Debug = true

	//beego.InsertFilter("/blog/*", beego.BeforeRouter, utils.BlogLoginFilter)
	orm.RunCommand()
	beego.Run()
}

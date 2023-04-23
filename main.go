package main

import (
	"fmt"
	_ "mahasiswa/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp(localhost:3306)/student")
	db, err := orm.GetDB("student")
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("DB Connected Succesfully ", db)
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

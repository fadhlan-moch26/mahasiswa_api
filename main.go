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
	orm.RegisterDataBase("default", "mysql", "os.Getenv("DB_USER"):os.Getenv("DB_PASS")@tcp(os.Getenv("DB_HOST"))/os.Getenv("DB_NAME")")
	db, err := orm.GetDB("os.Getenv("DB_NAME")")
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

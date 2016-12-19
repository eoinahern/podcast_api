package main

import (
	_ "podcast_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", "file:sqlite/podcast.db")

	if err != nil {
		panic(err)
	}

}

func main() {
	o := orm.NewOrm()
	o.Using("default")
	beego.Run()

}

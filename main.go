package main

import (
	"podcast_api/models"
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

	beego.Run()
	o := orm.NewOrm()
	o.Using("default")
	orm.RegisterModel(new(models.Admin))

	orm.RunSyncdb("default", false, true)

}

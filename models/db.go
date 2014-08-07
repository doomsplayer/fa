package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

var Engine *xorm.Engine

func init() {
	var err error

	Engine, err = xorm.NewEngine(beego.AppConfig.String("db"), beego.AppConfig.String("dbc"))
	if err != nil {
		beego.Critical("cannot open database:", err)
		time.Sleep(200 * time.Millisecond)
		os.Exit(1)
	}
	err = Engine.Sync2(
		new(File),
		new(Promotion),
		new(Carousel),
		new(PromotionType),
		new(TutorialType),
		new(Tutorial),
	)
	if beego.AppConfig.String("runmode") == "dev" {
		Engine.ShowSQL = true
	}
	if err != nil {
		beego.Critical("cannot sync model:", err)
		time.Sleep(200 * time.Millisecond)
		os.Exit(1)
	}
}

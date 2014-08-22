package main

import (
	_ "badmintonhome/docs"
	_ "badmintonhome/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.TemplateLeft = "{%"
	beego.TemplateRight = "%}"
	beego.Run()
}

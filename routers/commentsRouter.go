package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenuApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenuApi"],
		beego.ControllerComments{
			"Menu",
			"/mainBar",
			[]string{"get"},
			nil})

}

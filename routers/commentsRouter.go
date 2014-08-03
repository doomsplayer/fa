package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenu"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenu"],
		beego.ControllerComments{
			"Menu",
			"/mainBar",
			[]string{"get"},
			nil})

}

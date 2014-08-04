package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"],
		beego.ControllerComments{
			"UploadFile",
			"/upload",
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenuApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenuApi"],
		beego.ControllerComments{
			"Menu",
			"/mainBar",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenuApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:MainMenuApi"],
		beego.ControllerComments{
			"Carousel",
			"/mainMenu/carousel",
			[]string{"get"},
			nil})

}

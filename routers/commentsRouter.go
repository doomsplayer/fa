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

	beego.GlobalControllerRouter["badmintonhome/controllers/api:PromotionApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:PromotionApi"],
		beego.ControllerComments{
			"Get",
			"/promotion",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:PromotionApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:PromotionApi"],
		beego.ControllerComments{
			"AllTypes",
			"/promotiontypes",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:CarouselApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:CarouselApi"],
		beego.ControllerComments{
			"Add",
			"/carousel",
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:PromotionApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:PromotionApi"],
		beego.ControllerComments{
			"Put",
			"/promotion",
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:PromotionApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:PromotionApi"],
		beego.ControllerComments{
			"Delete",
			"/promotion",
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:PromotionApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:PromotionApi"],
		beego.ControllerComments{
			"AddTypes",
			"/promotiontypes",
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:PromotionApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/adminapi:PromotionApi"],
		beego.ControllerComments{
			"DelTypes",
			"/promotiontypes",
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:CarouselApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:CarouselApi"],
		beego.ControllerComments{
			"List",
			"/carousel",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"],
		beego.ControllerComments{
			"UploadFile",
			"/upload",
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"],
		beego.ControllerComments{
			"DeleteFile",
			"/upload",
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"] = append(beego.GlobalControllerRouter["badmintonhome/controllers/api:FileApi"],
		beego.ControllerComments{
			"FetchFile",
			"/upload",
			[]string{"get"},
			nil})

}

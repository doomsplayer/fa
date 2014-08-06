// @APIVersion 1.0.0
// @Title 羽球之家网站API
// @Description 羽球之家网站API。
// @Contact doomsplayer@gmail.com
package routers

import (
	"badmintonhome/controllers"
	"badmintonhome/controllers/adminapi"
	"badmintonhome/controllers/api"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ApiNS := beego.NewNamespace("/api",
		beego.NSNamespace("/admin",
			beego.NSInclude(&adminapi.CarouselApi{}, &adminapi.PromotionApi{}),
		),
		beego.NSNamespace("/common",
			beego.NSInclude(&api.MainMenuApi{}, &api.FileApi{}, &api.CarouselApi{},
				&api.PromotionApi{}),
		),
	)
	beego.AddNamespace(ApiNS)

}

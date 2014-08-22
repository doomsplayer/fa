// @APIVersion 1.0.0
// @Title 羽球之家网站API
// @Description 羽球之家网站API。
// @Contact doomsplayer@gmail.com
package routers

import (
	"badmintonhome/controllers"
	"badmintonhome/controllers/admin"
	"badmintonhome/controllers/adminapi"
	"badmintonhome/controllers/api"
	"badmintonhome/controllers/ueditor"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ApiNS := beego.NewNamespace("/api",
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&adminapi.CarouselApi{},
				&adminapi.PromotionApi{},
				&adminapi.TutorialApi{},
				&adminapi.VideoApi{},
			),
		),
		beego.NSNamespace("/common",
			beego.NSInclude(
				&api.MainMenuApi{},
				&api.FileApi{},
				&api.CarouselApi{},
				&api.PromotionApi{},
				&api.TutorialApi{},
				&api.VideoApi{},
			),
		),
	)

	beego.AddNamespace(ApiNS)
	var AdminNs = beego.NewNamespace(`/admin`)
	AdminNs.Include(&admin.AdminController{}, &admin.LoginController{})
	beego.AddNamespace(AdminNs)

	ue := beego.NewNamespace(`/ue`)
	ue.Include(&ueditor.UController{})
	beego.AddNamespace(ue)
}

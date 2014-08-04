// @APIVersion 1.0.0
// @Title badminton API
// @Description 羽球网。
// @Contact doomsplayer@gmail.com
package routers

import (
	"badmintonhome/controllers"
	"badmintonhome/controllers/api"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ApiNS := beego.NewNamespace("v1",
		beego.NSNamespace("/api",
			beego.NSInclude(&api.MainMenuApi{}),
		),
	)
	beego.AddNamespace(ApiNS)

}

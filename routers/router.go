// @APIVersion 1.0.0
// @Title 羽球之家网站API
// @Description 羽球之家网站API。
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
			beego.NSInclude(&api.MainMenuApi{}, &api.FileApi{}),
		),
	)
	beego.AddNamespace(ApiNS)

}

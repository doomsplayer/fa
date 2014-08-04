package lib

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) Prepare() {
	b.SetSession(`username`, `test`)

}

type AuthController struct {
	BaseController
}

func (a *AuthController) Prepare() {
	a.BaseController.Prepare()
	if a.GetSession(`username`) == nil {
		a.Abort("401")
		return
	}
}

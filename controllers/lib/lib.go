package lib

import (
	"github.com/astaxie/beego"
)

type more []string

func (m *more) Add(s string) string {
	*m = append(*m, s)
	return ``
}

type BaseController struct {
	beego.Controller
}

func (b *BaseController) Prepare() {
	b.Data[`moreStyles`] = &more{}
	b.Data[`moreScripts`] = &more{}
	b.Data["position"] = ``
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

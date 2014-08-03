package lib

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) Prepare() {

}

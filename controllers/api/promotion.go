package api

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type PromotionApi struct {
	lib.BaseController
}

// @Title promotion
// @Description 获得促销商品
// @Param from query int false 从第几项开始,0是第一个,默认0
// @Param num query int false 几个，默认1
// @Param type query string true 类型（球鞋拍子blahblah）
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /promotion [get]
func (p *PromotionApi) Get() {
	type list struct {
		From int    `form:"from" valid:"Min(0)"`
		Num  int    `form:"num" valid:"Min(0)"`
		Type string `form:"type" valid:"Required"`
	}
	l := new(list)
	err := p.ParseForm(l)
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}

	valid := validation.Validation{}
	b, err := valid.Valid(l)
	if err != nil {

		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}
	if !b {
		errmsg := ``
		for _, err := range valid.Errors {
			errmsg += fmt.Sprint(err.Key, " ", err.Message)
		}
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "form valid not pass: " + errmsg}
		p.ServeJson()
		return
	}

	if l.Num == 0 {
		l.Num = 1
	}

	pl := &models.PromotionSlice{}
	err = pl.Hot(l.Num, l.From, l.Type)
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}
	if len(*pl) == 0 {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "empty"}
		p.ServeJson()
		return
	}
	p.Data["json"] = map[string]interface{}{"ok": true, "promotions": pl}
	p.ServeJson()
	return

}

// @Title mainBar
// @Description 获取促销商品分类
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /promotiontypes [get]
func (p *PromotionApi) AllTypes() {
	tp := &models.PromotionTypeSlice{}
	err := tp.GetAll()
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}
	p.Data["json"] = map[string]interface{}{"ok": true, "promotiontypes": tp}
	p.ServeJson()
	return
}

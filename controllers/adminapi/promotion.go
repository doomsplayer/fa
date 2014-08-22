package adminapi

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
)

type PromotionApi struct {
	lib.AuthController
}

// @Title mainBar
// @Description 添加促销商品
// @Param title form string true 标题
// @Param title2 form string true 小标题
// @Param description form string true 描述
// @Param description2 form string true 小描述
// @Param type form string true 商品类型
// @Param link form string true 连接地址
// @Param detail form string true 商品细节
// @Param picid form string true 大图片id
// @Param promotetime form string true 促销时间，用格式 2014-09-01~2015-01-01
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /promotion [put]
func (p *PromotionApi) Put() {

	pro := new(models.Promotion)
	err := p.ParseForm(pro)
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}

	valid := validation.Validation{}
	b, err := valid.Valid(pro)
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
	promotetimeS := p.GetString("promotetime")

	s := strings.Split(promotetimeS, "~")
	starttime, _ := time.Parse("2006-01-02", s[0])
	endtime, _ := time.Parse("2006-01-02", s[1])
	pro.StartTime = starttime
	pro.EndTime = endtime
	err = pro.Put()
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}

	p.Data["json"] = map[string]interface{}{"ok": true, "promotion": pro}
	p.ServeJson()
	return

}

// @Title mainBar
// @Description 删除促销商品
// @Param id query int true 要删除的id
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /promotion [delete]
func (p *PromotionApi) Delete() {
	id, err := p.GetInt("id")
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}
	pro := new(models.Promotion)
	pro.Id = id
	err = pro.Delete()
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}
	p.Data["json"] = map[string]interface{}{"ok": true}
	p.ServeJson()
	return
}

// @Title PromotionTypes
// @Description 添加商品分类
// @Param name form string true 商品分类的名字
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /promotiontypes [put]
func (p *PromotionApi) AddTypes() {
	typs := p.GetString("name")
	if strings.TrimSpace(typs) == `` {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "name is empty"}
		p.ServeJson()
		return
	}
	tp := new(models.PromotionType)
	tp.Name = typs
	err := tp.Put()
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}
	p.Data["json"] = map[string]interface{}{"ok": true, "promotiontypes": tp}
	p.ServeJson()
	return
}

// @Title PromotionTypes
// @Description 删除商品分类
// @Param id query string true 商品分类的id
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /promotiontypes [delete]
func (p *PromotionApi) DelTypes() {
	typid, err := p.GetInt("id")
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}
	tp := new(models.PromotionType)
	tp.Id = typid

	err = tp.Delete()
	if err != nil {
		p.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		p.ServeJson()
		return
	}
	p.Data["json"] = map[string]interface{}{"ok": true}
	p.ServeJson()
	return
}

package adminapi

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type CarouselApi struct {
	lib.AuthController
}

// @Title add carousel
// @Description 添加首页轮播图
// @Param title form string true 标题
// @Param subtitle form string true 小标题
// @Param url form string true Url
// @Param picid form int true 图片的Id
// @Success 200 {string} 结果json
// @Failure 404 Not found errmsg
// @Failure 401 Need Permmision errmsg
// @Failure 500 Server Error errmsg
// @router /carousel [put]
func (c *CarouselApi) Add() {

	carousel := new(models.Carousel)
	err := c.ParseForm(carousel)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		c.ServeJson()
		return
	}

	valid := validation.Validation{}
	b, err := valid.Valid(carousel)
	if err != nil {

		c.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		c.ServeJson()
		return
	}
	if !b {
		errmsg := ``
		for _, err := range valid.Errors {
			errmsg += fmt.Sprint(err.Key, " ", err.Message)
		}
		c.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "form valid not pass: " + errmsg}
		c.ServeJson()
		return
	}
	err = carousel.Put()

	c.Data["json"] = map[string]interface{}{"ok": true, "carousels": carousel}
	c.ServeJson()
}

// @Title delete carousel
// @Description 添加首页轮播图
// @Param id form int true Id
// @Success 200 {string} 结果json
// @Failure 404 Not found errmsg
// @Failure 401 Need Permmision errmsg
// @Failure 500 Server Error errmsg
// @router /carousel [delete]
func (c *CarouselApi) Delete() {
	var (
		id, _ = c.GetInt("id")
	)
	carousel := new(models.Carousel)
	carousel.Id = id
	carousel.Delete()

	c.Data["json"] = map[string]interface{}{"ok": true, "carousels": carousel}
	c.ServeJson()
}

// @Title use carousel
// @Description 首页轮播图使能
// @Param id form int true carouselId
// @Param use form bool true 使能与否
// @Success 200 {string} 结果json
// @Failure 404 Not found errmsg
// @Failure 401 Need Permmision errmsg
// @Failure 500 Server Error errmsg
// @router /carouseluse [post]
func (c *CarouselApi) SetUse() {
	carousel := new(models.Carousel)
	var (
		id, _  = c.GetInt(`id`)
		use, _ = c.GetBool(`use`)
	)

	carousel.Id = id
	carousel.Get()
	carousel.Use = use
	carousel.Update()

	c.Data["json"] = map[string]interface{}{"ok": true, "carousels": carousel}
	c.ServeJson()
}

package api

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
)

type CarouselApi struct {
	lib.BaseController
}

// @Title list carousel
// @Description 拉取首页轮播图
// @Param num query int true 拉取几张
// @Success 200 {string} 轮播的json
// @Failure 404 Not found errmsg
// @Failure 401 Need Permmision errmsg
// @Failure 500 Server Error errmsg
// @router /carousel [get]
func (c *CarouselApi) List() {

	num, err := c.GetInt("num")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		c.ServeJson()
		return
	}
	cs := &models.CarouselSlice{}
	err = cs.GetUse(int(num))
	if err != nil {
		c.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		c.ServeJson()
		return
	}

	c.Data["json"] = map[string]interface{}{"ok": true, "carousels": cs}
	c.ServeJson()
}

package api

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type VideoApi struct {
	lib.BaseController
}

// @Title video
// @Description 获得视频
// @Param from query int false 从第几项开始,0是第一个,默认0
// @Param num query int false 几个，默认1
// @Param type query string false 类型 大话羽球blahblah
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /video [get]
func (t *VideoApi) Get() {
	type list struct {
		From int    `form:"from" valid:"Min(0)"`
		Num  int    `form:"num" valid:"Min(0)"`
		Type string `form:"type" valid:""`
	}
	l := new(list)
	err := t.ParseForm(l)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}

	valid := validation.Validation{}
	b, err := valid.Valid(l)
	if err != nil {

		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	if !b {
		errmsg := ``
		for _, err := range valid.Errors {
			errmsg += fmt.Sprint(err.Key, " ", err.Message)
		}
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "form valid not pass: " + errmsg}
		t.ServeJson()
		return
	}

	if l.Num == 0 {
		l.Num = 1
	}

	tl := &models.VideoSlice{}
	err = tl.All(l.Num, l.From, l.Type)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	if len(*tl) == 0 {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "empty"}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true, "videos": tl}
	t.ServeJson()
	return
}

// @Title video
// @Description 获得单一视频
// @Param id path id true id
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /video/:id [get]
func (t *VideoApi) Single() {
	ids := t.Ctx.Input.Param(`:id`)
	id, err := strconv.Atoi(ids)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	v := new(models.Video)
	v.Id = int64(id)
	err = v.Get()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true, "video": v}
	t.ServeJson()
	return
}

// @Title video
// @Description 获得热门视频
// @Param from query int false 从第几项开始,0是第一个,默认0
// @Param num query int false 几个，默认1
// @Param type query string true 类型 大话羽球blahblah
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /hotvideo [get]
func (t *VideoApi) Hot() {
	type list struct {
		From int    `form:"from" valid:"Min(0)"`
		Num  int    `form:"num" valid:"Min(0)"`
		Type string `form:"type" valid:""`
	}
	l := new(list)
	err := t.ParseForm(l)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}

	valid := validation.Validation{}
	b, err := valid.Valid(l)
	if err != nil {

		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	if !b {
		errmsg := ``
		for _, err := range valid.Errors {
			errmsg += fmt.Sprint(err.Key, " ", err.Message)
		}
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "form valid not pass: " + errmsg}
		t.ServeJson()
		return
	}

	if l.Num == 0 {
		l.Num = 1
	}

	tl := &models.VideoSlice{}
	err = tl.Hot(l.Num, l.From, l.Type)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	if len(*tl) == 0 {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "empty"}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true, "videos": tl}
	t.ServeJson()
	return
}

// @Title video
// @Description 获取视频分类
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /videotypes [get]
func (t *VideoApi) AllTypes() {
	tt := &models.VideoTypeSlice{}
	err := tt.GetAll()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true, "videotypes": tt}
	t.ServeJson()
	return
}

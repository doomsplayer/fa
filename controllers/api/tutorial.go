package api

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type TutorialApi struct {
	lib.BaseController
}

// @Title tutorial
// @Description 获得教程
// @Param from query int false 从第几项开始,0是第一个,默认0
// @Param num query int false 几个，默认1
// @Param type query string true 类型 大话羽球blahblah
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /tutorial [get]
func (t *TutorialApi) Get() {
	type list struct {
		From int    `form:"from" valid:"Min(0)"`
		Num  int    `form:"num" valid:"Min(0)"`
		Type string `form:"type" valid:"Required"`
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

	tl := &models.TutorialSlice{}
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
	t.Data["json"] = map[string]interface{}{"ok": true, "tutorials": tl}
	t.ServeJson()
	return
}

// @Title tutorial
// @Description 获得热门教程
// @Param from query int false 从第几项开始,0是第一个,默认0
// @Param num query int false 几个，默认1
// @Param type query string true 类型 大话羽球blahblah
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /hottutorial [get]
func (t *TutorialApi) Hot() {
	type list struct {
		From int    `form:"from" valid:"Min(0)"`
		Num  int    `form:"num" valid:"Min(0)"`
		Type string `form:"type" valid:"Required"`
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

	tl := &models.TutorialSlice{}
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
	t.Data["json"] = map[string]interface{}{"ok": true, "tutorials": tl}
	t.ServeJson()
	return
}

// @Title tutorial
// @Description 获取文章分类
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /tutorialtypes [get]
func (t *TutorialApi) AllTypes() {
	tt := &models.TutorialTypeSlice{}
	err := tt.GetAll()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true, "tutorialtypes": tt}
	t.ServeJson()
	return
}

// @Title tutorial
// @Description 获得教程详细内容
// @Param id path int true id
// @Success 200 {string} 文章内容json
// @Failure 404 Not found
// @router /tutorial/:id [get]
func (t *TutorialApi) Single() {
	ids := t.Ctx.Input.Param(`:id`)
	id, err := strconv.Atoi(ids)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}

	tot := new(models.Tutorial)
	tot.Id = int64(id)

	err = tot.Get()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true, "tutorial": tot}
	t.ServeJson()
	return
}

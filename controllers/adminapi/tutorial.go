package adminapi

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"
)

type TutorialApi struct {
	lib.AuthController
}

// @Title tutorial
// @Description 添加教程
// @Param title form string true 标题
// @Param type form string true 教程类型
// @Param picid form int true 大图片id
// @Param content form string true 教程内容
// @Param author form string false 教程作者
// @Param source form string false 教程来源
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /tutorial [put]
func (t *TutorialApi) Put() {

	tut := new(models.Tutorial)
	err := t.ParseForm(tut)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}

	valid := validation.Validation{}
	b, err := valid.Valid(tut)
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

	err = tut.Put()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}

	t.Data["json"] = map[string]interface{}{"ok": true, "tutorial": tut}
	t.ServeJson()
	return

}

// @Title tutorial
// @Description 删除教程
// @Param id query int true 要删除的id
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /tutorial [delete]
func (t *TutorialApi) Delete() {
	id, err := t.GetInt("id")
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	tut := new(models.Tutorial)
	tut.Id = id
	err = tut.Delete()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true}
	t.ServeJson()
	return
}

// @Title tutorialTypes
// @Description 添加文章分类
// @Param name form string true 文章分类的名字
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /tutorialtypes [put]
func (t *TutorialApi) AddTypes() {
	typs := t.GetString("name")
	if strings.TrimSpace(typs) == `` {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "name is empty"}
		t.ServeJson()
		return
	}
	tt := new(models.TutorialType)
	tt.Name = typs
	err := tt.Put()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true, "tutorialtypes": tt}
	t.ServeJson()
	return
}

// @Title TutorialTypes
// @Description 删除文章分类
// @Param id query string true 文章分类的id
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /tutorialtypes [delete]
func (t *TutorialApi) DelTypes() {
	typid, err := t.GetInt("id")
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	tt := new(models.TutorialType)
	tt.Id = typid

	err = tt.Delete()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true}
	t.ServeJson()
	return
}

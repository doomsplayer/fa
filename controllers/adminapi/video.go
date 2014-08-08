package adminapi

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"
)

type VideoApi struct {
	lib.AuthController
}

// @Title video
// @Description 添加视频
// @Param title form string true 标题
// @Param type form string true 视频类型
// @Param picid form int true 大图片id
// @Param content form string true 视频内容（url）
// @Param author form string false 视频作者
// @Param source form string false 视频来源
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /video [put]
func (t *VideoApi) Put() {

	tut := new(models.Video)
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

	t.Data["json"] = map[string]interface{}{"ok": true, "video": tut}
	t.ServeJson()
	return

}

// @Title video
// @Description 删除视频
// @Param id query int true 要删除的id
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /video [delete]
func (t *VideoApi) Delete() {
	id, err := t.GetInt("id")
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	tut := new(models.Video)
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

// @Title videoTypes
// @Description 添加视频分类
// @Param name form string true 视频分类的名字
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /videotypes [put]
func (t *VideoApi) AddTypes() {
	typs := t.GetString("name")
	if strings.TrimSpace(typs) == `` {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": "name is empty"}
		t.ServeJson()
		return
	}
	tt := new(models.VideoType)
	tt.Name = typs
	err := tt.Put()
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	t.Data["json"] = map[string]interface{}{"ok": true, "videotypes": tt}
	t.ServeJson()
	return
}

// @Title VideoTypes
// @Description 删除视频分类
// @Param id query string true 视频分类的id
// @Success 200 {string} 列表的json
// @Failure 404 Not found
// @router /videotypes [delete]
func (t *VideoApi) DelTypes() {
	typid, err := t.GetInt("id")
	if err != nil {
		t.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		t.ServeJson()
		return
	}
	tt := new(models.VideoType)
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

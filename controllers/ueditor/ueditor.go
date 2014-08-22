package ueditor

import (
	"badmintonhome/controllers/lib"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type UController struct {
	lib.BaseController
}

// @router / [get,post]
func (this *UController) Main() {
	switch this.GetString("action") {
	case "config":
		{
			f, _ := os.Open(`config.json`)
			b, _ := ioutil.ReadAll(f)
			this.Ctx.Output.Body(b)
			f.Close()
			return
		}
	case "uploadimage":
		{

			filename := fmt.Sprint(time.Now().Unix())
			err := this.SaveToFile("upfile", "static/upload/ue/"+filename)

			if err != nil {
				beego.Critical("get upload file fail:", err)
			}
			_, header, _ := this.GetFile("upfile")
			this.Data[`json`] = map[string]interface{}{
				"filetype": filepath.Ext(header.Filename),
				"original": header.Filename,
				"state":    "SUCCESS",
				"url":      "/static/upload/ue/" + filename,
			}
			this.ServeJson()
		}
	}
}

package ueditor

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"fmt"
	"io/ioutil"
	"os"
)

type UController struct {
	lib.BaseController
}

// @router / [get,post]
func (this *UController) Main() {
	fmt.Print("be")
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

			file, header, err := this.GetFile("upfile")
			if err != nil {
				this.Ctx.Output.Status = 500
				this.Ctx.Output.Json(map[string]interface{}{"ok": false, "errmsg": err.Error()}, false, false)
				return
			}
			username := `admin`

			dbfile := new(models.File)

			data, _ := ioutil.ReadAll(file)

			err = dbfile.SaveToFile(header.Filename, username, data)
			if err != nil {
				this.Data[`json`] = map[string]interface{}{
					"title":    header.Filename,
					"original": header.Filename,
					"state":    "FAIL",
				}
				this.ServeJson()
				return
			} else {
				this.Data[`json`] = map[string]interface{}{
					"title":    header.Filename,
					"original": header.Filename,
					"state":    "SUCCESS",
					"url":      "/" + dbfile.Path,
				}
				this.ServeJson()
				return
			}

		}
	case `listimage`:
		{
			start, _ := this.GetInt(`start`)
			num, _ := this.GetInt(`size`)
			dir, _ := os.Open(`static/upload/ue`)
			fileinfos, _ := dir.Readdir(-1)
			m := map[string]interface{}{}
			m[`state`] = `SUCCESS`
			m["start"] = start
			m["total"] = len(fileinfos)
			s := []map[string]string{}
			for num = num - 1; num >= 0; num-- {
				s = append(s, map[string]string{
					"url": "/static/upload/ue/" + fileinfos[num+start].Name(),
				})
			}
			m[`list`] = s
			this.Data[`json`] = m
			this.ServeJson()
			return
		}
	}
}

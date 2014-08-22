package api

import (
	"badmintonhome/controllers/lib"
	"badmintonhome/models"
	"io/ioutil"
)

type FileApi struct {
	lib.AuthController
}

// @Title Upload File
// @Description 上传文件
// @Param file form file true 上传的文件
// @Success 200 {string} 文件Id的json
// @Failure 404 Not found errmsg
// @Failure 401 Need Permmision errmsg
// @Failure 500 Server Error errmsg
// @router /upload [put,post]
func (f *FileApi) UploadFile() {

	file, header, err := f.GetFile("file")
	if err != nil {
		f.Ctx.Output.Status = 500
		f.Ctx.Output.Json(map[string]interface{}{"ok": false, "errmsg": err.Error()}, false, false)
		return
	}
	username := f.GetSession(`username`).(string)

	dbfile := new(models.File)

	data, _ := ioutil.ReadAll(file)

	err = dbfile.SaveToFile(header.Filename, username, data)
	if err != nil {
		f.Ctx.Output.Status = 500
		f.Ctx.Output.Json(map[string]interface{}{"ok": false, "errmsg": err.Error()}, false, false)
		return
	} else {
		f.Data[`json`] = map[string]interface{}{"ok": true, "fileId": dbfile.Id}
		f.ServeJson()
	}

}

// @Title Delete File
// @Description 删除文件
// @Param id query int true 文件id
// @Success 200 {string} 返回json
// @Failure 404 Not found errmsg
// @Failure 401 Need Permmision errmsg
// @Failure 500 Server Error errmsg
// @router /upload [delete]
func (f *FileApi) DeleteFile() {
	id, err := f.GetInt("id")
	if err != nil {
		f.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		f.ServeJson()
		return
	}

	dbfile := new(models.File)

	dbfile.Id = id
	err = dbfile.Delete()

	if err != nil {
		f.Ctx.Output.Status = 500
		f.Ctx.Output.Json(map[string]interface{}{"ok": false, "errmsg": err.Error()}, false, false)
		return
	} else {
		f.Data[`json`] = map[string]interface{}{"ok": true}
		f.ServeJson()
		return
	}

}

// @Title Get File
// @Description 获得文件
// @Param id query int true 文件id
// @Success 200 {string} 返回json
// @Failure 404 Not found errmsg
// @Failure 401 Need Permmision errmsg
// @Failure 500 Server Error errmsg
// @router /upload [get]
func (f *FileApi) FetchFile() {
	id, err := f.GetInt("id")
	if err != nil {
		f.Data["json"] = map[string]interface{}{"ok": false, "errmsg": err.Error()}
		f.ServeJson()
		return
	}

	dbfile := new(models.File)

	dbfile.Id = id
	err = dbfile.Get()

	if err != nil {
		f.Ctx.Output.Status = 500
		f.Ctx.Output.Json(map[string]interface{}{"ok": false, "errmsg": err.Error()}, false, false)
		return
	} else {
		f.Data[`json`] = map[string]interface{}{"ok": true, "filename": dbfile.Name, "filepath": dbfile.Path, "fileid": dbfile.Id}
		f.ServeJson()
		return
	}

}

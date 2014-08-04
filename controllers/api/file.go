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
// @Param type form string true 上传文件的类型
// @Success 200 {string} 文件Id的json
// @Failure 404 Not found errmsg
// @Failure 401 Need Permmision errmsg
// @Failure 500 Server Error errmsg
// @router /upload [put]
func (f *FileApi) UploadFile() {

	file, header, err := f.GetFile("file")
	if err != nil {
		f.Ctx.Output.Status = 500
		f.Ctx.Output.Json(map[string]interface{}{"ok": false, "errmsg": err}, false, false)
		return
	}
	username := f.GetSession(`username`).(string)

	dbfile := new(models.File)

	data, _ := ioutil.ReadAll(file)

	err = dbfile.SaveToFile(header.Filename, username, data)
	if err != nil {
		f.Ctx.Output.Status = 500
		f.Ctx.Output.Json(map[string]interface{}{"ok": false, "errmsg": err}, false, false)
		return
	} else {
		f.Data[`json`] = map[string]interface{}{"ok": true, "fileId": dbfile.Id}
		f.ServeJson()
	}

}

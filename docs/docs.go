package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

var rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/api","description":""}],"info":{"title":"badminton API","description":"羽球网。","contact":"doomsplayer@gmail.com"}}`
var subapi string = `{"/api":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/api","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/mainBar","description":"","operations":[{"httpMethod":"GET","nickname":"mainBar","type":"","summary":"获得首页目录","responseMessages":[{"code":200,"message":"目录json","responseModel":""},{"code":404,"message":"Not found","responseModel":""}]}]}]}}`
var rootapi swagger.ResourceListing

var apilist map[string]*swagger.ApiDeclaration

func init() {
	basepath := "v1"
	err := json.Unmarshal([]byte(rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = basepath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
